package city

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"wc-api/infraestructure/handler/common"
	"wc-api/model"
)

type UseCase interface {
	GetCities(filters model.Filters) model.Cities
	GetByID(id int64) (*model.City, error)
}

type handler struct {
	useCase UseCase
}

func newHandler(useCase UseCase) handler {
	return handler{
		useCase: useCase,
	}
}

func (h *handler) getCities(c *gin.Context) {
	filters, err := common.ValidateQueryParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Message: fmt.Sprintf(model.BadRequestErr, err.Error())})
		return
	}

	cities := h.useCase.GetCities(filters)
	if len(cities) == 0 || cities == nil {
		c.JSON(http.StatusNotFound, model.Error{Message: fmt.Sprintf(model.DataNotFoundErr, "cities")})
		return
	}

	c.JSON(http.StatusOK, cities)
}

func (h *handler) getCity(c *gin.Context) {
	cityIDParam := c.Param("id")
	cityID, err := strconv.Atoi(cityIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Message: fmt.Sprintf(model.IDNotANumberMessageErr, cityIDParam)})
		return
	}

	city, err := h.useCase.GetByID(int64(cityID))
	if errors.Is(err, model.ErrNotFound) {
		c.JSON(http.StatusNotFound, model.Error{Message: fmt.Sprintf(model.DataNotFoundErr, "city")})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error{Message: fmt.Sprintf(model.UnexpectedErrOccurred, err.Error())})
		return
	}

	c.JSON(http.StatusOK, city)
}
