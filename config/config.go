package config

import (
	"TDD-GoAPI/controller"
	"TDD-GoAPI/repository"
	"TDD-GoAPI/service"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

type Config struct {
	routes *controller.Controller
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
func ConnectToDatabase() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"))
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	fmt.Printf("Successfully connected to %v", os.Getenv("DB_HOST"))
	return nil
}

func MigrateDatabase() error {
	repo := repository.New(DB)

	if err := repo.Book.Migrate(); err != nil {
		return err
	}
	return nil
}

func Echo() (*echo.Echo, error) {
	e := echo.New()
	r := repository.New(DB)
	s := service.New(r)
	c := controller.New(s)
	controller.SetupRoutes(e, c)
	return e, nil
}

func StartServer() error {
	if err := ConnectToDatabase(); err != nil {
		return err
	}
	if err := MigrateDatabase(); err != nil {
		return err
	}
	e, err := Echo()
	if err != nil {
		return fmt.Errorf("failed to initialize Echo: %v", err)
	}
	if err := e.Start(":8080" + os.Getenv("PORT")); err != nil {
		return err
	}
	return nil
}
