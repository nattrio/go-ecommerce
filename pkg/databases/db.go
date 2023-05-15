package databases

import (
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/nattrio/go-ecommerce/config"
)

func DbConnect(cfg config.IDbconfig) *sqlx.DB {

	// Connect
	db, err := sqlx.Connect("pgx", cfg.Url())
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns())
	return db
}
