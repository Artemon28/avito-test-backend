package repository

import (
	"avito-test-backend/internal/structures"
	"gorm.io/gorm"
	"time"
)

type OrderPostgres struct {
	db *gorm.DB
}

/*
type Order struct {
	Id          int    `json:"id"`
	Fromuserid  int    `json:"fromuserid"`
	Touserid    int    `json:"touserid"`
	Serviceid   int    `json:"serviceid"`
	orderid   int    `json:"serviceid"`
	Amount      int    `json:"amount"`
	Date        int    `json:"date"`
	Description string `json:"description"`
}
*/

func NewOrderPostgres(db *gorm.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) CreateOrder(fromuserid, touserid, serviceid, orderid, amount int,
	date time.Time, description string) (int, error) {
	order := structures.Order{
		Fromuserid:  fromuserid,
		Touserid:    touserid,
		Serviceid:   serviceid,
		Orderid:     orderid,
		Amount:      amount,
		Date:        date,
		Description: description,
	}
	err := r.db.Create(&order).Error

	if err != nil {
		return 0, err
	}
	return order.Id, nil
}

func (r *OrderPostgres) GetMonthOrders(month int, year int) ([]structures.Order, error) {
	var orders []structures.Order
	err := r.db.Where(`EXTRACT(MONTH FROM date) = ? AND EXTRACT(YEAR FROM date) = ? AND description = 'deposit'`, month, year).
		Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderPostgres) GetTransactions(userid int, sortOrder string) ([]structures.Order, error) {
	var orders []structures.Order
	err := r.db.Order(sortOrder).Where(`fromuserid = ? OR touserid =?`, userid, userid).
		Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
