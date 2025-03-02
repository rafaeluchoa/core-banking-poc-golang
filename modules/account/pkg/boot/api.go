package boot

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "nk/account/docs" // swagger docs

	"github.com/gofiber/swagger"
)

type Controller interface {
	AddRoutes(app *fiber.App)
}

type APIConfig struct {
	Port string
	Bind string
}

type APIApp struct {
	App  *fiber.App
	port string
}

func NewAPIApp(config *APIConfig) *APIApp {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		EnablePrintRoutes:     true,
	})

	app.Use("/api/*", logger.New())
	app.Get("/swagger/*", swagger.HandlerDefault)

	return &APIApp{
		App:  app,
		port: fmt.Sprintf("%s:%s", config.Bind, config.Port),
	}
}

func (s *APIApp) Run(done chan error) {
	s.App.Hooks().OnListen(func(l fiber.ListenData) error {
		if fiber.IsChild() {
			return nil
		}

		log.Printf("Listening on %s:%s", l.Host, l.Port)
		done <- nil

		return nil
	})

	err := s.App.Listen(s.port)
	if err != nil {
		done <- err
	}
}

func (s *APIApp) AddController(ctr Controller) {
	ctr.AddRoutes(s.App)
}
