package portfolio

import (
	"fmt"
	"github.com/go-pg/pg/orm"
	"portfolio/db"
	// "portfolio/models"
	"testing"
)

func TestSetUpSchema(t *testing.T) {
	models := []interface{}{
		// (*models.Skills)(nil),
		// (*models.Services)(nil),
		// (*models.Experience)(nil),
		// (*models.MessageMe)(nil),
		// (*models.Users)(nil),
	}
	for _, model := range models {
		err := db.DB.Model(model).CreateTable(&orm.CreateTableOptions{})
		if err != nil {
			fmt.Println("could not setup", err)
		}
	}
}
