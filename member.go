package main

type Member struct {
	Id int32	`json:"member_id"	db:"id, primarykey, autoincrement"`
	Name string	`json:"member_name"	db:"name"`
}
