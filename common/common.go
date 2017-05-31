package common

import(
	"log"
)

const(
	URL_ALL_BLOG_LIST = "http://www.keyakizaka46.com/s/k46o/diary/member/list"
	// --- blog's url which upper from file name (scheme, host, directory) --- //
	BLOG_UPPER_URL = "http://www.keyakizaka46.com/"
	IMAGE_UPPER_URL = "http://cdn.keyakizaka46.com/"
	NEWS_UPPER_URL = "http://www.keyakizaka46.com/s/k46o/news/detail/"
	MEMBER_UPPER_URL = "http://www.keyakizaka46.com/s/k46o/search/artist/"

	NEWS_LIST_URL = "http://www.keyakizaka46.com/s/k46o/news/list?ima=000"
)

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatal("[error] " + msg, err)
	}
}
