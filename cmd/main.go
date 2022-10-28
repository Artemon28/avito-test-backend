package main

import (
	_ "avito-test-backend/docs"
	"avito-test-backend/internal"
	"avito-test-backend/internal/repository"
	"avito-test-backend/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
)

// @title           Avito Backend Test Task
// @version         1.0
// @description		Rest API microservice that allowed to manage users bank accounts and create reports for accounting

// @contact.name   Artemiy Chaykov

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	repo := repository.NewRepository(db)
	service := services.NewService(repo)
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
