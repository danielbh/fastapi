package main

import (
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port   int    `envconfig:"port" default:3000`
	APIKEY string `envconfig:"apikey"`
}

func main() {
	startServer(getConfig())
}

func getConfig() *Config {
	var cfg Config
	err := envconfig.Process("funspacestuff", &cfg)

	if err != nil {
		panic(err)
	}

	return &cfg
}

func startServer(config *Config) {
	app := fiber.New()

	client := resty.New()

	app.Get("/", func(c *fiber.Ctx) {
		resp, err := client.R().Get("https://api.nasa.gov/neo/rest/v1/feed?start_date=2015-09-07&end_date=2015-09-08&api_key=" + config.APIKEY)

		if err != nil {
			c.Status(http.StatusInternalServerError)
		} else {
			c.Send(resp)
		}
	})

	app.Listen(config.Port)
}
