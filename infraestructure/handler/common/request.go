package common

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"wc-api/model"
)

func ValidateQueryParams(c *gin.Context) (model.Filters, error) {
	filters := model.Filters{}

	cityIDParam := c.Query("city_id")
	group := c.Query("group")
	roundParam := c.Query("round")
	status := c.Query("status")
	homeID := c.Query("home_id")
	awayID := c.Query("away_id")
	country := c.Query("country")
	name := c.Query("name")
	fromParam := c.Query("from")
	toParam := c.Query("to")

	if cityIDParam != "" {
		cityID, err := strconv.Atoi(cityIDParam)
		if err != nil {
			return filters, fmt.Errorf("city_id")
		}

		filters.CityID = int64(cityID)
	}

	if roundParam != "" {
		round, err := strconv.Atoi(roundParam)
		if err != nil {
			return filters, fmt.Errorf("round")
		}

		filters.Round = round
	}

	if fromParam != "" {
		from, err := time.Parse("2006-01-02", fromParam)
		if err != nil {
			return filters, fmt.Errorf("from date")
		}

		filters.From = from
	}

	if toParam != "" {
		to, err := time.Parse("2006-01-02", toParam)
		if err != nil {
			return filters, fmt.Errorf("to date")
		}

		filters.To = to
	}

	filters.Group = group
	filters.Status = status
	filters.HomeID = homeID
	filters.AwayID = awayID
	filters.Country = country
	filters.Name = name

	return filters, nil
}
