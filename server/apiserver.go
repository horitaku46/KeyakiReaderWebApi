package main

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-gorp/gorp"
	"os"
	"log"
	"net/http"
)

const(
	// --- for accessing database --- //
	db_user = "keyaki"
	db_name = "keyaki"
	db_host = "sqlserver"
	db_port = "3306"
	// --- for accessing database --- //
	http_port = ":9000"
	// --- blog's url which upper from file name (scheme, host, directory) --- //
	blog_uppder_url = "http://www.keyakizaka46.com/s/k46o/diary/member/list"
	image_uppder_url = "http://cdn.keyakizaka46.com/images/"
	news_uppder_url = "http://www.keyakizaka46.com/s/k46o/news/detail/"
	member_uppder_url = "http://www.keyakizaka46.com/s/k46o/artist/"
)

var (
	db_passwd = os.Getenv("DB_PASSWD")
	http_funcs = []ActionPair{
		{"/blogs/all/get", getAllBlogs},
		{"/blogs/individual/get", getIndividualBlogs},
		{"/news/get", getNews},
		{"/members/get", getMembers},
		{"/images/get", getImages},
	}
	dbmap *gorp.DbMap
)

type ActionPair struct {
	path string
	function func(http.ResponseWriter, *http.Request)
}

func main() {

	// -- ready database --- //
	dbmap = initDb();
	defer dbmap.Db.Close()

	// --- register functions to htttp server --- //
	for i := range http_funcs {
		pair := http_funcs[i]
		http.HandleFunc(pair.path, pair.function)
	}
	log.Fatal( http.ListenAndServe(http_port, nil) )
}

func initDb() (dbmap *gorp.DbMap) {

	// --- connect to database --- //
	db, err := sql.Open("mysql", db_user + ":" + db_passwd + "@tcp(" + db_host + ":" + db_port + ")/" + db_name + "?parseTime=true&loc=Asia%2FTokyo")
	checkErr(err, "sql.Open() is failed.")

	dbmap = &gorp.DbMap{ Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"} }

	// --- setup for logging --- //
	file, err := os.Open(`./log/gorp.log`)
	checkErr(err, "Cannot open log file.")
	dbmap.TraceOn("[gorp]", log.New(file, "apiserver:", log.Lmicroseconds))

	// --- ready tables --- //
	dbmap.AddTableWithName(Blog{}, "blogs").SetKeys(true, "Id")
	dbmap.AddTableWithName(Member{}, "members").SetKeys(true, "Id")
	dbmap.AddTableWithName(News{}, "news").SetKeys(true, "Id")
	dbmap.AddTableWithName(Image{}, "images").SetKeys(true, "Id")
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Creating tables is failed.")

	return
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatal("[error] " + msg, err)
	}
}

