package main

type Blog struct {
	Id int32		`json:"blog_id"			db:"id, primarykey, autoincrement"`
	Title string		`json:"blog_title"		db:"title"`
	Link string		`json:"blog_url"		db:"link_utl"`
	Writer int32		`json:"blog_writer"		db:"writer_id"`
	Thumbnail string	`json:"blog_image_url"		db:"tumbnail_url"`
	Updated string		`json:"blog_update_time"	db:"updated"`
}
