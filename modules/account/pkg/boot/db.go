package boot

import (
	"database/sql"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

type DBConfig struct {
	URL      string
	User     string
	Password string
}

type DBApp struct {
	config *DBConfig
	db     *sql.DB
}

func NewDBApp(config *DBConfig) *DBApp {
	return &DBApp{
		config: config,
	}
}

func (s *DBApp) Run(done chan error) {
	pgxConfig, err := pgxpool.ParseConfig(s.config.URL)
	if err != nil {
		done <- err
		log.Panicf("Error on parse url: %v", err)
	}

	pgxConfig.ConnConfig.User = s.config.User
	pgxConfig.ConnConfig.Password = s.config.Password

	connStr := stdlib.RegisterConnConfig(pgxConfig.ConnConfig)
	db, err := sql.Open("pgx", connStr)

	if err != nil {
		done <- err
		log.Panicf("Error on connect %v", err)
	}

	err = db.Ping()
	if err != nil {
		done <- err
		log.Panicf("Error pinging db: %v", err)
	}

	s.db = db

	log.Println("DB Connected")
	done <- nil
}

func (s DBApp) GetDB() *sql.DB {
	return s.db
}
