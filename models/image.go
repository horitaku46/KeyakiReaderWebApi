package models

type Image struct {
	Id int32		`db:"id, primarykey, autoincrement"`
	Url string		`db:"url, notnull"`
	ArticleId int32		`db:"article_id, notnull"`
}
