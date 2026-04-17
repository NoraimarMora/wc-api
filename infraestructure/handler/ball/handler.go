package ball

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"wc-api/model"
)

type UseCase interface {
	GetBall() *model.Ball
}

type handler struct {
	useCase UseCase
}

func newHandler(useCase UseCase) handler {
	return handler{
		useCase: useCase,
	}
}

func (h *handler) getBall(c *gin.Context) {
	ball := h.useCase.GetBall()
	if ball == nil {
		c.JSON(http.StatusNotFound, model.Error{Message: fmt.Sprintf(model.DataNotFoundErr, "ball")})
		return
	}

	c.JSON(http.StatusOK, ball)
}
