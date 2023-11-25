package config

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5"

	"github.com/evrentan/tasks/db"
)

func GetDbConnection(cfg AppConfig, logger *Logger) *pgx.Conn {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Db.Host, cfg.Db.Port, cfg.Db.Username, cfg.Db.Password, cfg.Db.Database)

	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		logger.Fatalf("Error connecting to the %s database. Error: %v", cfg.Db.Database, err)
	}

	applyDbMigration(cfg, logger)

	return conn
}

func applyDbMigration(cfg AppConfig, logger *Logger) {
	migrationSource, err := iofs.New(db.Migrations, "migrations")
	if err != nil {
		logger.Fatalf("Error creating migration source. Error: %v", err)
	}

	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Db.Username, cfg.Db.Password, cfg.Db.Host, cfg.Db.Port, cfg.Db.Database)

	migration, err := migrate.NewWithSourceInstance("iofs", migrationSource, databaseUrl)

	if err != nil {
		logger.Fatalf("Error creating migration instance. Error: %v", err)
	}

	err = migration.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Fatalf("Error applying migration. Error: %v", err)
	}

	logger.Infof("Migration applied successfully")
}
