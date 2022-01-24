package models

type Users struct {
	tablename struct{} `pg:"users"`
	Id        int64    `pg:"id,pk" json:"id,omitempty"`
	Name      string   `pg:"name" json:"name"`
	Email     string   `pg:"email" json:"email"`
	Password  string   `pg:"password" json:"password"`
	IsAdmin   string   `pg:"is_admin" json:"is_admin"`
	group     int16    `pg:"group" json:"group"`
}
