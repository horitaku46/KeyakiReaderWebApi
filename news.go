package main

type News struct {
	Id int32	`db:"id, primarykey, autoincrement"`
	Link string	`db:"link_url"`
	Title string	`db:"title"`
	Category string	`db:"category"`
	Updated string	`db:"updated"`
}
