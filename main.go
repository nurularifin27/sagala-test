package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sagala/factory"
	"sagala/internal/routes"
	"sagala/pkg/middleware"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func init() {
	if _, err := os.Stat(".env"); err == nil {
		_ = godotenv.Load(".env")
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func main() {
	port := viper.GetString("app.port")
	if port == "" {
		port = "8080"
	}

	factory := factory.NewFactory()
	app := gin.New()

	app.Use(middleware.RequestID())
	app.Use(middleware.Logger())
	app.Use(middleware.ErrorHandler())

	routes.SetupRoutes(app, factory)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: app,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
