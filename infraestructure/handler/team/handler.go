package team

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"wc-api/infraestructure/handler/common"
	"wc-api/model"
)

type UseCase interface {
	GetTeams(filters model.Filters) []*model.Team
	GetByID(id string) (*model.TeamResponse, error)
}

type handler struct {
	useCase UseCase
}

func newHandler(useCase UseCase) handler {
	return handler{
		useCase: useCase,
	}
}

func (h *handler) getTeams(c *gin.Context) {
	filters, err := common.ValidateQueryParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Message: fmt.Sprintf(model.BadRequestErr, err.Error())})
		return
	}

	teams := h.useCase.GetTeams(filters)
	if len(teams) == 0 || teams == nil {
		c.JSON(http.StatusNotFound, model.Error{Message: fmt.Sprintf(model.DataNotFoundErr, "teams")})
		return
	}

	c.JSON(http.StatusOK, teams)
}

func (h *handler) getTeam(c *gin.Context) {
	teamID := c.Param("id")

	city, err := h.useCase.GetByID(teamID)
	if errors.Is(err, model.ErrNotFound) {
		c.JSON(http.StatusNotFound, model.Error{Message: fmt.Sprintf(model.DataNotFoundErr, "team")})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error{Message: fmt.Sprintf(model.UnexpectedErrOccurred, err.Error())})
		return
	}

	c.JSON(http.StatusOK, city)
}
