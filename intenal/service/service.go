package service

import (
	"expanses_rest_api/intenal/expense"
	"expanses_rest_api/intenal/repository"
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

type Service struct {
	Authorization
	Expense
	ExpenseList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Expense: NewExpenseService(repos.Expense),
	}
}
