package sound

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"wc-api/model"
)

type UseCase interface {
	GetSound() *model.Sound
}

type handler struct {
	useCase UseCase
}

func newHandler(useCase UseCase) handler {
	return handler{
		useCase: useCase,
	}
}

func (h *handler) getSound(c *gin.Context) {
	sound := h.useCase.GetSound()
	if sound == nil {
		c.JSON(http.StatusNotFound, model.Error{Message: fmt.Sprintf(model.DataNotFoundErr, "sound")})
		return
	}

	c.JSON(http.StatusOK, sound)
}
