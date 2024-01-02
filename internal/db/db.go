package db

import (
	"database/sql"
	"github.com/Uikola/ybsProductTask/internal/config"
	sl "github.com/Uikola/ybsProductTask/internal/src/logger"
	"log/slog"
)

func InitDB(cfg *config.Config, log *slog.Logger) *sql.DB {
	db, err := sql.Open(cfg.DriverName, cfg.ConnString)
	if err != nil {
		log.Info("failed to connect to the database", sl.Err(err))
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Info("failed to ping the database", sl.Err(err))
		return nil
	}
	return db
}
