package ranking

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"wc-api/model"
)

type UseCase interface {
	GetWorldRanking() model.WorldRankingResponse
}

type handler struct {
	useCase UseCase
}

func newHandler(useCase UseCase) handler {
	return handler{
		useCase: useCase,
	}
}

func (h *handler) getWorldRanking(c *gin.Context) {
	worldRanking := h.useCase.GetWorldRanking()
	if len(worldRanking) == 0 || worldRanking == nil {
		c.JSON(http.StatusNotFound, model.Error{Message: fmt.Sprintf(model.DataNotFoundErr, "world ranking")})
		return
	}

	c.JSON(http.StatusOK, worldRanking)
}
