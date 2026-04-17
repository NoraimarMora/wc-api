package match

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
	GetMatches(filters model.Filters) []*model.Match
	GetByID(id int64) (*model.MatchResponse, error)
}

type handler struct {
	useCase UseCase
}

func newHandler(useCase UseCase) handler {
	return handler{
		useCase: useCase,
	}
}

func (h *handler) getMatches(c *gin.Context) {
	filters, err := common.ValidateQueryParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Message: fmt.Sprintf(model.BadRequestErr, err.Error())})
		return
	}

	matches := h.useCase.GetMatches(filters)
	if len(matches) == 0 {
		c.JSON(http.StatusNotFound, model.Error{Message: fmt.Sprintf(model.DataNotFoundErr, "matches")})
		return
	}

	c.JSON(http.StatusOK, matches)
}

func (h *handler) getByID(c *gin.Context) {
	matchIDParam := c.Param("id")
	matchID, err := strconv.Atoi(matchIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Message: fmt.Sprintf(model.IDNotANumberMessageErr, matchIDParam)})
		return
	}

	match, err := h.useCase.GetByID(int64(matchID))
	if errors.Is(err, model.ErrNotFound) {
		c.JSON(http.StatusNotFound, model.Error{Message: fmt.Sprintf(model.DataNotFoundErr, "match")})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error{Message: fmt.Sprintf(model.UnexpectedErrOccurred, err.Error())})
		return
	}

	c.JSON(http.StatusOK, match)
}
