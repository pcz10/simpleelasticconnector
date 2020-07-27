package persistance

import (
	"context"
	"encoding/json"
	"fmt"
	models "todo/model"

	"github.com/olivere/elastic"
)

type Manager interface {
	FindAll() []byte
	Add(note *models.Note) error
}

type DBmanager struct {
	client *elastic.Client
}

var ElasticClient = initializeESCClient()

func GetElasticClient() DBmanager {
	return ElasticClient
}

func initializeESCClient() DBmanager {
	elasticClient, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))
	if err != nil {
		panic("err")
	} else {
		fmt.Println("Elasticsearch initialized")
	}
	return DBmanager{client: elasticClient}
}

func (elasticClient DBmanager) FindAll() []models.Note {
	var notes []models.Note
	ctx := context.Background()
	searchService := GetElasticClient().client.Search().Index("notes")
	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("Search error. Err=", err)
	}
	for _, hit := range searchResult.Hits.Hits {
		var note models.Note
		err := json.Unmarshal(hit.Source, &note)
		if err != nil {
			fmt.Println("Unmarshall note error. Err=", err)
		}
		notes = append(notes, note)
	}
	fmt.Println("notes :", notes)
	return notes
}

func (elasticClient DBmanager) Add(note []byte) error {
	ctx := context.Background()
	data := string(note)
	ind, err := elasticClient.client.Index().Index("notes").BodyJson(data).Do(ctx)

	if err != nil {
		panic(err)
	}
	fmt.Println("Insertion Successful", ind)
	return nil
}
