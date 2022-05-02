package models

import "portfolio/db"

type Experience struct {
	tablename  struct{} `pg:"experience"`
	Id         int64    `pg:"id,pk" json:"id,omitempty"`
	Name       string   `pg:"name" binding:"required" json:"name"`
	Technology string   `pg:"technology" binding:"required" json:"technology"`
	Company    string   `pg:"company" binding:"required" json:"company"`
	StartDate  string   `pg:"start_date" binding:"required" json:"start_date"`
	EndDate    string   `pg:"end_date" binding:"required" json:"end_date"`
	Order      string   `pg:"order" json:"order"`
	UserId     int64    `pg:"user_id" json:"user_id"`
}

func (e *Experience) Add(user_id int64) error {
	e.UserId = user_id
	_, err := db.DB.Model(e).Insert()
	return err
}

func GetAllExperience() ([]Experience, error) {
	var exp []Experience
	err := db.DB.Model(&exp).Order("order ASC").Select()
	return exp, err
}

func GetAllExperienceByUser(user int64) ([]Experience, error) {
	var exp []Experience
	err := db.DB.Model(&exp).Order("order ASC").Where("user_id = ?", user).Select()
	return exp, err
}
