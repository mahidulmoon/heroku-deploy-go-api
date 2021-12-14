package portfolio

import (
	"fmt"
	"portfolio/models"
	"testing"

	"portfolio/db"

	"github.com/go-pg/pg/orm"
)

func TestSetUpSchema(t *testing.T) {
	models := []interface{}{
		(*models.Skills)(nil),
	}
	for _, model := range models {
		err := db.DB.Model(model).CreateTable(&orm.CreateTableOptions{})
		if err != nil {
			fmt.Println("could not setup", err)
		}
	}
}
