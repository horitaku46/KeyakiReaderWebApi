package main

import(
	Id		`db:"id, primarykey, autoincrement"`
	Url		`db:"url"`
	ArticleId	`db:"article_id"`
	Updated		`db:"updated"`
)
