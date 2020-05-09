package app

import (
	"fmt"
	"strconv"

	"github.com/DatTran1597/golang-starter/model"
)

//NumberOfUser is a key for cache number of user
const NumberOfUser = "numberofuser"

func (a *App) CreateUser(user *model.User) error {
	err := a.Store.User().CreateUser(user)
	if err != nil {
		return err
	}

	nou, _ := a.Store.User().CountUsers()
	go a.cacheNumberOfUser(nou)
	go a.Search.IndexUser(user)
	return nil
}

func (a *App) GetUsers() ([]model.User, error) {
	return a.Store.User().GetUsers()
}

func (a *App) CountUser() (int, error) {
	nof, _ := a.getNumberOfUserfromCache()
	if nof != 0 {
		return nof, nil
	}

	nof, err := a.Store.User().CountUsers()
	if err != nil {
		return 0, err
	}

	return nof, nil
}

func (a *App) SearchUserByName(name string) ([]model.User, error) {
	ids, _, err := a.Search.SearchUserByName(name)
	fmt.Println("ids", ids)
	if err != nil {
		return nil, err
	}

	users, err := a.Store.User().GetUserByIds(ids)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (a *App) DeleteUser(userID int) error {
	go a.Search.DeleteUser(userID)
	return a.Store.User().DeleteUser(userID)
}

func (a *App) cacheNumberOfUser(number int) error {
	err := a.Cache.Set(NumberOfUser, number)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) getNumberOfUserfromCache() (int, error) {
	nof, err := a.Cache.Get(NumberOfUser)
	if err != nil {
		return 0, err
	}

	number, _ := strconv.Atoi(nof.(string))
	return number, nil
}
