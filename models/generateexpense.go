package models

type GenerateExpense struct {
	tablename struct{} `pg: "generateexpense"`
	Id        int64    `pg: "id,pk" json: "id,omitempty"`
	Date      string   `pg:"date" binding: "required" json:"date"`
	Month     string   `pg:"month" binding: "required" json:"month"`
	Monthdate string   `pg:"month_date" binding: "required" json:"monthdate"`
	Price     string   `pg:"price" binding: "required" json:"price"`
	Time      string   `pg:"time" binding: "required" json:"time"`
	Type      string   `pg:"type" binding: "required" json:"type"`
	UserDate  string   `pg:"userDate" binding: "required" json:"userDate"`
	Weekday   string   `pg:"week_day" binding: "required" json:"weekday"`
	Year      string   `pg:"year" binding: "required" json:"year"`
}
