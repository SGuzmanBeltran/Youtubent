package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"youtubent/internal/config"
	"youtubent/pgk/postgres"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	pool, err := postgres.NewDBPool(config.Envs.DBUrl)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Postgres connected successfully")
	}
	defer pool.Close()

	startServices(pool)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		log.Println("Shutting down server...")
		if err := app.Shutdown(); err != nil {
			log.Fatal("Server forced to shutdown:", err)
		}
	}()

	if err := app.Listen(":3000"); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server error:", err)
	}
}

func startServices(dbPool *pgxpool.Pool) {
	// newDB := db.New(dbPool)
}
