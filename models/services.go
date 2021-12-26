package models

import "portfolio/db"

type Services struct {
	tablename  struct{} `pg:"services"`
	Id         int64    `pg:"id,pk" json:"id,omitempty"`
	Name       string   `pg:"name" binding:"required" json:"name"`
	Technology string   `pg:"technology" binding:"required" json:"technology"`
	Github     string   `pg:"github" binding:"required" json:"github"`
	Order      string   `pg:"order" json:"order"`
}

func (s *Services) Add() error {
	_, err := db.DB.Model(s).Insert()
	return err
}

func GetServices() ([]Services, error) {
	var services []Services
	err := db.DB.Model(&services).Select()
	return services, err
}
