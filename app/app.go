package app

import (
	"github.com/DatTran1597/golang-starter/config"
	"github.com/DatTran1597/golang-starter/model"
	"github.com/DatTran1597/golang-starter/service/cache"
	"github.com/DatTran1597/golang-starter/service/search"
	"github.com/DatTran1597/golang-starter/store"
	"github.com/DatTran1597/golang-starter/store/postgresql"

	"github.com/rs/zerolog/log"
)

type App struct {
	Config model.Configuration

	Store  store.Store
	Cache  cache.Cacher
	Search search.SearchService
}

func New(cfg string) (*App, error) {
	c, err := config.Load(cfg)
	if err != nil {
		return nil, err
	}
	app := &App{
		Config: *c,
	}

	log.Info().Msg("Server is initializing...")
	app.Store = postgresql.NewPostgres(app.Config.SQLSettings)
	app.Cache, _ = cache.NewRedisCacher(&app.Config.CacheSettings)
	app.Search, err = search.NewElasticSearch(&app.Config.SearchSettings)
	if err != nil {
		return nil, err

	}

	return app, nil
}
