package postgres

import (
	"fmt"

	"github.com/Trendyol/post_service/configs"
	"github.com/jmoiron/sqlx"

	// _ "github.com/go-sql-driver/postgres"
	_ "github.com/lib/pq"
)

func NewClient(cfg configs.Config) (*sqlx.DB, error) {
	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)
	connDb, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		return nil, err
	}
	return connDb, nil

}
