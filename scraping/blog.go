package scraping

import(
	tf "github.com/HoritakuDev/KeyakiReaderWebApi/time_formatter"
	"github.com/HoritakuDev/KeyakiReaderWebApi/common"
	"github.com/HoritakuDev/KeyakiReaderWebApi/models"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"io/ioutil"
	"log"
	"strconv"
)

const (
	BLOG_PARAM_SPECIFY_PAGE = "?page="
	BLOG_SAVING_DIR = "/var/tmp/keyaki/blog.html"
	BLOG_MAX_PAGE = 420
)

type BlogImgPair struct {
	Blog models.Blog
	Images []models.Image
}

// Insert Articles and its Images.
func insertBlogImgPairs(articles []BlogImgPair) {
	for _, pair := range articles {
		if len(pair.Images) > 0 {
			pair.Blog.Thumbnail = pair.Images[0].Url
		} else {
			pair.Blog.Thumbnail = ""
		}
		dbmap.Insert(&pair.Blog)
		for _, img := range pair.Images {
			img.ArticleId = pair.Blog.Id
			dbmap.Insert(&img)
		}
	}
}

// Scrape all blogs.
func scrapeAllBlogs() {
	for i := BLOG_MAX_PAGE; i >= 0; i-- {
		articles, _ :=scrapeBlog(common.URL_ALL_BLOG_LIST + BLOG_PARAM_SPECIFY_PAGE + strconv.Itoa(i))
		insertBlogImgPairs(articles)
	}
}

// Scrape blogs which recently added.
func scrapeRecentBlogs() {
	should_continue := true	// indicates unreached an article which stored database?
	var articles []BlogImgPair
	var tmp []BlogImgPair
	for i := 0; should_continue; i++ {
		tmp, should_continue = scrapeBlog(common.URL_ALL_BLOG_LIST + BLOG_PARAM_SPECIFY_PAGE + strconv.Itoa(i))
		articles = append(tmp, articles...)
	}

	// --- insert into database --- //
	insertBlogImgPairs(articles)
}

// Scrape articles from one HTML file.
func scrapeBlog(url string) (articles []BlogImgPair, should_continue bool) {
	log.Println("start scraping from : " + url)
	downloadFile(url, BLOG_SAVING_DIR)
	// --- Read articles from one HTML file. --- //
	file_infos, _ := ioutil.ReadFile(BLOG_SAVING_DIR)
	str_reader := strings.NewReader(string(file_infos))
	doc, err := goquery.NewDocumentFromReader(str_reader)

	should_continue = true
	if err != nil {
		log.Println(err)
	}
	articles = make([]BlogImgPair, 20)
	doc.Find("article").Each(func(index int, article *goquery.Selection) {
		var tmp_art BlogImgPair
		tmp_art, should_continue = scrapeArticle(article)
		// --- Read articles from one HTML file. --- //
		if should_continue { // should_continue.(Target article is not registerd yet)
			articles = append(articles[:1], articles[0:]...)
			articles[0] = tmp_art
		}
	})

	return
}

// Get BlogImgPair from an article and judge continue scraping.
func scrapeArticle(article *goquery.Selection) (datum BlogImgPair, should_continue bool) {
	header := article.Find("div.innerHead")
	content := article.Find("div.box-article")
	bottom := article.Find("div.box-bottom")

	url, _ := header.Find("a").Attr("href")
	title := header.Find("a").Text()
	writer := strings.TrimSpace(header.Find("p.name").Text())
	updated_str := strings.TrimSpace(bottom.Find("li").First().Text())

	should_continue = true

	// --- get member's id from database --- //
	tmp_id, _ := dbmap.SelectInt("SELECT ID FROM members WHERE name = '" + writer + "'", &models.Member{})
	if tmp_id == 0 {
		log.Println("Member's name is not founded : " + writer)
		return
	}

	// --- represent this article --- //
	datum = BlogImgPair{
		Blog: models.Blog{
			Title: title,
			Link: url,
			Writer: int32(tmp_id),
		},
		Images: make([]models.Image, 0, 20),
	}
	tf.SetTimeInJST(updated_str, &datum.Blog)

	// --- If this article exists on database, this function doesn't work. --- //
	var tmp_articles []models.Blog
	dbmap.Select(&tmp_articles, "SELECT * FROM blogs WHERE link_url = '" + datum.Blog.Link + "'")
	if(len(tmp_articles) != 0) {
		should_continue = false
		return
	}
	log.Println(datum.Blog.Link)

	// --- scrapes images which are included in an ariticle --- //
	content.Find("img").Each(func(_ int, img_tag *goquery.Selection) {
		tmp_url, _ := img_tag.Attr("src")
		// -- appends tips of url to list  -- //
		splited_url := strings.SplitAfterN(tmp_url, "/", 4)
		datum.Images = append(datum.Images, models.Image{Url: splited_url[len(splited_url)-1]})
	})
	return
}
