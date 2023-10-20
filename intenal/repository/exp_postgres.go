package repository

import (
	"expanses_rest_api/intenal/expense"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const (
	expenseSchema = "expenses"
	expenseTable  = "expense"
)

type Exp_Postgres struct {
	db *sqlx.DB
}

func NewExp_Postgres(db *sqlx.DB) *Exp_Postgres {
	return &Exp_Postgres{db: db}
}

func (r *Exp_Postgres) AddExpense(exp expense.Expense) error {
	query := fmt.Sprintf("INSERT INTO %s.%s (type, value, date, comment, description) VALUES ($1, $2, current_date, $3, $4)", expenseSchema, expenseTable)
	r.db.QueryRow(query, exp.Type, exp.Value, exp.Comment, exp.Description)
	return nil
}
func (r *Exp_Postgres) UpdateExpense(exp expense.Expense) error {
	// query := fmt.Sprintf("INSERT INTO %s.%s (type, value, date, comment, description) VALUES ($1, $2, $3, $4, $5)", expenseSchema, expenseTable)
	query := fmt.Sprintf("UPDATE %s.%s SET type=$1, value=$2, comment=$3, description=$4 WHERE id=$5", expenseSchema, expenseTable)
	r.db.QueryRow(query, exp.Type, exp.Value, exp.Comment, exp.Description, exp.Id)

	return nil
}

func (r *Exp_Postgres) GetExpense(id int64) *expense.Expense {
	exp := new(expense.Expense)
	exp.Id = id
	query := fmt.Sprintf("SELECT type, value, date, comment, description FROM expenses.expense WHERE id=$1")
	rows, err := r.db.Query(query, id)
	for rows.Next() {
		err = rows.Scan(&exp.Type, &exp.Value, &exp.Date, &exp.Comment, &exp.Description)
		if err != nil {
			logrus.Fatal(err)
		}
	}

	return exp
}

func (r *Exp_Postgres) GetLastExpenses(limit int64, offset int64) []expense.Expense {
	exp := []expense.Expense{}
	query := fmt.Sprintf("SELECT id, type, value, date, comment, description FROM expenses.expense ORDER BY id DESC LIMIT $1 offset $2")
	rows, err := r.db.Query(query, limit, offset)
	for rows.Next() {
		expense := expense.Expense{}
		err = rows.Scan(&expense.Id, &expense.Type, &expense.Value, &expense.Date, &expense.Comment, &expense.Description)
		if err != nil {
			logrus.Fatal(err)
			return nil
		}
		exp = append(exp, expense)
	}
	return exp
}
func (r *Exp_Postgres) GetAllExpenses() []expense.Expense {
	exp := []expense.Expense{}
	query := fmt.Sprintf("SELECT id, type, value, date, comment, description FROM expenses.expense")
	rows, err := r.db.Query(query)
	for rows.Next() {
		expense := expense.Expense{}
		err = rows.Scan(&expense.Id, &expense.Type, &expense.Value, &expense.Date, &expense.Comment, &expense.Description)
		if err != nil {
			logrus.Fatal(err)
			return nil
		}
		exp = append(exp, expense)
	}
	return exp
}
