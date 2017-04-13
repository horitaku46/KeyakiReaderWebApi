package server

import(
	"github.com/HoritakuDev/KeyakiReaderWebApi/models"
	"github.com/HoritakuDev/KeyakiReaderWebApi/common"
	"net/http"
	"encoding/json"
	"regexp"
	"strconv"
	"log"
)

const(
	// --- for sql --- ///
	LIMIT_BLOG_ARTICLE = "20"
	LIMIT_NEWS_ARTICLE = "20"
	DATA_ORDER = "ORDER BY id ASC"
)

var(
	// --- for validation parameters --- ///
	ID_PATTERN = regexp.MustCompile(`^[0-9]{1,}$`)
)

func setCommonHeader(w *http.ResponseWriter)  {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(http.StatusOK)
}

func parseParameters(req *http.Request) (params map[string]interface{}) {
	p_str := req.FormValue("params")
	log.Println(p_str)
	json.Unmarshal([]byte(p_str), &params)
	return
}

func validateIntParam(params map[string]interface{}, key string) (int, bool) {
	if tmp, ok := params[key]; ok {
		log.Println("interface{} : " + key)
		if tmp_int, ok := tmp.(float64); ok {
			log.Println("float64 : " + key)
			return int(tmp_int), true
		}
	}
	return 0, false
}


func getAllBlogs(w http.ResponseWriter, req *http.Request) {
	setCommonHeader(&w)

	// --- "scope" is ID-range of records which will be gathered --- //
	scope := make(map[string]int)

	params := parseParameters(req)
	log.Println(params)

	// The parameters has "bottom_id".
	if oldest, ok := validateIntParam(params, "bottom_id"); ok {
		scope["start"] = oldest-21
		scope["end"] = oldest-1
	// has "top_id".
	} else if newest, ok := validateIntParam(params, "top_id"); ok {
		scope["start"] = newest+1
		tmp_id, _ := dbmap.SelectInt("SELECT MAX(id) FROM blogs")
		scope["end"] = int(tmp_id)
	// doesn't have either items.
	} else {
		max_id, _ := dbmap.SelectInt("SELECT MAX(id) FROM blogs")
		scope["start"] = int(max_id-20)
		scope["end"] = int(max_id)
	}

	var blogs models.ApiBlogList
	if err := blogs.SelectAllBetween(dbmap, scope); err == nil {
		if response, err := json.Marshal(blogs); err == nil {
			w.Write( response )
			return
		}
	}
	w.Write( []byte("[]") )
}

func getIndividualBlogs(w http.ResponseWriter, req *http.Request) {
	setCommonHeader(&w)

	// --- "scope" is ID-range of records which will be gathered --- //
	scope := make(map[string]int)

	params := parseParameters(req)
	log.Println(params)

	if newest, ok := validateIntParam(params, "top_id"); ok {
		scope["top_id"] = newest
	} else if oldest, ok := validateIntParam(params, "bottom_id"); ok {
		scope["bottom_id"] = oldest
	} else {
		tmp_int64, _ := dbmap.SelectInt("SELECT MAX(id) FROM blogs")
		scope["bottom_id"] = int(tmp_int64)
	}

	if member_id, ok := validateIntParam(params, "member_id"); ok {
		var blogs models.ApiBlogList
		if err := blogs.SelectIndiBetween(dbmap, scope, member_id); err == nil {
			if response, err := json.Marshal(blogs); err == nil {
				w.Write( response )
				return
			}
		}
	}
	w.Write( []byte("[]") )
}

func getNews(w http.ResponseWriter, req *http.Request) {
	setCommonHeader(&w)
	params := parseParameters(req)
	log.Println(params)

	var news []models.News
	var condition string
	if newest, ok := validateIntParam(params, "top_id"); ok {
		condition = " id > " + strconv.Itoa(newest)
	} else if oldest, ok := validateIntParam(params, "bottom_id"); ok {
		condition = " id BETWEEN " + strconv.Itoa(oldest-21) + " AND " + strconv.Itoa(oldest-1)
	} else {
		record_num, _ := dbmap.SelectInt("SELECT MAX(id) FROM news")
		condition = " id BETWEEN " + strconv.Itoa(int(record_num-21)) + " AND " + strconv.Itoa(int(record_num-1))
	}
	if _, err := dbmap.Select( &news, "SELECT * FROM news WHERE" + condition + " ORDER BY id ASC"); err == nil {
		response, _ := json.Marshal(news)
		w.Write(response)
		return
	}
	w.Write( []byte("[]") )
}

func getMembers(w http.ResponseWriter, req *http.Request) {
	setCommonHeader(&w)
	var members []models.Member
	if _, err := dbmap.Select(&members, "SELECT * FROM members ORDER BY id ASC"); err == nil {
		if response, err := json.Marshal(members); err == nil {
			w.Write( response )
			return
		}
	}
	w.Write( []byte("[]") )
}

func getImages(w http.ResponseWriter, req *http.Request) {
	setCommonHeader(&w)
}

func getAllConf(w http.ResponseWriter, req *http.Request) {
	setCommonHeader(&w)
	conf_map := map[string]string{
		"blog_upper_url": common.BLOG_UPPDER_URL,
		"image_upper_url": common.IMAGE_UPPDER_URL,
	}
	json_byte, _ := json.Marshal(conf_map)
	w.Write( json_byte )
}

// --- test method --- //
func echoTest(w http.ResponseWriter, req *http.Request) {
	setCommonHeader(&w)
	params := parseParameters(req)
	buff, _ := json.Marshal(&params)
	w.Write( buff )
}
