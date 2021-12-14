package portfolio

import (
	"fmt"
	"testing"

	"portfolio/db"

	"github.com/go-pg/pg/orm"
)

func TestSetUpSchema(t *testing.T) {
	models := []interface{}{
		// (*models.Skills)(nil),
		// (*models.Services)(nil),
	}
	for _, model := range models {
		err := db.DB.Model(model).CreateTable(&orm.CreateTableOptions{})
		if err != nil {
			fmt.Println("could not setup", err)
		}
	}
}
