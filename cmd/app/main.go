package main

import (
	"context"
	"github.com/Uikola/ybsProductTask/internal/config"
	"github.com/Uikola/ybsProductTask/internal/db"
	"github.com/Uikola/ybsProductTask/internal/server"
	"github.com/Uikola/ybsProductTask/pkg/zlog"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustConfig()

	log := zlog.Default(true, "dev", zerolog.InfoLevel)
	log.Info().Msg("starting application")

	dataBase := db.InitDB(cfg, log)

	router := chi.NewRouter()
	server.Router(dataBase, router, log)

	log.Info().Msg("starting server")

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
			log.Error().Err(err).Msg("failed to start server")
		}
	}()

	log.Info().Msg("server started")

	<-done

	log.Info().Msg("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("failed to stop server")
		return
	}
	defer dataBase.Close()
	log.Info().Msg("server stopped")
}
