package models

import "time"

type News struct {
	Id int32		`db:"id, primarykey, autoincrement"`
	Link string		`db:"link_url, notnull"`
	Title string		`db:"title, notnull"`
	Category string		`db:"category, notnull"`
	Updated time.Time	`db:"updated, notnull"`
}
