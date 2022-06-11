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
	UserId    int64    `pg:"user_id" json:"user_id"`
}

func (e *Expense) AddExpense(user_id int64) error {
	e.UserId = user_id
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

func GetFilterExpense(year string, month string, day string, user_id int64) ([]Expense, error) {
	var expense []Expense
	var err error
	if year != "" && month != "" && day == "" {
		err = db.DB.Model(&expense).Where("year = ?", year).Where("month = ?", month).Where("user_id = ?", user_id).Select()
	} else if month == "" {
		err = db.DB.Model(&expense).Where("year = ?", year).Where("user_id = ?", user_id).Select()
	} else if year == "" {
		err = db.DB.Model(&expense).Where("month = ?", month).Where("user_id = ?", user_id).Select()
	} else if year != "" && month != "" && day != "" {
		err = db.DB.Model(&expense).Where("year = ?", year).Where("month = ?", month).Where("monthdate = ?", day).Where("user_id = ?", user_id).Select()
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
	Income    string   `pg:"income" binding: "required" json:"income"`
	Time      string   `pg:"time" binding: "required" json:"time"`
	Type      string   `pg:"type" binding: "required" json:"type"`
	UserDate  string   `pg:"userDate" binding: "required" json:"userDate"`
	Weekday   string   `pg:"week_day" binding: "required" json:"weekday"`
	Year      string   `pg:"year" binding: "required" json:"year"`
	UserId    int64    `pg:"user_id" json:"user_id"`
}

func (g *GenerateExpense) GetGenerateInfo(user_id int64) (string, error) {
	var getexpense []GenerateExpense
	err := db.DB.Model(&getexpense).Where("year = ?", g.Year).Where("month = ?", g.Month).Where("user_id = ?", user_id).Select()
	if err != nil {
		println("Query Error")
		return "no data found", err
	} else {
		if len(getexpense) == 0 {
			var getdata []Expense
			err := db.DB.Model(&getdata).Where("year = ?", g.Year).Where("month = ?", g.Month).Where("user_id = ?", user_id).Select()
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
func (g *GenerateExpense) AddGenExp(amount string, income string, user_id int64) error {
	g.Price = amount
	g.Income = income
	g.Type = "backup"
	g.UserId = user_id
	_, err := db.DB.Model(g).Insert()
	return err
}

func AllGeneData(user_id int64) ([]GenerateExpense, error) {
	var genexp []GenerateExpense
	err := db.DB.Model(&genexp).Where("user_id = ?", user_id).Select()
	return genexp, err

}

func DeleteGenExp(year string, month string, user_id int64) error {
	var model []Expense
	_, err := db.DB.Model(&model).Where("year = ?", year).Where("month = ?", month).Where("user_id = ?", user_id).Delete()
	return err
}
