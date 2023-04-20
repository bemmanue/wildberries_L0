package server

import (
	"database/sql"
	"fmt"
	"github.com/bemmanue/wildberries_L0/internal/broker/nats"
	"github.com/bemmanue/wildberries_L0/internal/cache/mapcache"
	"github.com/bemmanue/wildberries_L0/internal/config"
	"github.com/bemmanue/wildberries_L0/internal/store/sqlstore"
	_ "github.com/lib/pq"
	"log"
)

// Start ...
func Start(config *config.Config) error {
	db, err := newDB(*config.Database)
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)
	cache, _ := mapcache.New(store)

	broker, err := nats.New(config.Nuts, store, cache)
	if err != nil {
		log.Fatalln(err)
	}

	if err := broker.Subscribe(config.Nuts.Subject); err != nil {
		log.Fatalln(err)
	}

	srv := newServer(store, cache)

	return srv.router.Run(config.BindAddr)
}

// newDB ...
func newDB(config config.DatabaseConfig) (*sql.DB, error) {
	databaseURL := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.User,
		config.Password,
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
