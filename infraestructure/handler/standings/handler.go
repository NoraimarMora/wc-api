package standings

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"wc-api/model"
)

type UseCase interface {
	GetAll() model.StandingsByGroupResponse
	GetByGroup(group string) ([]*model.StandingsResponse, error)
}

type handler struct {
	useCase UseCase
}

func newHandler(useCase UseCase) handler {
	return handler{
		useCase: useCase,
	}
}

func (h *handler) getAll(c *gin.Context) {
	standings := h.useCase.GetAll()
	if len(standings) == 0 || standings == nil {
		c.JSON(http.StatusNotFound, model.Error{Message: fmt.Sprintf(model.DataNotFoundErr, "standings")})
		return
	}

	c.JSON(http.StatusOK, standings)
}

func (h *handler) getByGroup(c *gin.Context) {
	group := c.Param("group")

	standings, err := h.useCase.GetByGroup(group)
	if errors.Is(err, model.ErrNotFound) {
		c.JSON(http.StatusNotFound, model.Error{Message: fmt.Sprintf(model.DataNotFoundErr, "standings")})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error{Message: fmt.Sprintf(model.UnexpectedErrOccurred, err.Error())})
		return
	}

	c.JSON(http.StatusOK, standings)

}
