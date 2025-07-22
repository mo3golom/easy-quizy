package main

import (
	"context"
	"fmt"
	"log"
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

	// Add request logging middleware
	r.Use(func(c *gin.Context) {
		c.Next()
		if c.Writer.Status() == 403 {
			log.Printf("[403 DEBUG] PATH: %s | IP: %s | PlayerID: %s | Origin: %s",
				c.FullPath(),
				c.ClientIP(),
				c.Request.Header.Get("X-Player-Id"),
				c.Request.Header.Get("Origin"),
			)
		}
	})

	// Configure CORS for development and production
	corsConfig := cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Player-ID"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	// Set allowed origins based on environment
	if gin.Mode() == gin.DebugMode {
		// Development: Allow Vite dev server and SSR server
		corsConfig.AllowOrigins = []string{
			"http://localhost:5173", // Vite dev server
			"http://localhost:3000", // SSR server
			"http://localhost:4173", // Vite preview
		}
	} else {
		// Production: More flexible CORS for proxy scenarios
		corsConfig.AllowAllOrigins = true
	}

	r.Use(cors.New(corsConfig))

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
