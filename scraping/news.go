package scraping

import(
	tf "github.com/HoritakuDev/KeyakiReaderWebApi/time_formatter"
	"github.com/HoritakuDev/KeyakiReaderWebApi/models"
	"github.com/HoritakuDev/KeyakiReaderWebApi/common"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"io/ioutil"
	"log"
	"time"
)

const(
	DATETIME_NEWS_FORMAT = "200601"
	NEWS_SAVING_DIR = "/var/tmp/keyaki/news.html"
)

func scrapeAllNews() {
	loc, _ := time.LoadLocation(tf.LOCATION_TOKYO)
	target_time := time.Date(2015, 10, 1, 0, 0, 0, 0,  loc)
	end_time := time.Now()
	for end_time.After(target_time) {
		scrapeNews(common.NEWS_LIST_URL + "&dy=" + target_time.Format(DATETIME_NEWS_FORMAT))
		target_time = target_time.AddDate(0, 1, 0)
	}
}

func scrapeRecentNews() {
	scrapeNews(common.NEWS_LIST_URL)
}

func scrapeNews(url string) {

	log.Println("start scraping from : " + url)
	downloadFile(url, NEWS_SAVING_DIR)

	// --- read html file --- //
	file_infos, _ := ioutil.ReadFile(NEWS_SAVING_DIR)
	str_reader := strings.NewReader(string(file_infos))
	doc, err := goquery.NewDocumentFromReader(str_reader)
	if err != nil {
		log.Println(err)
	}

	// --- scraping --- //
	news_list := make([]models.News, 0, 30)
	doc.Find("div.box-news").Find("li").Each(func(_ int, article *goquery.Selection) {
		url, _ := article.Find("a").Attr("href")
		title := article.Find("a").Text()
		category := strings.TrimSpace(article.Find("div.category").Text())
		updated_str := strings.TrimSpace(article.Find("div.date").Text())

		news := models.News{
			Link: url,
			Title: title,
			Category: category,
		}
		tf.SetTimeInJST(updated_str, &news)

		tmp_id, _ := dbmap.SelectInt("SELECT id FROM news where title = '" + news.Title + "' updated = '" + tf.FormatTimeToStr(news.Updated) + "'", &models.News{})
		if tmp_id != 0 {
			return
		}
		// add news to head of list
		news_list = append(news_list[:1], news_list[0:]...)
		news_list[0] = news
	})
	for _, news := range news_list {
		dbmap.Insert(&news)
	}
}
