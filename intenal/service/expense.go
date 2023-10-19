package service

import (
	"expanses_rest_api/intenal/expense"
	"expanses_rest_api/intenal/repository"
)

type ExpenseService struct {
	repo repository.Expense
}

func NewExpenseService(repo repository.Expense) *ExpenseService {
	return &ExpenseService{repo: repo}
}

func (s *ExpenseService) AddExpense(exp expense.Expense) error {
	return s.repo.AddExpense(exp)
}
func (s *ExpenseService) UpdateExpense(exp expense.Expense) error {
	return s.repo.UpdateExpense(exp)
}
