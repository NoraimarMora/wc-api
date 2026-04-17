package event

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"wc-api/model"
)

type UseCase interface {
	GetAll() model.Events
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
	events := h.useCase.GetAll()
	if len(events) == 0 {
		c.JSON(http.StatusNotFound, model.Error{Message: fmt.Sprintf(model.DataNotFoundErr, "events")})
		return
	}

	c.JSON(http.StatusOK, events)
}
