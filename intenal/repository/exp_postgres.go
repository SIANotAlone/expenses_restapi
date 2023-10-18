package repository

import (
	"expanses_rest_api/intenal/expense"
	"fmt"

	"github.com/jmoiron/sqlx"
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
	query := fmt.Sprintf("INSERT INTO %s.%s (type, value, date, comment, description) VALUES ($1, $2, $3, $4, $5)", expenseSchema, expenseTable)
	r.db.QueryRow(query, exp.Type, exp.Value, exp.Date, exp.Comment, exp.Description)
	return nil
}
