package models

import (
	"portfolio/db"
)

type Services struct {
	tablename  struct{} `pg:"services"`
	Id         int64    `pg:"id,pk" json:"id,omitempty"`
	Name       string   `pg:"name" binding:"required" json:"name"`
	Technology string   `pg:"technology" binding:"required" json:"technology"`
	Github     string   `pg:"github" binding:"required" json:"github"`
	Order      string   `pg:"order" json:"order"`
	UserId     int64    `pg:"user_id"`
}

func (s *Services) Add(user_id int64) error {
	s.UserId = user_id
	_, err := db.DB.Model(s).Insert()
	return err
}

func GetServices() ([]Services, error) {
	var services []Services
	err := db.DB.Model(&services).Order("order ASC").Select()
	return services, err
}
