package handler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "UserId"
)

func (s *CatHandler) UserIdentity(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		header := c.Request().Header.Get(authorizationHeader)
		if header == "" {
			newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
			return err
		}
		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
			return err
		}
		UserId, err := s.service.ParseToken(headerParts[1])
		if err != nil {
			newErrorResponse(c, http.StatusUnauthorized, err.Error())
			return err
		}

		c.Set(userCtx, UserId)
		return next(c)
	}

}
