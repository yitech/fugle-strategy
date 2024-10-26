package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yitech/fugle-strategy/internal/env"
)

func NewSystemAPI(rg *gin.RouterGroup) {
	h := systemHandler{}
	srg := rg.Group("/system")
	srg.GET("/version", h.getVersion)
}

type systemHandler struct{}

type versionResponse struct {
	Version    string `json:"version"`
	BuildTime  string `json:"buildTime"`
	CommitHash string `json:"commitHash"`
}

// getVersion godoc
// @Summary Get version information
// @Description Returns the application's current version, build time, and commit hash
// @Tags System
// @Produce json
// @Success 200 {object} versionResponse
// @Router /v1/system/version [get]
func (s systemHandler) getVersion(c *gin.Context) {
	c.JSON(200, versionResponse{
		Version:    env.Version,
		BuildTime:  env.BuildTime,
		CommitHash: env.CommitHash,
	})
}
