package news

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"wc-api/model"
)

type UseCase interface {
	GetAll() []*model.News
	GetByID(id int64) (*model.News, error)
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
	news := h.useCase.GetAll()
	if len(news) == 0 {
		c.JSON(http.StatusNotFound, model.Error{Message: fmt.Sprintf(model.DataNotFoundErr, "news")})
		return
	}

	c.JSON(http.StatusOK, news)
}

func (h *handler) getByID(c *gin.Context) {
	newsIDParam := c.Param("id")
	newsID, err := strconv.Atoi(newsIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Message: fmt.Sprintf(model.IDNotANumberMessageErr, newsIDParam)})
		return
	}

	city, err := h.useCase.GetByID(int64(newsID))
	if errors.Is(err, model.ErrNotFound) {
		c.JSON(http.StatusNotFound, model.Error{Message: fmt.Sprintf(model.DataNotFoundErr, "news")})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error{Message: fmt.Sprintf(model.UnexpectedErrOccurred, err.Error())})
		return
	}

	c.JSON(http.StatusOK, city)
}
