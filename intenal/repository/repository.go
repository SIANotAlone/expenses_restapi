package repository

import (
	"expanses_rest_api/intenal/expense"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
}

type Expense interface {
	AddExpense(exp expense.Expense) error
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
