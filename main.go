package main

import (
	_articleHttpDelivery "backend/article/delivery/http"
	_articleRepo "backend/article/repository/mysql"
	_articleUcase "backend/article/usecase"
	"backend/domain"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware" // Import Echo middleware
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, falling back to system environment variables")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	serverPort := os.Getenv("SERVER_PORT")

	rootDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort)
	tempDB, err := gorm.Open(mysql.Open(rootDsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL server: %v", err)
	}

	err = tempDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`;", dbName)).Error
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&domain.Article{})

	e := echo.New()

	// Add the Logger middleware right here!
	e.Use(middleware.Logger())
	// Optional: Add Recover middleware to prevent panics from crashing the server
	e.Use(middleware.Recover())

	e.Validator = &CustomValidator{validator: validator.New()}

	timeoutContext := time.Duration(2) * time.Second
	ar := _articleRepo.NewMysqlArticleRepository(db)
	au := _articleUcase.NewArticleUsecase(ar, timeoutContext)

	_articleHttpDelivery.NewArticleHandler(e, au)

	log.Printf("Server running on port %s...", serverPort)
	e.Start(serverPort)
}
