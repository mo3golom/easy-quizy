package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	txmanager "github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	gameAPI "easy-quizy/api/v1/game"
	gameRepo "easy-quizy/internal/repositories/game"
	gameUC "easy-quizy/internal/usecase/game"
)

func main() {
	if _, err := os.Stat(".env"); err == nil {
		// path/to/whatever exists
		err := godotenv.Load(".env")
		if err != nil {
			panic(err)
		}
	}

	ctx := context.Background()

	// Get SERVER_PORT from env, fallback to 8080
	port := os.Getenv("SERVER_PORT")
	if _, err := strconv.Atoi(port); err != nil {
		port = "8080"
	}

	source := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		loadEnvValue("DB_USER"),
		loadEnvValue("DB_PASSWORD"),
		loadEnvValue("DB_HOST"),
		loadEnvValue("DB_PORT"),
		loadEnvValue("DB_NAME"),
		loadEnvValue("DB_SSL"),
	)

	db, err := sqlx.ConnectContext(
		ctx,
		"postgres",
		source,
	)
	if err != nil {
		panic(err)
	}

	trm := txmanager.Must(trmsqlx.NewDefaultFactory(db))
	trmsqlxGetter := trmsqlx.DefaultCtxGetter

	repository := gameRepo.NewRepository(db, trmsqlxGetter)
	usecase := gameUC.NewUsecase(repository, trm)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	gameHandler := gameAPI.NewHandler(usecase)
	gameHandler.Register(&r.RouterGroup)

	runErr := r.Run(":" + port)
	if runErr != nil {
		logrus.Fatalf("Failed to start server: %v", runErr)
	}
}

func loadEnvValue(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic("env value doesn't exists")
	}

	return value
}
