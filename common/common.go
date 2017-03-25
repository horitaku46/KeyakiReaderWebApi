package common

import(
	"log"
)

const(
	// --- blog's url which upper from file name (scheme, host, directory) --- //
	BLOG_UPPDER_URL = "http://www.keyakizaka46.com/s/k46o/diary/member/list"
	IMAGE_UPPDER_URL = "http://cdn.keyakizaka46.com/"
	NEWS_UPPDER_URL = "http://www.keyakizaka46.com/s/k46o/news/detail/"
	MEMBER_UPPDER_URL = "http://www.keyakizaka46.com/s/k46o/search/artist"

	NEWS_LIST_URL = "http://www.keyakizaka46.com/s/k46o/news/list?ima=000"
)

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatal("[error] " + msg, err)
	}
}
