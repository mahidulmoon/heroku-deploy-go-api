package models

import "portfolio/db"

type Skills struct {
	tablename struct{} `pg:"skills"`
	Id        int64    `pg:"id,pk" json:"id,omitempty"`
	Name      string   `pg:"name" binding:"required" json:"name"`
	Progress  string   `pg:"progress" binding:"required" json:"progress"`
	Order     string   `pg:"order" json:"order"`
}

func (s *Skills) Add() error {
	_, err := db.DB.Model(s).Insert()

	return err
}

func GetAllSkills() ([]Skills, error) {
	var skills []Skills
	err := db.DB.Model(&skills).Order("order ASC").Select()
	return skills, err
}

func DeleteSkill(id int) error {
	var model Skills
	_, err := db.DB.Model(&model).Where("id = ?", id).Delete()
	return err
}
