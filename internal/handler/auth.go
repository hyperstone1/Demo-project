package handler

import (
	"net/http"

	"github.com/hyperstone1/TestCRUD/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body model.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]

func (s *CatHandler) SignUp(c echo.Context) error {
	var input model.User

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	id, err := s.service.CreateUser(input)
	if err != nil {
		logrus.Println("Ошибка signUp CU")
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
func (s *CatHandler) SignIn(c echo.Context) error {
	var input signInInput

	if err := c.Bind(&input); err != nil {
		logrus.Println("Ошибка signIn Input")
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	token, err := s.service.GenerateJWT(input.Username, input.Password)
	if err != nil {
		logrus.Println("This user is not on db")
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
