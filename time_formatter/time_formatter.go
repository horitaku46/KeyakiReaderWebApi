package time_formatter

import(
	"github.com/HoritakuDev/KeyakiReaderWebApi/models"
	"time"
)

const (
	LOCATION_TOKYO = "Asia/Tokyo"
	BLOG_DATETIME_FORMAT = "2006/01/02 15:04"
	NEWS_DATETIME_FORMAT = "2006.01.02"
	MEMBER_DATETIME_FORMAT = ""
	MYSQL_DATETIME_FORMAT = "2006-01-02 15:04:05"
)

func SetTimeInJST(time_str string, model interface{}) {

	loc, err := time.LoadLocation(LOCATION_TOKYO)

	if err != nil {
		loc = time.FixedZone(LOCATION_TOKYO, 9*60*60)
	}

	switch v := model.(type) {
	case *models.Blog:
		v.Updated, _ = time.ParseInLocation(BLOG_DATETIME_FORMAT, time_str, loc)
	case *models.News:
		v.Updated, _ = time.ParseInLocation(NEWS_DATETIME_FORMAT, time_str, loc)
	case *models.Member:
		v.Updated, _ = time.ParseInLocation(MEMBER_DATETIME_FORMAT, time_str, loc)
	}
}

func FormatTimeToStr(datetime time.Time) string {
	return datetime.Format(MYSQL_DATETIME_FORMAT)
}
