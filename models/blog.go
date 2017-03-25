package models

import (
	"time"
)

type Blog struct {
	Id int32		`db:"id, primarykey, autoincrement" json:"blog_id"`
	Title string		`db:"title, notnull"                json:"blog_title"`
	Link string		`db:"link_url, notnull"             json:"blog_url"`
	Writer int32		`db:"writer_id, notnull"	    json:"writer_id"`
	Thumbnail string	`db:"thumbnail_url"                  json:"blog_image_url"`
	Updated time.Time	`db:"updated, notnull"              json:"blog_update_time"`
}
