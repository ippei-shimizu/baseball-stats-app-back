package main

import (
	"baseball-stats-app-back/internal/delivery/http"
	"baseball-stats-app-back/internal/repository"
	"baseball-stats-app-back/internal/usecase"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db, err := initDB()
	if err != nil {
		log.Fatal("データベース接続失敗", err)
	}

	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)

	router := gin.Default()
	userHandler := http.NewUserHandler(userUseCase)
	router.POST("/auth/google", userHandler.AuthenticateGoogleUser)

	port := os.Getenv("PORT")
	log.Println("サーバー起動 :", port)
	router.Run(":" + port)
}

func initDB() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("データベース接続成功")
	return db, nil
}
