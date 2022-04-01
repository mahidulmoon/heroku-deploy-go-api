package models

type Expense struct {
	tablename struct{} `pg:"expense"`
	Id        int64    `pg:"id,pk" json:"id,omitempty"`
	Date      string   `pg:"date" binding: "required" json:"date"`
	Month     string   `pg:"month" binding: "required" json:"month"`
	MonthDate string   `pg:"month_date" binding: "required" json:"month_date"`
	Price     string   `pg:"price" binding: "required" json:"price"`
	Text      string   `pg:"text" binding: "required" json:"text"`
	Time      string   `pg:"time" binding: "required" json:"time"`
	Type      string   `pg:"type" binding: "required" json:"type"`
	UserDate  string   `pg:"userDate" binding: "required" json:"userDate"`
	WeekDay   string   `pg:"week_day" binding: "required" json:"week_day"`
	Year      string   `pg:"year" binding: "required" json:"year"`
}
