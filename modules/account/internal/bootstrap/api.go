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

	port := fmt.Sprintf("%s:%s", config.Bind, config.Port)
	log.Printf("Loading on %s", port)

	return &ApiApp{
		app:  app,
		port: port,
	}
}

func (s *ApiApp) Run() error {
	return s.app.Listen(s.port)
}

func (s *ApiApp) AddController(ctr Controller) {
	ctr.AddRoutes(s.app)
}
