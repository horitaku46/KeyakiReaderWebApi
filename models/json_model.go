package models

import(
	"time"
	"log"
	"strconv"
	"github.com/go-gorp/gorp")

const(
	SQL_SELECT_API_BLOGS = `
		SELECT
			blogs.id		AS id,
			blogs.title		AS title,
			blogs.link_url		AS link_url,
			members.name		AS writer_name,
			blogs.thumbnail_url	AS thumbnail_url,
			blogs.updated		AS updated
		FROM
			blogs
		LEFT OUTER JOIN
			members
		ON
			blogs.writer_id = members.id
		WHERE
	`
)

type ApiBlog struct {
	Id int32	`db:"id" json:"blog_id"`
	Title string		`db:"title, notnull"                json:"blog_title"`
	Url string	`db:"link_url" json:"blog_url"`
	Writer string	`db:"writer_name" json:"blog_writer"`
	Image string	`db:"thumbnail_url" json:"blog_image_url"`
	Updated time.Time	`db:"updated, notnull"              json:"blog_update_time"`
}

type ApiBlogList []ApiBlog

func (this *ApiBlogList) SelectIndiBetween(dbmap *gorp.DbMap, scope map[string]int, member_id int) (err error) {
	log.Println(scope)
	sql := SQL_SELECT_API_BLOGS + " blogs.writer_id = " + strconv.Itoa(member_id) + " "
	if _, ok := scope["top_id"]; ok {
		sql += " AND blogs.id > :top_id "
	} else if _, ok := scope["bottom_id"]; ok {
		sql = `
			SELECT * FROM (
		` + sql + `
			AND blogs.id < :bottom_id
			ORDER BY blogs.id DESC
			LIMIT 20
			) AS TMP
			ORDER BY ID ASC
		`
	}
	log.Println(sql)
	if _, err := dbmap.Select(this, sql, scope); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (this *ApiBlogList) SelectAllBetween(dbmap *gorp.DbMap, scope map[string]int) (err error) {
	sql := SQL_SELECT_API_BLOGS + `
			blogs.id BETWEEN :start AND :end
		ORDER BY
			blogs.id
		ASC
		`
	log.Println(sql)
	if _, err := dbmap.Select(this, sql, scope); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

type ApiImage struct {
	Id int32	`db:"id" json:"blog_id"`
	Title string	`db:"title, notnull"                json:"blog_title"`
	Url string	`db:"link_url" json:"blog_url"`
	Images []string	`db:"url" json:"blog_image_url_array"`
	Updated time.Time	`db:"updated, notnull"              json:"blog_update_time"`
}

func SelectMemberImages(dbmap *gorp.DbMap, scope map[string]int, writer_id int) (images []ApiImage, err error) {

	sql := `SELECT
			id,
			title,
			link_url,
			updated
		FROM blogs
		WHERE writer_id = 
	` + strconv.Itoa(writer_id) + " "
	if _, ok := scope["top_id"]; ok {
		sql += " AND id > :top_id "
	} else if _, ok := scope["bottom_id"]; ok {
		sql = `
			SELECT * FROM (
			` + sql + `
			AND id < :bottom_id
			ORDER BY id DESC
			LIMIT 20
			) AS TMP
			ORDER BY id ASC
		`
	}
	log.Println(sql)
	_, err = dbmap.Select(&images, sql, scope)
	if err != nil {
		log.Println(err)
		return
	}
	for i:= 0; i < len(images); i++ {
		/*
		var imgs []Image
		_, err = dbmap.Select( &(imgs), "SELECT * FROM images WHERE article_id = " + strconv.Itoa(int(article.Id)))
		for _, img := range imgs {
			article.Images = append(article.Images, img.Url)
		}
		log.Println(imgs)
		*/
		_, err = dbmap.Select( &(images[i].Images), "SELECT url FROM images WHERE article_id = " + strconv.Itoa(int(images[i].Id)))
		log.Println(images[i].Images)
	}
	return
}
