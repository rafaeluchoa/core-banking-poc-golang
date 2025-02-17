package bootstrap

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type MigrationConfig struct {
	Dir  string
	Name string
}

type MigrationApp struct {
	config *MigrationConfig
	db     *sql.DB
}

func NewMigrationApp(config *MigrationConfig, db *sql.DB) *MigrationApp {
	return &MigrationApp{
		config: config,
		db:     db,
	}
}

func (s *MigrationApp) Run(done chan error) {
	driver, err := postgres.WithInstance(s.db, &postgres.Config{})
	if err != nil {
		done <- err
		log.Panicf("Error on migration %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		s.config.Dir, s.config.Name, driver)
	if err != nil {
		done <- err
		log.Panicf("Error on migration %v", err)
	}

	err = m.Up()
	if err != nil {
		if err != migrate.ErrNoChange {
			done <- err
			log.Panicf("Error on migration %v", err)
		}
	}

	log.Println("Migration done")
	done <- nil
}
