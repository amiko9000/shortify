package handler

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func Login(c *fiber.Ctx) error {
	type Login struct {
		Identity string `validate:"required" json:"identity"`
		Password string `validate:"required" json:"password"`
	}

	login := new(Login)
	if err := c.BodyParser(&login); err != nil {
		log.Println("Wrong input data")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if login.Identity != "admin" || login.Password != "password" {
		return c.SendStatus(fiber.StatusForbidden)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Admin Ivanovich"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	type Response struct {
		Token string `json:"token"`
	}

	res := Response{Token: t}
	return c.JSON(&res)
}