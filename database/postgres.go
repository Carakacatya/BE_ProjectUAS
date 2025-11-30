package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"projectuasbe/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Postgres *pgxpool.Pool

func InitPostgres() {
	cfg := config.AppConfig

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
	)

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("❌ Failed parsing PostgreSQL config: %v", err)
	}

	poolConfig.MaxConns = 10
	poolConfig.MinConns = 2
	poolConfig.MaxConnIdleTime = 5 * time.Minute

	conn, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatalf("❌ Failed connect PostgreSQL: %v", err)
	}

	// test ping
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatalf("❌ PostgreSQL unreachable: %v", err)
	}

	Postgres = conn
	log.Println("✅ PostgreSQL connected successfully")
}
