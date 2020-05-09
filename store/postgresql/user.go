package postgresql

import (
	"fmt"

	"github.com/DatTran1597/golang-starter/model"
	"github.com/DatTran1597/golang-starter/store"
)

const usersTable = "users"

type PostgresUserStore struct {
	postgres *PostgresStore
}

func NewUserStore(pg *PostgresStore) store.UserStore {
	return &PostgresUserStore{
		postgres: pg,
	}
}

func (p *PostgresUserStore) CreateUser(user *model.User) error {
	return p.postgres.db.Create(user).Error
}

func (p *PostgresUserStore) GetUserById(id int) (*model.User, error) {
	return nil, nil
}

func (p *PostgresUserStore) GetUsers() ([]model.User, error) {
	users := []model.User{}
	err := p.postgres.db.Where("deleted_at is null").Find((&users)).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (p *PostgresUserStore) GetUserByIds(ids []int) ([]model.User, error) {
	users := []model.User{}
	err := p.postgres.db.Where("deleted_at is null and id in (?)", ids).Find((&users)).Error
	fmt.Println("pq err:", err)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (p *PostgresUserStore) CountUsers() (int, error) {
	var count int
	err := p.postgres.db.Model(&model.User{}).Where("deleted_at is null").Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (p *PostgresUserStore) UpdateUser(user *model.User) error {
	return nil
}

func (p *PostgresUserStore) DeleteUser(id int) error {
	return p.postgres.db.Where("id = ?", id).Delete(&model.User{}).Error
}
