package store

import "github.com/DatTran1597/golang-starter/model"

type Store interface {
	User() UserStore
}

type UserStore interface {
	CreateUser(user *model.User) error
	GetUserById(id int) (*model.User, error)
	GetUserByIds(ids []int) ([]model.User, error)
	GetUsers() ([]model.User, error)
	CountUsers() (int, error)
	UpdateUser(user *model.User) error
	DeleteUser(id int) error
}
