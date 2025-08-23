package database

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/Trypion/ecommerce/order-service/internal/config"
)

func NewConnection(cfg *config.Config) (*gorm.DB, error) {
	// Build DSN
	dsn := buildDSN(cfg)

	// Run migrations first
	if err := runMigrations(dsn); err != nil {
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	// Connect with GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	log.Println("Database connection established successfully")
	return db, nil
}

func runMigrations(databaseURL string) error {
	// Get migrations path
	migrationsPath := getMigrationsPath()

	// Create migrator
	migrator, err := NewMigrator(databaseURL, migrationsPath)
	if err != nil {
		return err
	}
	defer migrator.Close()

	// Run migrations
	return migrator.Up()
}

func getMigrationsPath() string {
	// Get current file path
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "./internal/database/migrations"
	}

	// Get directory of current file and append migrations
	dir := filepath.Dir(filename)
	return filepath.Join(dir, "migrations")
}

func buildDSN(cfg *config.Config) string {
	if cfg.DatabaseURL != "" {
		return cfg.DatabaseURL
	}
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)
}
