package internal

import (
	"avito-test-backend/internal/repository"
	"avito-test-backend/internal/structures"
	"encoding/csv"
	"errors"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type service struct {
	r *repository.Repository
}

func NewService(rep *repository.Repository) *service {
	return &service{r: rep}
}

func (s *service) Deposit(fromuserid, userid, orderid, serviceid int, amount int) (structures.User, error) {
	user, err := s.r.GetUser(userid)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return structures.User{}, err
		}
		s.r.CreateUser(userid)
	}
	u, err := s.r.UpdateAmount(userid, user.Amount+amount)
	if err != nil {
		return structures.User{}, err
	}
	_, err = s.r.CreateOrder(fromuserid, userid, serviceid, orderid, amount, time.Now(), "deposit")
	if err != nil {
		return structures.User{}, err
	}
	return u, nil
}

func (s *service) Book(userid int, bookamount int) (structures.User, error) {
	user, err := s.r.GetUser(userid)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return structures.User{}, err
		}
		return structures.User{}, errors.New("insufficient funds")
	}
	if user.Amount < bookamount {
		return structures.User{}, errors.New("insufficient funds")
	}
	u, err := s.r.UpdateBookAmount(userid, user.Amount-bookamount, user.Bookamount+bookamount)
	if err != nil {
		return structures.User{}, err
	}
	return u, nil
}

func (s *service) UnBook(userid int, bookamount int) (structures.User, error) {
	user, err := s.r.GetUser(userid)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return structures.User{}, err
		}
		return structures.User{}, errors.New("insufficient funds")
	}
	if user.Bookamount < bookamount {
		return structures.User{}, errors.New("insufficient funds")
	}
	u, err := s.r.UpdateBookAmount(userid, user.Amount+bookamount, user.Bookamount-bookamount)
	if err != nil {
		return structures.User{}, err
	}
	return u, nil
}

func (s *service) Withdraw(userid, touserId, orederid, serviceid int, amount int) (structures.User, error) {
	user, err := s.r.GetUser(userid)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			s.UnBook(userid, amount)
			return structures.User{}, err
		}
		return structures.User{}, errors.New("insufficient funds")
	}
	if user.Bookamount < amount {
		return structures.User{}, errors.New("insufficient funds")
	}
	u, err := s.r.UpdateBookAmount(userid, user.Amount, user.Bookamount-amount)
	if err != nil {
		s.UnBook(userid, amount)
		return structures.User{}, err
	}
	_, err = s.r.CreateOrder(userid, touserId, serviceid, orederid, amount, time.Now(), "withdraw")
	if err != nil {
		s.UnBook(userid, amount)
		return structures.User{}, err
	}
	return u, nil
}

func (s *service) Balance(userid int) (int, error) {
	user, err := s.r.GetUser(userid)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return 0, err
		}
		return 0, nil
	}
	return user.Amount, nil
}

/*
get from db all Deposit orders from this month
*/
type record struct {
	id     int
	amount int
}

func rightPath(url string) string {
	return strings.ReplaceAll(url, "\\", "/")
}

func (s *service) Report(month int, year int) (string, error) {
	orders, err := s.r.GetMonthOrders(month, year)
	if err != nil {
		return "", err
	}
	services := make(map[int]int)
	for _, order := range orders {
		services[order.Serviceid] += order.Amount
	}
	records := make([]record, 0, len(services))
	for i, v := range services {
		records = append(records, record{id: i, amount: v})
	}
	os.Mkdir("reports", 0755)
	os.Chdir("reports")
	path, _ := os.Getwd()
	defer os.Chdir("..")
	file := "reportmonth" + strconv.Itoa(month) + "year" + strconv.Itoa(year) + ".csv"
	path = filepath.Join(path, file)
	csvFile, err := os.Create(rightPath(file))
	defer csvFile.Close()
	if err != nil {
		return "", err
	}
	csvwriter := csv.NewWriter(csvFile)
	defer csvwriter.Flush()
	var data [][]string
	for _, record := range records {
		row := []string{strconv.Itoa(record.id), strconv.Itoa(record.amount)}
		data = append(data, row)
	}
	csvwriter.WriteAll(data)
	return path, nil
}

func (s *service) Transactions(userid int, sortOrder string) ([]structures.Order, error) {
	if sortOrder != "date" && sortOrder != "amount" && sortOrder != "" {
		return nil, errors.New("incorrect order request")
	}
	orders, err := s.r.GetTransactions(userid, sortOrder)
	if err != nil {
		return nil, err
	}
	for i, _ := range orders {
		orders[i].Id = i + 1
	}
	return orders, nil
}
