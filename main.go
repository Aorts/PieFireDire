package main

import (
	"github.com/Aorts/PieFireDire/config"
	"github.com/Aorts/PieFireDire/handler"
	"github.com/Aorts/PieFireDire/internal/httputil"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"log"
)

func main() {

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(errors.Wrap(err, "Unable to initial config."))
	}

	httpReq := httputil.InitHttpClient(
		cfg.HTTP.TimeOut,
		cfg.HTTP.MaxIdleConn,
		cfg.HTTP.MaxIdleConnPerHost,
		cfg.HTTP.MaxConnPerHost,
	)
	app := fiber.New()
	app.Get("/getMeat", handler.GetBeefHandler(handler.GetMeat(httputil.NewHttpGet(httpReq)), handler.CountAllMeat()))
	app.Get("/hello", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello world")
	})
	app.Listen(":8080")

	// ########################################################################################################################################################

}
