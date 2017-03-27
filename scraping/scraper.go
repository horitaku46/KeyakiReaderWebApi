package scraping

import(
	"github.com/go-gorp/gorp"
	"time"
	"log"
	"net/http"
	"io/ioutil"
	"flag"
)

const(
	INTERVAL_MIN = 1
)

var(
	dbmap *gorp.DbMap
)

func StartScraping(arg_dbmap *gorp.DbMap) {

	// --- check arguments (Is it includes "init" option?) --- //
	var should_scrape_all_blogs bool
	flag.BoolVar(&should_scrape_all_blogs, "init", false, "Do initial scraping")
	flag.Parse()

	log.Println("scraping cron is activated.")
	dbmap = arg_dbmap
	scrapeMembers()
	if should_scrape_all_blogs {
		scrapeAllBlogs()
		scrapeAllNews()
	}
	go func() {
		for {
			select {
			case <-time.After(INTERVAL_MIN * time.Minute):
				scrape()
				log.Println("run scraping")
			}
		}
	}()
}

func scrape() {
	scrapeMembers()
	scrapeRecentBlogs()
	scrapeRecentNews()
}

func downloadFile(url string, file_path string) {
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		log.Println("file downloading")
		log.Println(err);
		return
	}
	body, _ := ioutil.ReadAll(res.Body)
	if err := ioutil.WriteFile(file_path, body, 0644); err != nil {
		log.Println("file saving")
		log.Println(err)
	}
}
