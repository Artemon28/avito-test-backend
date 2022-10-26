package repository

import (
	"avito-test-backend/internal/structures"
	"gorm.io/gorm"
	"time"
)

type User interface {
	CreateUser(userid int) (int, error)
	GetUser(id int) (structures.User, error)
	UpdateAmount(userId, amount int) (structures.User, error)
	UpdateBookAmount(userId, amount, bookAmount int) (structures.User, error)
}

type Order interface {
	CreateOrder(fromuserid, touserid, serviceid, orderid, amount int,
		date time.Time, description string) (int, error)
	GetMonthOrders(month int, year int) ([]structures.Order, error)
	GetTransactions(userid int, sortOrder string) ([]structures.Order, error)
}

type Repository struct {
	User
	Order
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:  NewUserPostgres(db),
		Order: NewOrderPostgres(db),
	}
}
