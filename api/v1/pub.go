package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yitech/fugle-strategy/pkg/pubsub"
)

func NewPubAPI(rg *gin.RouterGroup, pub *pubsub.Publisher) {
	h := pubHandler{
		pub: pub,
	}
	srg := rg.Group("/pub")
	srg.POST("/offline/kline", h.publishOfflineKline)
}

type pubHandler struct {
	pub *pubsub.Publisher
}

type publishOfflineKlineRequest struct {
	Symbol    string `json:"symbol"`
	Topic     string `json:"topic"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Interval  int    `json:"interval"`
}

type publishOfflineKlineResponse struct {
	ID string `json:"id"`
}

// publishOfflineKline godoc
// @Summary Publish offline kline data
// @Description Publishes offline kline data for a specified symbol, date range, and interval.
// @Tags Kline
// @Accept json
// @Produce json
// @Param request body publishOfflineKlineRequest true "Offline Kline Request"
// @Success 200 {object} publishOfflineKlineResponse
// @Router /v1/pub/offline/kline [post]
func (h pubHandler) publishOfflineKline(c *gin.Context) {

	c.JSON(200, publishOfflineKlineResponse{
		ID: "123",
	})
}
