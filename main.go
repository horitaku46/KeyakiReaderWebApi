package main

import(
	"github.com/HoritakuDev/KeyakiReaderWebApi/models"
	"github.com/HoritakuDev/KeyakiReaderWebApi/server"
	"github.com/HoritakuDev/KeyakiReaderWebApi/scraping"
	"github.com/HoritakuDev/KeyakiReaderWebApi/common"
	"github.com/go-gorp/gorp"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"flag"
)

const HTTP_PORT = ":9000"

var(
	// --- for accessing database --- //
	DB_PASSWD = os.Getenv("DB_PASSWD")
	DB_USER = os.Getenv("DB_USER")
	DB_NAME = os.Getenv("DB_NAME")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = "3306"
)


func main() {

	should_scrape_all_blogs := checkFlags()

	// -- ready database --- //
	server_dbmap := initDb();
	cron_dbmap := initDb();
	defer server_dbmap.Db.Close()
	defer cron_dbmap.Db.Close()

	scraping.StartScraping(cron_dbmap, should_scrape_all_blogs)
	server.ActivateServer(server_dbmap, HTTP_PORT)
}

func initDb() (dbmap *gorp.DbMap) {

	// --- connect to database and create mapping entity for database --- //
	db, err := sql.Open("mysql", DB_USER + ":" + DB_PASSWD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?parseTime=true&loc=Asia%2FTokyo")
	common.CheckErr(err, "sql.Open() is failed.")

	dbmap = &gorp.DbMap{ Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"} }

	// --- ready tables --- //
	dbmap.AddTableWithName(models.Blog{}, "blogs").SetKeys(true, "Id")
	dbmap.AddTableWithName(models.Member{}, "members").SetKeys(true, "Id")
	dbmap.AddTableWithName(models.News{}, "news").SetKeys(true, "Id")
	dbmap.AddTableWithName(models.Image{}, "images").SetKeys(true, "Id")
	dbmap.AddTableWithName(models.Client{}, "clients").SetKeys(true, "Id")
	err = dbmap.CreateTablesIfNotExists()
	common.CheckErr(err, "Creating tables is failed.")

	return

}

func checkFlags() (should_scrape_all_blogs bool) {
	// --- check arguments (Is it includes "init" option?) --- //
	flag.BoolVar(&should_scrape_all_blogs, "init", false, "Do initial scraping")

	flag.Parse()
	return
}
