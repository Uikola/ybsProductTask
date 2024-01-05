package db

import (
	"database/sql"
	"github.com/Uikola/ybsProductTask/internal/config"
	"github.com/rs/zerolog"
)

func InitDB(cfg *config.Config, log zerolog.Logger) *sql.DB {
	db, err := sql.Open(cfg.DriverName, cfg.ConnString)
	if err != nil {
		log.Info().Err(err).Msg("failed to connect to the database")
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Info().Err(err).Msg("failed to ping the database")
		return nil
	}
	return db
}
