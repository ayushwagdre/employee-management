package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/newrelic/go-agent/v3/integrations/nrhttprouter"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/zap"

	"practice/api/routes"
	"practice/config"
	"practice/lib/cache"
	"practice/lib/db"
	"practice/lib/logger"
)

var ctx context.Context

const (
	cacheSize       = 10 * 1000 * 1000 * 2
	reqReadTimeout  = 10 * time.Second
	reqWriteTimeout = 10 * time.Second
	reqIdleTimeout  = 60 * time.Second

	redisMaxIdleConnection     = 20
	redisMaxActiveConnection   = 20
	redisIdleConnectionTimeout = 30 * time.Second
	redisMaxConnectionLifetime = 24 * time.Hour
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

	cache.InitRedisCache(&cache.RedisCacheOpts{
		DB:                    config.Redis.DB,
		Host:                  config.Redis.Host,
		MaxIdleConnection:     redisMaxIdleConnection,
		MaxActiveConnection:   redisMaxActiveConnection,
		IdleConnectionTimeout: redisIdleConnectionTimeout,
		MaxConnectionLifetime: redisMaxConnectionLifetime,
	})

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
	defer cache.CloseRedisCache()
	defer db.Get().Close()

	config := config.NewConfig()
	newRelicApp, _ := newrelic.NewApplication()
	router := nrhttprouter.New(newRelicApp)

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
