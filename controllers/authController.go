package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leandromello/api-auth-jwt-token/databases"
	"github.com/leandromello/api-auth-jwt-token/models"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name: data["name"],
		Email: data["email"],
		Password: password,
	}

	databases.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	databases.DB.Where("email = ?", data["email"]).Find(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Usuário não encontrado!",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Password incorreto!",
		})
	}

	// claims := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.StandardClaims{
	// 	Issuer: 		strconv.Itoa(int(user.Id)),
	// 	ExpiresAt: 	time.Now().Add(time.Hour * 24).Unix(), // 1 dia
	// })

	// token, err := claims.SignedString([]byte(SecretKey))

	// if err != nil {
	// 	c.Status(fiber.StatusInternalServerError)
	// 	return c.JSON(fiber.Map{
	// 		"message": "Não foi possível fazer login",
	// 	})
	// }

	// cookie := fiber.Cookie{
	// 	Name: 		"jwt",
	// 	Value: 		token,
	// 	Expires: 	time.Now().Add(time.Hour * 24),
	// 	HTTPOnly: true,
	// }

	// c.Cookie(&cookie)

	// return c.JSON(fiber.Map{
	// 	"message": "Sucesso",
	// })

	return c.JSON(user)
}