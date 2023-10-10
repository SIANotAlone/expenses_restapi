package service

import "expanses_rest_api/intenal/repository"

type Authorization interface {
}

type Expense interface {
}
type ExpenseList interface {
}

type Service struct {
	Authorization
	Expense
	ExpenseList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
