package models

import (
	"portfolio/db"
)

type Expense struct {
	tablename struct{} `pg:"expense"`
	Id        int64    `pg:"id,pk" json:"id,omitempty"`
	Date      string   `pg:"date" binding: "required" json:"date"`
	Month     string   `pg:"month" binding: "required" json:"month"`
	Monthdate string   `pg:"month_date" binding: "required" json:"monthdate"`
	Price     string   `pg:"price" binding: "required" json:"price"`
	Text      string   `pg:"text" binding: "required" json:"text"`
	Time      string   `pg:"time" binding: "required" json:"time"`
	Type      string   `pg:"type" binding: "required" json:"type"`
	UserDate  string   `pg:"userDate" binding: "required" json:"userDate"`
	Weekday   string   `pg:"week_day" binding: "required" json:"weekday"`
	Year      string   `pg:"year" binding: "required" json:"year"`
}

func (e *Expense) AddExpense() error {
	_, err := db.DB.Model(e).Insert()
	return err
}

func GetAllExpense() ([]Expense, error) {
	var exp []Expense
	err := db.DB.Model(&exp).Select()

	return exp, err
}

type FilterExpense struct {
	Year  string `json: "year"`
	Month string `json: "month"`
	Day   string `json: "day"`
}

func GetFilterExpense(year string, month string, day string) ([]Expense, error) {
	var expense []Expense
	var err error
	if year != "" && month != "" && day == "" {
		err = db.DB.Model(&expense).Where("year = ?", year).Where("month = ?", month).Select()
	} else if month == "" {
		err = db.DB.Model(&expense).Where("year = ?", year).Select()
	} else if year == "" {
		err = db.DB.Model(&expense).Where("month = ?", month).Select()
	} else if year != "" && month != "" && day != "" {
		err = db.DB.Model(&expense).Where("year = ?", year).Where("month = ?", month).Where("monthdate = ?", day).Select()
	}

	return expense, err

}

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

func (g *GenerateExpense) GetGenerateInfo() (string, error) {
	var getexpense []GenerateExpense
	err := db.DB.Model(&getexpense).Where("year = ?", g.Year).Where("month = ?", g.Month).Select()
	if err != nil {
		println("Query Error")
		return "no data found", err
	} else {
		if len(getexpense) == 0 {
			var getdata []Expense
			err := db.DB.Model(&getdata).Where("year = ?", g.Year).Where("month = ?", g.Month).Select()
			if err != nil {
				println("Query Error")
				return "no data found", err
			} else {
				return "data found", err
			}
		} else {
			return "already data found", err
		}
	}
}
func (g *GenerateExpense) AddGenExp(amount string) error {
	g.Price = amount
	g.Type = "backup"
	_, err := db.DB.Model(g).Insert()
	return err
}

func AllGeneData() ([]GenerateExpense, error) {
	var genexp []GenerateExpense
	err := db.DB.Model(&genexp).Select()
	return genexp, err

}
