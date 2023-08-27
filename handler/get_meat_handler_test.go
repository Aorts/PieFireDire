package handler_test

import (
	"fmt"
	handler2 "github.com/Aorts/PieFireDire/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func initGetMeat(app *fiber.App) *http.Response {
	var newReq *http.Request
	newReq = httptest.NewRequest(fiber.MethodPost, "/getMeat", nil)
	newReq.Header.Add("Content-Type", "application/json")
	newReq.Header.Add("Accept-Language", "en-EN")
	res, _ := app.Test(newReq)
	return res
}
func TestGetMeat(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		app := fiber.New()

		handler := handler2.GetBeefHandler(func() (*handler2.TextFile, error) {
			return &handler2.TextFile{}, nil
		}, func(text string) (res handler2.WordCount) {
			return res
		})
		app.Post("/getMeat", handler)

		assert.Equal(t, fiber.StatusOK, initGetMeat(app).StatusCode)
	})
	t.Run("err", func(t *testing.T) {
		app := fiber.New()

		handler := handler2.GetBeefHandler(func() (*handler2.TextFile, error) {
			return &handler2.TextFile{}, fmt.Errorf("something")
		}, func(text string) (res handler2.WordCount) {
			return res
		})
		app.Post("/getMeat", handler)

		assert.Equal(t, fiber.StatusInternalServerError, initGetMeat(app).StatusCode)
	})
}
