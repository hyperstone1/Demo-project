package handler

import (
	"errors"
	"net/http"

	"github.com/hyperstone1/TestCRUD/internal/model"
	"github.com/hyperstone1/TestCRUD/internal/service"

	"github.com/labstack/echo/v4"
)

// @Summary Create cat
// @Security ApiKeyAuth
// @Tags Cats
// @Description create cat
// @ID create-cat
// @Accept  json
// @Produce  json
// @Param input body model.Cats true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [post]

type CatHandler struct {
	service *service.CatService
}

func New(serv *service.CatService) (*CatHandler, error) {
	if serv == nil {
		return nil, errors.New("handler.New error")
	}
	return &CatHandler{serv}, nil
}

func (s *CatHandler) Create(c echo.Context) error {
	/*id := c.Get(userCtx)
	if id != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	}*/

	cat := new(model.Cat)
	if err := c.Bind(cat); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := s.service.Create(*cat)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "successfully added")
}

func (s *CatHandler) Get(c echo.Context) error {
	/*id := c.Get(userCtx)
	if id != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	}*/

	cat, err := s.service.Get()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, cat)
}

func (s *CatHandler) Update(c echo.Context) error {
	/*id := c.Get(userCtx)
	if id != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	}*/

	cat := new(model.Cat)
	if err := c.Bind(cat); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := s.service.Update(*cat)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "successfully updated")
}
func (s *CatHandler) Delete(c echo.Context) error {

	/*idU := c.Get(userCtx)
	if idU != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"id": idU,
		})
	}*/

	var err error
	id := c.Param("id")
	if id == "" {
		return err
	}
	err = s.service.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "successfully deleted")
}
func (s *CatHandler) GetById(c echo.Context) error {
	id := c.Param("id")

	cat, err := s.service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, cat)
}
