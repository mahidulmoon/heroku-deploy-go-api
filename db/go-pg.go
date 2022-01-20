package db

import (
	"github.com/go-pg/pg"
)

var DB *pg.DB

func init() {
	// db_string := "postgres://wcrhovsjeqfkeb:f437021135d05e7792cd7317b4a4e7a74b35ea1ce8e021cd8d71a3564f3cba49@ec2-52-54-38-229.compute-1.amazonaws.com:5432/dcmuducnduvq6e"
	// db_string := "postgres://gwpmkgjmksxdde:036bde826a1cc9b0ee49c8db92d2379cb5362bb3a1ae22a4dce1cb6ce629fe1e@ec2-54-204-99-176.compute-1.amazonaws.com:5432/dcrl4mi3tkc8r8"
	db_string := "postgres://postgres:1234@localhost:5432/adminpanel?sslmode=disable"
	opt, err := pg.ParseURL(db_string)
	if err != nil {
		panic(err)
	}

	DB = pg.Connect(opt)
}
