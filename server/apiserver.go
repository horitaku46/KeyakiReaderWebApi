package server

import(
	"github.com/go-gorp/gorp"
	"net/http"
	"log"
)
var (
	dbmap *gorp.DbMap

	http_funcs = []ActionPair{
		{"/blogs/all/get", getAllBlogs},
		{"/blogs/individual/get", getIndividualBlogs},
		{"/news/get", getNews},
		{"/members/get", getMembers},
		{"/images/get", getImages},
		{"/conf/get", getAllConf},
	}
)

type ActionPair struct {
	path string
	function func(http.ResponseWriter, *http.Request)
}

func ActivateServer(arg_dbmap *gorp.DbMap, port string) {
	dbmap = arg_dbmap

	// --- register functions to htttp server --- //
	for i := range http_funcs {
		pair := http_funcs[i]
		http.HandleFunc(pair.path, pair.function)
	}
	log.Fatal( http.ListenAndServe(port, nil) )
}

