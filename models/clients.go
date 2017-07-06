package models

import(
	"github.com/go-gorp/gorp"
)

type Client struct {
	Id int32		`db:"id, primarykey, autoincrement"`
	Token string		`db:"token, notnull"`
}

// Get target clients from database.
func GetClientsFromDB(dbmap *gorp.DbMap, requirements string) ([]Client, error) {
				var clients []Client
				if _, err := dbmap.Select(&clients, "SELECT * FROM clients " + requirements); err != nil {
								return nil, err
				}
				return clients, nil
}
