package record

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"wc-api/model"
)

type UseCase interface {
	GetAll() []*model.Record
	GetByID(id int64) (*model.Record, error)
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
	worldRanking := h.useCase.GetAll()
	if len(worldRanking) == 0 || worldRanking == nil {
		c.JSON(http.StatusNotFound, model.Error{Message: fmt.Sprintf(model.DataNotFoundErr, "record")})
		return
	}

	c.JSON(http.StatusOK, worldRanking)
}

func (h *handler) getByID(c *gin.Context) {
	recordIDParam := c.Param("id")
	recordID, err := strconv.Atoi(recordIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Message: fmt.Sprintf(model.IDNotANumberMessageErr, recordIDParam)})
		return
	}

	city, err := h.useCase.GetByID(int64(recordID))
	if errors.Is(err, model.ErrNotFound) {
		c.JSON(http.StatusNotFound, model.Error{Message: fmt.Sprintf(model.DataNotFoundErr, "record")})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error{Message: fmt.Sprintf(model.UnexpectedErrOccurred, err.Error())})
		return
	}

	c.JSON(http.StatusOK, city)
}
