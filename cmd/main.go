package main

import (
	"fmt"
	"loan-app/config"
	"loan-app/internal/app"
	"loan-app/pkg/database/gorm"
	"loan-app/pkg/fiber"
	"loan-app/pkg/logger"
	"loan-app/pkg/middleware"
	"loan-app/pkg/validator"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"

	_ "loan-app/docs/swagger"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// @title Loan App API
// @version 1.0
// @description This is a sample server for a Loan Application.
// @termsOfService http://swagger.io/terms/
// @contact.name Ghalmas Shanditya Putra Agung
// @contact.email ghalmas.shanditya.putra.agung@gmail.com
// @host localhost:3000
// @BasePath /v1
// @schemes http https
// @securityDefinitions.apikey bearer
// @in header
// @name Authorization
// @description Enter the token with the `Bearer ` prefix, e.g. "Bearer abcde12345"
func main() {
	conf := config.Read()

	log := logger.NewLogger(conf)
	log.Info("initialized logger")

	newValidator := validator.NewValidator()
	log.Info("initialized validator")

	// Connect to PostgresSQL under the GORM ORM
	db := gorm.NewGormDB(conf, log)
	log.Infof("database connected, host: %s", conf.Database.Host)
	defer db.Close()

	// Initialize fiber application
	fiberApp := fiber.NewFiber(conf, log)
	log.Infof("initialized fiber server")

	// Setup middleware
	middleware.SetupMiddleware(fiberApp, conf, log)
	log.Infof("setup middleware for fiber server")

	// Bootstrap application
	app.Bootstrap(&app.BootstrapConfig{
		App:       fiberApp,
		DB:        db.DB(),
		Log:       log,
		Config:    conf,
		Validator: newValidator,
	})

	log.Info("setup exception middleware")
	middleware.SetupExceptionMiddleware(fiberApp)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	var serverShutdown sync.WaitGroup
	go func() {
		_ = <-signalChan
		log.Info("Gracefully shutting down...")
		serverShutdown.Add(1)
		defer serverShutdown.Done()
		_ = fiberApp.ShutdownWithTimeout(60 * time.Second)
	}()

	log.Infof("starting server on port %d...", conf.App.Port)
	if err := fiberApp.Listen(fmt.Sprintf("%s:%d", conf.App.Host, conf.App.Port)); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

	serverShutdown.Wait()
	log.Info("Running cleanup tasks...")
}
