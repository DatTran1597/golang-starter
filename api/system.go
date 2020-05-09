package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *API) HealthCheck() {
	api.BaseRouters.Root.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "start at port"+api.App.Config.ServiceSettings.Port)
	})
}
