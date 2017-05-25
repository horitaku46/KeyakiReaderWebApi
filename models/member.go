package models

//import ( "time" )

type Member struct {
	Id int32		`db:"id, primarykey, autoincrement" json:"member_id"`
	Thumbnail string		`db:"thumbnail, notnull"    json:"thmbnail_url"`
	Name string		`db:"name, notnull"                 json:"member_name"`
	Ruby string		`db:"ruby, notnull"                 json:"member_furigana"`
	/*
	BirthDay time.Time	`db:"birthday, notnull"		    json:"member_birthday"`
	Height int32
	Weight int32
	BloodType string
	Constellatoin string
	Hometown string
	MsgUrl string
	Team int32		`db:"team, notnull"		    json:""`
	*/
}
