package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/DatTran1597/golang-starter/model"
	"github.com/gin-gonic/gin"
)

func (api *API) InitUser() {
	api.BaseRouters.Users.POST("", api.userRegister)
	api.BaseRouters.Users.GET("", api.getUserHandler)
	api.BaseRouters.Users.DELETE("/:userid", api.deleteUser)
}

func (api *API) userRegister(c *gin.Context) {
	user := &model.User{}
	err := c.ShouldBindJSON(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	err = api.App.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, user)
}

func (api *API) getUserHandler(c *gin.Context) {
	var err error
	var res interface{}
	switch {
	case c.Query("query") == "numberofuser":
		res, err = api.App.CountUser()
	case c.Query("name") != "":
		name := strings.ToLower(c.Query("name"))
		res, err = api.App.SearchUserByName(name)
	default:
		res, err = api.App.GetUsers()
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, res)
}

func (api *API) deleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	err = api.App.DeleteUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusNoContent, nil)
}
