package models

import "time"

type News struct {
	Id int32		`db:"id, primarykey, autoincrement"	json:"news_id"`
	Link string		`db:"link_url, notnull"			json:"news_url"`
	Title string		`db:"title, notnull"			json:"news_title"`
	Category string		`db:"category, notnull"			json:"news_category"`
	Updated time.Time	`db:"updated, notnull"			json:"news_update_time"`
}
