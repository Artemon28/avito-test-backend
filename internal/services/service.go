package services

import (
	"avito-test-backend/internal/repository"
	"avito-test-backend/internal/structures"
)

//go:generate mockgen -source=service.go -destination=mocks/service_mock.go

type UserInterface interface {
	Deposit(fromuserid, userid, orderid, serviceid int, amount int) (structures.User, error)
	Book(userid int, bookamount int) (structures.User, error)
	UnBook(userid int, bookamount int) (structures.User, error)
	Withdraw(userid, touserId, orederid, serviceid int, amount int) (structures.User, error)
	Balance(userid int) (int, error)
	Report(month int, year int) (string, error)
	Transactions(userid int, sortOrder string) ([]structures.Order, error)
}

type Service struct {
	UserInterface
}

func NewService(rep *repository.Repository) *Service {
	return &Service{UserInterface: NewUserService(rep)}
}
