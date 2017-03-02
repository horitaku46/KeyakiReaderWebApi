package main

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-gorp/gorp"
	"os"
	"log"
)

const(
	db_user = "keyaki"
	db_name = "keyaki"
	db_host = "sqlserver"
	db_port = "3306"
)
var (
	db_passwd = os.Getenv("DB_PASSWD")
)

func initDb() (dbmap *gorp.DbMap) {

	log.Println("mysql", db_user + ":" + db_passwd + "@tcp(" + db_host + ":" + db_port + ")/" + db_name)
	db, err := sql.Open("mysql", db_user + ":" + db_passwd + "@tcp(" + db_host + ":" + db_port + ")/" + db_name)
	checkErr(err, "sql.Open() is failed.")

	dbmap = &gorp.DbMap{ Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"} }

	dbmap.AddTableWithName(Blog{}, "blogs").SetKeys(true, "Id")
	dbmap.AddTableWithName(Member{}, "members").SetKeys(true, "Id")
	dbmap.AddTableWithName(News{}, "news").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Creating tables is failed.")

	return
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatal("[error] " + msg, err)
	}

}

func main() {
	dbmap := initDb();
	defer dbmap.Db.Close()
}
