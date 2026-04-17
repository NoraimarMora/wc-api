package mascot

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
	GetMascots(filters model.Filters) []*model.Mascot
	GetByID(id int64) (*model.Mascot, error)
}

type handler struct {
	useCase UseCase
}

func newHandler(useCase UseCase) handler {
	return handler{
		useCase: useCase,
	}
}

func (h *handler) getMascots(c *gin.Context) {
	filters, err := common.ValidateQueryParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Message: fmt.Sprintf(model.BadRequestErr, err.Error())})
		return
	}

	mascots := h.useCase.GetMascots(filters)
	if len(mascots) == 0 {
		c.JSON(http.StatusNotFound, model.Error{Message: fmt.Sprintf(model.DataNotFoundErr, "mascots")})
		return
	}

	c.JSON(http.StatusOK, mascots)
}

func (h *handler) getByID(c *gin.Context) {
	mascotIDParam := c.Param("id")
	mascotID, err := strconv.Atoi(mascotIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Message: fmt.Sprintf(model.IDNotANumberMessageErr, mascotIDParam)})
		return
	}

	city, err := h.useCase.GetByID(int64(mascotID))
	if errors.Is(err, model.ErrNotFound) {
		c.JSON(http.StatusNotFound, model.Error{Message: fmt.Sprintf(model.DataNotFoundErr, "mascot")})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error{Message: fmt.Sprintf(model.UnexpectedErrOccurred, err.Error())})
		return
	}

	c.JSON(http.StatusOK, city)

}
