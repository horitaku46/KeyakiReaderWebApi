package models

import (
	"time"
)

type Member struct {
	Id int32		`db:"id, primarykey, autoincrement" json:"member_id"`
	Thumbnail string		`db:"thumbnail, notnull" json:"thmbnail_url"`
	Name string		`db:"name, notnull"                 json:"member_name"`
	Ruby string		`db:"ruby, notnull"                 json:"name_ruby"`
	Updated time.Time	`db:"updated, notnull"              json:"blog_update_time"`
}
