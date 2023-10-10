package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Expense interface {
}
type ExpenseList interface {
}

type Repository struct {
	Authorization
	Expense
	ExpenseList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
