package repository

import (
	"avito-test-backend/internal/structures"
	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(userid int) (int, error) {
	var user structures.User
	user.Id = userid
	result := r.db.Create(&user)

	if result.Error != nil {
		return 0, result.Error
	}
	return user.Id, nil
}

func (r *UserPostgres) GetUser(userId int) (structures.User, error) {
	var user structures.User
	err := r.db.First(&user, userId).Error
	if err != nil {
		return structures.User{}, err
	}
	return user, nil
}

func (r *UserPostgres) UpdateAmount(userId, amount int) (structures.User, error) {
	user := structures.User{Id: userId}
	err := r.db.Model(&user).Update("amount", amount).Error
	if err != nil {
		return structures.User{}, err
	}
	return user, nil
}

func (r *UserPostgres) UpdateBookAmount(userId, amount, bookAmount int) (structures.User, error) {
	user := structures.User{Id: userId}
	err := r.db.Model(&user).Updates(map[string]interface{}{"amount": amount, "bookamount": bookAmount}).Error
	if err != nil {
		return structures.User{}, err
	}
	return user, nil
}
