package search

import "github.com/DatTran1597/golang-starter/model"

type SearchService interface {
	Init() error
	IndexUser(*model.User) error
	SearchUserByName(name string) (ids []int, record int64, err error)
	DeleteUser(id int) error
}
