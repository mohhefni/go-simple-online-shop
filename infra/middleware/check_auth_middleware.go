package middleware

import (
	"log"
	"mohhefni/go-online-shop/infra/errorpkg"
	"mohhefni/go-online-shop/infra/responsepkg"
	"mohhefni/go-online-shop/internal/config"
	"mohhefni/go-online-shop/utility"
	"strings"

	"github.com/labstack/echo/v4"
)

func CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorization := c.Request().Header.Get("Authorization")
		if authorization == "" {
			return responsepkg.NewResponse(
				responsepkg.WithStatus(errorpkg.ErrUnauthorized),
			).Send(c)
		}

		splitToken := strings.Split(authorization, "Bearer ")
		if len(splitToken) != 2 {
			log.Println("Token Invalid")
			return responsepkg.NewResponse(
				responsepkg.WithStatus(errorpkg.ErrUnauthorized),
			).Send(c)
		}

		token := splitToken[1]

		publicId, role, err := utility.ValidateToken(token, config.Cfg.App.Encrytion.JWTSecret)
		if err != nil {
			log.Println("Token Invalid")
			return responsepkg.NewResponse(
				responsepkg.WithStatus(errorpkg.ErrUnauthorized),
			).Send(c)
		}

		c.Set("public_id", publicId)
		c.Set("role", role)

		return next(c)
	}
}
