package scraping

import(
	"github.com/go-gorp/gorp"
	"time"
	"log"
	"net/http"
	"io/ioutil"
)

const(
	INTERVAL_MIN = 5
)

var(
	dbmap *gorp.DbMap
)

// Activation of scraping routine.
func StartScraping(arg_dbmap *gorp.DbMap, should_scrape_all_blogs bool) {

	log.Println("scraping cron is activated.")
	dbmap = arg_dbmap
	scrapeMembers()
	if should_scrape_all_blogs {	// Do this process if '-init' options added. 
		scrapeAllBlogs()
		scrapeAllNews()
	}
	go func() {
		for {				// infinit loop
			select {	// one-scraping
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

// Download HTML file from the Internet
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
