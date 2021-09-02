package v1

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/abcbrasil/accesstoken", GerarToken())


}
