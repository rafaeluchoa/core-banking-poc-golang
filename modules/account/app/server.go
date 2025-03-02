package app

import (
	"database/sql"
	"nk/account/internal/ctr"
	"nk/account/pkg/boot"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CONFIG = "config"
)

func Start() {
	Run(".").Wait()
}

func Run(path string) *boot.Launcher {
	s := &server{
		path: path,
		c:    boot.NewContext(),
		l:    boot.NewLauncher(),
	}

	s.registerDB()
	s.registerMigration()

	s.registerBus()
	s.registerMongo()

	s.registerAccount()

	s.registerAPI()

	return s.l
}

type server struct {
	path string
	c    *boot.Context
	l    *boot.Launcher
}

func (s *server) registerDB() {
	db := boot.NewDBApp(
		boot.Load[boot.DBConfig](s.path, CONFIG, "db"),
	)
	s.l.Run(db)

	boot.Register(s.c, func(_ *boot.Context) *sql.DB {
		return db.GetDB()
	})
}

func (s *server) registerMigration() {
	db := boot.Get[sql.DB](s.c)

	migration := boot.NewMigrationApp(
		boot.Load[boot.MigrationConfig](s.path, CONFIG, "migration"),
		db,
		s.path,
	)

	s.l.Run(migration)
}

func (s *server) registerBus() {
	bus := boot.NewEventBus(
		boot.Load[boot.KafkaConfig](s.path, CONFIG, "bus"),
	)

	boot.Register(s.c, func(_ *boot.Context) *boot.EventBus {
		return bus
	})
}

func (s *server) registerMongo() {
	db := boot.NewMongoApp(
		boot.Load[boot.MongoConfig](s.path, CONFIG, "mongo"),
	)
	s.l.Run(db)

	boot.Register(s.c, func(_ *boot.Context) *mongo.Database {
		return db.GetDB()
	})
}

func (s *server) registerAPI() {
	apiApp := boot.NewAPIApp(
		boot.Load[boot.APIConfig](s.path, CONFIG, "api"),
	)

	s.addAuditMid(apiApp)

	apiApp.AddController(boot.Get[ctr.AccountCtr](s.c))

	s.l.Run(apiApp)
}

func (s *server) addAuditMid(apiApp *boot.APIApp) {
	mongo := boot.Get[mongo.Database](s.c)
	boot.NewAudit(apiApp.App, mongo)
}
