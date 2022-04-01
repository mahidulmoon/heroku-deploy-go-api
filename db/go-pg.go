package db

import (
	"github.com/go-pg/pg"
)

var DB *pg.DB

func init() {
	db_string := "postgres://xzlhwqxbuzodhw:c8d7ed40d70c2aecd15df0ba655474b9d59441ac86b7c6120d857846612f70c3@ec2-3-218-171-44.compute-1.amazonaws.com:5432/d6kum9sgkhni7f"
	// db_string := "postgres://postgres:1234@localhost:5432/adminpanel?sslmode=disable"
	opt, err := pg.ParseURL(db_string)
	if err != nil {
		panic(err)
	}

	DB = pg.Connect(opt)
}
