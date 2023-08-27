package handler

import (
	"github.com/Aorts/PieFireDire/internal/httputil"
	"github.com/gofiber/fiber/v2"
	"regexp"
	"strings"
)

type TextFile struct {
	Contents []byte
}

type WordCount map[string]int

type Response struct {
	Beef WordCount `json:"beef"`
}

func GetBeefHandler(getMeat GetMeatFunc, CountAllMeat CountAllMeatFunc) fiber.Handler {
	return func(c *fiber.Ctx) error {
		meatText, err := getMeat()
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Get Meat Fail")
		}
		meatList := CountAllMeat(string(meatText.Contents))

		return c.Status(fiber.StatusOK).JSON(Response{
			Beef: meatList,
		})
	}
}

type GetMeatFunc func() (res *TextFile, err error)

func GetMeat(caller httputil.HTTPGetFunc) GetMeatFunc {
	return func() (res *TextFile, err error) {
		value, err := caller()
		if err != nil {
			return nil, err
		}

		textFile := TextFile{
			Contents: value,
		}

		return &textFile, nil
	}
}

type CountAllMeatFunc func(text string) WordCount

func CountAllMeat() CountAllMeatFunc {
	return func(text string) WordCount {
		wordCounts := make(WordCount)
		splitter := regexp.MustCompile(`[ .,]+`)

		lines := strings.Split(text, "\n")
		for _, line := range lines {
			words := splitter.Split(line, -1)
			for _, word := range words {
				if len(word) > 0 {
					wordCounts[word]++
				}
			}
		}
		return wordCounts
	}
}
