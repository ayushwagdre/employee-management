package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"

	"practice/api/routes"
	"practice/config"
	"practice/lib/db"
	"practice/lib/logger"
)

var ctx context.Context

const (
	reqReadTimeout  = 10 * time.Second
	reqWriteTimeout = 10 * time.Second
	reqIdleTimeout  = 60 * time.Second
)

func panicHandler() {
	if r := recover(); r != nil {
		err, _ := r.(error)
		logger.E(context.Background(), err, "[Employee] Panic Exception", zap.Any("recover", r))
	}
}

func init() {
	config := config.NewConfig()
	ctx = context.Background()

	logLevel := logger.INFO

	logger.Init(logLevel, config.Env)

	dbClient := db.NewDB(&db.DBOpts{
		URL:                   config.DB.URL,
		MaxIdleConnection:     config.DB.MaxIdleConnections,
		MaxActiveConnection:   config.DB.MaxOpenConnections,
		MaxConnectionLifetime: time.Hour,
		DriverName:            db.PostgresDriver,
	})

	err := dbClient.Connect()
	if err != nil {
		panic(fmt.Sprintf("DB initialization failed: %v", err))
	}

}

func main() {
	defer panicHandler()
	defer db.Get().Close()

	config := config.NewConfig()
	router := httprouter.New()

	routes.Init(router)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Port),
		WriteTimeout: reqWriteTimeout,
		ReadTimeout:  reqReadTimeout,
		IdleTimeout:  reqIdleTimeout,
		Handler:      router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				logger.E(ctx, err, "server failed to start")
			}
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done
	logger.I(ctx, "server shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.I(ctx, "failed shutdown gracefully")
	}
}
