package models

type MessageMe struct {
	tablename struct{} `pg:"messageme"`
	Id        int64    `pg:"id,pk" json:"id,omitempty"`
	Name      string   `pg:"name" binding:"required" json:"name"`
	Email     string   `pg: "email" binding: "required" json:"email"`
	Phone     string   `pg: "phone" binding: "required" json:"phone"`
	Message   string   `pg: "message" binding: "required" json:"message"`
}
