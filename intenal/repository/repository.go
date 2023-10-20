package repository

import (
	"expanses_rest_api/intenal/expense"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
}

type Expense interface {
	AddExpense(exp expense.Expense) error
	UpdateExpense(exp expense.Expense) error
	GetExpense(id int64) *expense.Expense
	GetLastExpenses(limit int64, offset int64) []expense.Expense
	GetAllExpenses() []expense.Expense
}
type ExpenseList interface {
}

type Repository struct {
	Authorization
	Expense
	ExpenseList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Expense: NewExp_Postgres(db),
	}
}
