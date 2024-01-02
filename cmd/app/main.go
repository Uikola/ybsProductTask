package main

import (
	"context"
	"github.com/Uikola/ybsProductTask/internal/config"
	"github.com/Uikola/ybsProductTask/internal/db"
	"github.com/Uikola/ybsProductTask/internal/server"
	sl "github.com/Uikola/ybsProductTask/internal/src/logger"
	"github.com/Uikola/ybsProductTask/pkg/logger"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustConfig()

	log := logger.SetupLogger(cfg.Env)

	log.Info("starting application")

	dataBase := db.InitDB(cfg, log)

	router := chi.NewRouter()
	server.Router(dataBase, router, log)

	log.Info("starting server")

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr:         cfg.Port,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error("failed to start server")
		}
	}()

	log.Info("server started")

	<-done

	log.Info("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("failed to stop server", sl.Err(err))
		return
	}
	defer dataBase.Close()
	log.Info("server stopped")
}
