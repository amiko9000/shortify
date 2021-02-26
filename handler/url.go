package handler

import (
	"github.com/gofiber/fiber/v2"
	"shortify/database"
	"shortify/model"
	"shortify/util"
)

func GetOriginalURL(c *fiber.Ctx) error {
	token := c.Params("token")
	db := database.DB

	url := new(model.URL)
	db.First(&url, "token = ?", token)
	if url.RedirectURL == "" {
		return c.SendStatus(fiber.StatusNotFound)
	}

	type Response struct {
		OriginalURL string `json:"original_url"`
		RedirectURL string `json:"redirect_url"`
	}

	res := Response{
		OriginalURL: c.BaseURL() + "/s/" + token,
		RedirectURL: url.RedirectURL,
	}
	return c.JSON(&res)
}

func GetAllURLs(c *fiber.Ctx) error {
	db := database.DB
	var urls []model.URL
	db.Find(&urls)
	return c.JSON(&urls)
}

func CreateURL(c *fiber.Ctx) error {
	type Input struct {
		URL string `json:"url"`
	}

	input := new(Input)
	if err := c.BodyParser(input); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	url := model.URL{
		Token: util.GenerateToken(),
		RedirectURL: input.URL,
	}

	db := database.DB
	db.Create(&url)
	return c.JSON(&url)
}

func GetURL(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	url := new(model.URL)
	db := database.DB
	// TODO
	db.First(&url, "id = ?", uuid)
	return c.JSON(&url)
}

func DeleteURL(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	db := database.DB
	// TODO
	db.Delete(&model.URL{}, "id = ?", uuid)
	return c.SendStatus(fiber.StatusOK)
}