package store

import (
	"database/sql"

	_ "github.com/lib/pq" // драйвер для database/sql под postgress (анонимный)
)

type Store struct {
	config *Config
	db     *sql.DB
}

func NewDb(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {

	// db, err := sql.Open("postgres", "user=test password=test dbname=test sslmode=disable")
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	// пингуем, для того чтобы бд создало реальное соединение
	// а не ленивое
	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}