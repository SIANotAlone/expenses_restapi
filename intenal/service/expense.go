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
func (s *ExpenseService) GetExpense(id int64) *expense.Expense {
	return s.repo.GetExpense(id)
}

func (s *ExpenseService) GetLastExpenses(limit int64, offset int64) []expense.Expense {
	return s.repo.GetLastExpenses(limit, offset)
}

func (s *ExpenseService) GetAllExpenses() []expense.Expense {
	return s.repo.GetAllExpenses()
}
