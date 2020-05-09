package search

import (
	"context"
	"fmt"
	"strconv"

	"github.com/DatTran1597/golang-starter/model"
	elastic "github.com/olivere/elastic/v7"
)

const userMapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"properties":{
			"name":{
				"type":"text"
			}
		}
	}
}`
const userIndex = "users"

type UserSearch struct {
	Name string `json:"name"`
}

type ElasticSearch struct {
	client *elastic.Client
}

func NewElasticSearch(settings *model.SearchSetting) (SearchService, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(settings.ConnectionURL),
		// elastic.SetBasicAuth(settings.UserName, settings.Password),
		elastic.SetSniff(settings.Sniff),
	)

	if err != nil {
		return nil, err
	}

	_, _, err = client.Ping(settings.ConnectionURL).Do(context.Background())
	if err != nil {
		return nil, err
	}
	es := &ElasticSearch{
		client: client,
	}

	err = es.Init()
	if err != nil {
		return nil, err
	}

	return es, nil
}

func (es *ElasticSearch) Init() error {
	exist, err := es.client.IndexExists(userIndex).Do(context.Background())
	if err != nil {
		return err
	}

	if !exist {
		_, err := es.client.CreateIndex(userIndex).BodyString(userMapping).Do(context.Background())
		if err != nil {
			return err
		}
	}

	return nil
}

func (es *ElasticSearch) IndexUser(user *model.User) error {
	userDoc := &UserSearch{
		Name: user.Name,
	}

	_, err := es.client.Index().Index(userIndex).
		Id(strconv.Itoa(user.ID)).BodyJson(userDoc).Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (es *ElasticSearch) SearchUserByName(name string) (ids []int, record int64, err error) {
	fmt.Println("name query", name)
	query := elastic.NewCommonTermsQuery("name", name)
	searchResult, err := es.client.Search().Index("users").Query(query).Pretty(true).Do(context.Background())
	if err != nil {
		return nil, 0, err
	}

	ids = make([]int, searchResult.TotalHits())
	for i, hit := range searchResult.Hits.Hits {
		id, _ := strconv.Atoi(hit.Id)
		if id > 0 {
			ids[i] = id
		}
	}

	return ids, searchResult.TotalHits(), nil
}

func (es *ElasticSearch) DeleteUser(id int) error {
	_, err := es.client.Delete().Index(userIndex).Id(strconv.Itoa(id)).Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
