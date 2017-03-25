package main

import(
	"net/http"
	"encoding/json"
	"regexp"
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
}

func getAllBlogs(w http.ResponseWriter, req *http.Request) {
	var blogs []Blog
	var condition string
	setCommonHeader(&w)
	if newest := req.FormValue("top_id"); ID_PATTERN.MatchString(newest) {
		condition = "id > " + newest
	} else if oldest := req.FormValue("bottom_id"); ID_PATTERN.MatchString(oldest) {
		condition = "id < " + oldest + " LIMIT " + LIMIT_BLOG_ARTICLE
	}
	if _, err := dbmap.Select( &blogs, "SELECT * FROM blogs WHERE " + condition ); err != nil {
		if response, err := json.Marshal(blogs); err == nil {
			w.Write( response )
			return
		}
	}
	w.Write( []byte("[]") )
}

func getIndividualBlogs(w http.ResponseWriter, req *http.Request) {
	setCommonHeader(&w)
	var blogs []Blog
	if member_id := req.FormValue("member_id"); ID_PATTERN.MatchString(member_id) {
		if _, err := dbmap.Select( &blogs, "SELECT * FROM blogs WHERE writer_id = " + member_id + " ORDER BY id ASC"); err == nil {
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
	var news []News
	var condition string
	if newest := req.FormValue("top_id"); ID_PATTERN.MatchString(newest) {
		condition = " id > " + newest
	} else if oldest := req.FormValue("bottom_id"); ID_PATTERN.MatchString(oldest) {
		condition = " id < " + oldest + " LIMIT " + LIMIT_NEWS_ARTICLE
	}
	condition += " ORDER BY id ASC"
	if _, err := dbmap.Select( &news, "SELECT * FROM news WHERE" + condition + "ORDER BY id ASC"); err == nil {
		response, _ := json.Marshal(news)
		w.Write(response)
		return
	}
	w.Write( []byte("[]") )
}

func getMembers(w http.ResponseWriter, req *http.Request) {
	setCommonHeader(&w)
	var members []Member
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
