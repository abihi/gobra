package main

import (
	"bytes"
	"fmt"
	httpserver "gobra/http"
	"gobra/match"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
)

var (
	httpAddress = envString("HTTP_ADDRESS", ":80")

	postgresUser     = envString("POSTGRES_USER", "gobra_service")
	postgresPassword = envString("POSTGRES_PASSWORD", "")
	postgresHost     = envString("POSTGRES_HOST", "postgres")
	postgresPort     = envString("POSTGRES_PORT", "5432")
	postgresDB       = envString("POSTGRES_DB", "gobra_service")
	postgresSSL      = envString("POSTGRES_SSL", "disable")

	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
)

func main() {
	// Postgres setup.
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		postgresUser,
		postgresPassword,
		postgresHost,
		postgresPort,
		postgresDB,
		postgresSSL,
	)

	fmt.Println("dsn: ", dsn)

	// Postgres repository
	// matchRepository := postgres.matchRepository{
	// 	Connector: postgresConnector,
	// }

	// Services
	matchService := match.Service{
		// MatchRepository: matchRepository,
	}

	errorChannel := make(chan error)

	// HTTP transport.
	go func() {
		httpServer := httpserver.Server{
			Address:      httpAddress,
			Logger:       logger,
			MatchService: matchService,
			Timeout:      10 * time.Second,
		}
		errorChannel <- httpServer.Open()
	}()

	// Capture interrupts.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errorChannel <- fmt.Errorf("got signal: %s", <-c)
	}()

	// Wait for any error.
	if err := <-errorChannel; err != nil {
		logger.Fatal(err)
	}
}

func envString(key string, fallback string) string {
	if value, ok := syscall.Getenv(key); ok {
		return value
	}
	return fallback
}
