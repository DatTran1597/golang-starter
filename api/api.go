package api

import (
	"github.com/DatTran1597/golang-starter/app"
	"github.com/gin-gonic/gin"
)

type Routers struct {
	Root  *gin.Engine
	APIv1 *gin.RouterGroup

	Users *gin.RouterGroup
}

type API struct {
	App         *app.App
	BaseRouters *Routers
}

func Init(app *app.App, root *gin.Engine) *API {
	api := &API{
		App:         app,
		BaseRouters: &Routers{},
	}

	api.BaseRouters.Root = root
	api.BaseRouters.APIv1 = api.BaseRouters.Root.Group("/api/v1")
	api.BaseRouters.Users = api.BaseRouters.APIv1.Group("/users")

	api.HealthCheck()
	api.InitUser()
	return api
}

func (a *API) Run() error {
	err := a.BaseRouters.Root.Run(a.App.Config.ServiceSettings.Port)
	if err != nil {
		return err
	}
	return nil
}
