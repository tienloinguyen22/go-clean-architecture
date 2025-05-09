package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/tienloinguyen22/go-clean-architecture/internal/domain/service"
	"github.com/tienloinguyen22/go-clean-architecture/internal/infrastructure/configs"
	"github.com/tienloinguyen22/go-clean-architecture/internal/infrastructure/database"
	"github.com/tienloinguyen22/go-clean-architecture/internal/infrastructure/event"
	"github.com/tienloinguyen22/go-clean-architecture/internal/interface/api"
)

func main() {
	// Init configs
	configs := configs.InitAppConfigs()
	fmt.Printf("Welcome to Go Clean Architecture Project!\n")
	fmt.Printf("Postgres Host: %s\n", configs.PostgresConfig.Host)
	fmt.Printf("Postgres Port: %v\n", configs.PostgresConfig.Port)
	fmt.Printf("Redis Host: %s\n", configs.RedisConfig.Host)
	fmt.Printf("Redis Port: %v\n", configs.RedisConfig.Port)
	fmt.Printf("Server running on port: %v\n", configs.Port)

	// Setup database
	db, err := database.NewPostgresDB(&database.PostgresConfig{
		Host:     configs.PostgresConfig.Host,
		Port:     configs.PostgresConfig.Port,
		Username: configs.PostgresConfig.Username,
		Password: configs.PostgresConfig.Password,
		DBName:   configs.PostgresConfig.DBName,
		SSLMode:  configs.PostgresConfig.SSLMode,
	})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	defer db.Close()

	// Setup pubsub
	ps, err := event.NewPubSub(&event.RedisConfig{
		Host:     configs.RedisConfig.Host,
		Port:     configs.RedisConfig.Port,
		Password: configs.RedisConfig.Password,
		DB:       configs.RedisConfig.DB,
	})
	if err != nil {
		log.Fatal("Failed to connect to redis", err)
	}
	defer ps.Close()

	// Setup repositories
	userRepository := database.NewUserRepository(db)

	// Setup services
	userService := service.NewUserService(userRepository, ps)

	// Setup handlers
	healthHandler := api.NewHeathHandler()
	userHandler := api.NewUserHandler(userService)

	// Setup router
	router := chi.NewRouter()
	router.Get("/health", healthHandler.HandleHealthCheck)
	router.Route("/api/v1", func(r chi.Router) {
		r.Post("/users", userHandler.HandleCreateUser)
		r.Get("/users/{id}", userHandler.HandleGetUserByID)
		r.Put("/users/{id}", userHandler.HandleUpdateUser)
		r.Delete("/users/{id}", userHandler.HandleDeleteUser)
	})

	// Start server
	server := &http.Server{
		Addr:    ":" + strconv.Itoa(configs.Port),
		Handler: router,
	}
	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-stop
		log.Println("Shutting down server...")
		shutdownCtx, _ := context.WithTimeout(context.Background(), 20*time.Second)
		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("Server shutdown timeout. Force exit")
			}
		}()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Fatal("Failed to shutdown server", err)
		}
		log.Println("Server stopped")
		serverStopCtx()
	}()

	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal("Failed to start server", err)
	}
	<-serverCtx.Done()
}
