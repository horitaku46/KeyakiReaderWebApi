package scraping

import(
	//tf "github.com/HoritakuDev/KeyakiReaderWebApi/time_formatter"
	"github.com/HoritakuDev/KeyakiReaderWebApi/common"
	"github.com/HoritakuDev/KeyakiReaderWebApi/models"
	"github.com/PuerkitoBio/goquery"
	//"time"
	"strings"
	"io/ioutil"
	"log"
)

const(
	MEMBER_ALL_PARAM = "?ima=000"
	MEMBER_SAVING_DIR = "/var/tmp/keyaki/member.html"
)

func scrapeMembers() {
	downloadFile(common.MEMBER_UPPER_URL + MEMBER_ALL_PARAM, MEMBER_SAVING_DIR)
	file_infos, _ := ioutil.ReadFile(MEMBER_SAVING_DIR)
	str_reader := strings.NewReader(string(file_infos))
	doc, err := goquery.NewDocumentFromReader(str_reader)
	if err != nil {
		log.Println(err)
	}

	// --- read member's information from article --- //
	scraped_members := make([]models.Member, 0, 50)
	scrape_member := func(_ int, member_info *goquery.Selection) {
		scraped_members = append( scraped_members, getMemberFromLI(member_info) )
	}
	tmp_div := doc.Find("div.sorted.sort-default.current").Find("div.box-member")
	tmp_div.First().Find("li").Each(scrape_member) // keyaki-zaka64
	tmp_div.Next().Find("li").Each(scrape_member) // hiragana-keyaki

	// --- check already existed in database --- //
	var existed_members []models.Member
	dbmap.Select(&existed_members, "SELECT * FROM members")
	for _, sm := range scraped_members {
		insertUpdateIfNeeded(sm, &existed_members)
	}
}

func getMemberFromLI(member_info *goquery.Selection) (member models.Member) {
	member = models.Member{
		Name: strings.TrimSpace(member_info.Find("p.name").Text()),
		Ruby: strings.TrimSpace(member_info.Find("p.furigana").Text()),
	}
	member.Thumbnail, _ = member_info.Find("img").Attr("src")
	return
}

func insertUpdateIfNeeded(chk_target models.Member, domain *[]models.Member) {
	for _, m := range *domain {
		if chk_target.Name == m.Name && chk_target.Ruby == m.Ruby {
			if chk_target.Thumbnail != (common.IMAGE_UPPER_URL + m.Thumbnail) {
				splited_url := strings.SplitAfterN(chk_target.Thumbnail, "/", 4)
				m.Thumbnail = splited_url[len(splited_url)-1]
				dbmap.Update(&m)
			}
			return
		}
	}
	dbmap.Insert(&chk_target)
}
