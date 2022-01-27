package models

import "portfolio/db"

type MessageMe struct {
	tablename struct{} `pg:"messageme"`
	Id        int64    `pg:"id,pk" json:"id,omitempty"`
	Name      string   `pg:"name" binding:"required" json:"name"`
	Email     string   `pg: "email" binding: "required" json:"email"`
	Phone     string   `pg: "phone" binding: "required" json:"phone"`
	Message   string   `pg: "message" binding: "required" json:"message"`
}

func GetAllMessage() ([]MessageMe, error) {
	var msg []MessageMe
	err := db.DB.Model(&msg).Select()
	return msg, err
}

func (m *MessageMe) AddMsg() error {
	_, err := db.DB.Model(m).Insert()
	return err
}
