package service

import (
	"database/sql"
	"fmt"
	"github.com/bemmanue/wildberries_L0/internal/store/sqlstore"
	_ "github.com/lib/pq"
	"net/http"
)

// Start ...
func Start(config *Config) error {
	db, err := newDB(*config.Database)
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)

	srv := newServer(store)

	return http.ListenAndServe(config.BindAddr, srv.router)
}

// newDB ...
func newDB(config DatabaseConfig) (*sql.DB, error) {
	databaseURL := fmt.Sprintf(
		"host=localhost port=%d user=%s dbname=%s sslmode=%s",
		config.Port,
		config.User,
		config.Name,
		config.SSLMode,
	)

	// connect to database
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	// check database connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
