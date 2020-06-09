package middlewares

import (
	"github.com/gofiber/fiber"
	"net/http"
	"sopordiversao/authc"
	"strings"
)

func Authentication(c *fiber.Ctx) {
	userToken := string(c.Fasthttp.Request.Header.Peek("Authorization"))
	//userEmail:=c.Query()
	if userToken == "" {
		c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Ta logado não Tio(a)",
		})
		return
	}
	split := strings.Split(userToken, " ")
	if len(split) != 2 || split[0] != "Bearer" {
		c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Não é JWT :D",
		})
		return
	}
	ok := authc.ValidateToken(split[01])
	if ok != nil {
		c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Não deu ruim legal",
		})
		return
	}
	c.Next()
	return
}
