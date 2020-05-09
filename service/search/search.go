package search

import "git.lozi.space/beast/be-mandat/model"

type SearchService interface {
	Init() error
	IndexUser(*model.User) error
	SearchUserByName(name string) (ids []int, record int64, err error)
	DeleteUser(id int) error
}
