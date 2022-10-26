package main

import (
	"avito-test-backend/internal"
	"avito-test-backend/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

/*
PUT /deposit/userId:/amount:
PUT /book/userId:/serviceId:/orderId:/amount:
PUT /unBook/orderId:
PUT /withdraw/userID:/serviceId:/orderID:/amount:
GET /balance/userId: - return JSON with int

GET /report/data: - make csv for every service return Url to csv
GET /transactionlist - return json array
GET /trnsactionListByDate
GET /transactionListBySum
*/

/*
user
	-id
	-amount
	-bookamount

order
	-id
	-fromuserid
	-touserid
	-serviceid
	-amount
	-date
	-description
*/

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("error initialization")
	}

	db, err := internal.NewPostgresDB(internal.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	router := gin.Default()
	repo := repository.NewRepository(db)
	service := internal.NewService(repo)
	handler := internal.NewHandler(service)

	router.PUT("/deposit", handler.Deposit)
	router.PUT("/book", handler.Book)
	router.PUT("/unbook", handler.UnBook)
	router.PUT("/withdraw", handler.Withdraw)
	router.GET("/balance/:id", handler.Balance)
	router.GET("/report", handler.Report)
	router.GET("/transactions", handler.Transactions)

	router.Run(viper.GetString("port"))
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
