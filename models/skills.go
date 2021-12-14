package models

import "portfolio/db"

type Skills struct {
	tablename struct{} `pg:"skills"`
	Id        int64    `pg:"id,pk" json:"id,omitempty"`
	Name      string   `pg:"name" binding:"required" json:"name"`
	Progress  string   `pg:"progress" binding:"required" json:"progress"`
}

func (s *Skills) Add() error {
	_, err := db.DB.Model(s).Insert()

	return err
}
