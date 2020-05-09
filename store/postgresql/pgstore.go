package postgresql

import (
	"github.com/rs/zerolog/log"

	"github.com/DatTran1597/golang-starter/model"
	"github.com/DatTran1597/golang-starter/store"
	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	store.Store
	settings model.SQLSettings
	db       *gorm.DB

	user store.UserStore
}

func NewPostgres(settings model.SQLSettings) *PostgresStore {
	p := &PostgresStore{
		settings: settings,
	}

	p.initConnection()
	p.user = NewUserStore(p)
	return p
}

func (p *PostgresStore) initConnection() {
	db, err := gorm.Open(p.settings.DriverName, p.settings.URI)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open SQL connection")
	}
	p.db = db
}

func (p *PostgresStore) User() store.UserStore {
	return p.user
}
