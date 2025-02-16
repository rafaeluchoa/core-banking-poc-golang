package bootstrap

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	AddRoutes(app *fiber.App)
}

type ApiConfig struct {
	Port string
	Bind string
}

type ApiApp struct {
	app  *fiber.App
	port string
}

func NewApiApp(config *ApiConfig) *ApiApp {

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		EnablePrintRoutes:     true,
	})

	return &ApiApp{
		app:  app,
		port: fmt.Sprintf("%s:%s", config.Bind, config.Port),
	}
}

func (s *ApiApp) Run(done chan error) {
	s.app.Hooks().OnListen(func(l fiber.ListenData) error {
		if fiber.IsChild() {
			return nil
		}

		log.Printf("Loading on %s:%s", l.Host, l.Port)
		done <- nil

		return nil
	})

	err := s.app.Listen(s.port)
	if err != nil {
		done <- err
		log.Panicf("Error on run app: %s\n", err)
	}
}

func (s *ApiApp) AddController(ctr Controller) {
	ctr.AddRoutes(s.app)
}
