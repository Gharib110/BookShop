package elasticsearch

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	log2 "log"
	"os"
	"time"
)

const (
	envEsHosts = "ELASTIC_HOSTS"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string, string) (*elastic.GetResult, error)
	Search(string, elastic.Query) (*elastic.SearchResult, error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	logger := log2.Default()

	client, err := elastic.NewClient(
		elastic.SetURL(os.Getenv(envEsHosts)),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(logger),
		elastic.SetInfoLog(logger),
	)
	if err != nil {
		panic(err)
	}
	Client.setClient(client)

	// Create the index if it does not exists
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) Index(index string, docType string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().
		Index(index).
		Type(docType).
		BodyJson(doc).
		Do(ctx)

	if err != nil {
		log2.Printf(fmt.Sprintf("error when trying to index document in index %s\n", index), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Get(index string, docType string, id string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := c.client.Get().
		Index(index).
		Type(docType).
		Id(id).
		Do(ctx)
	if err != nil {
		log2.Printf(fmt.Sprintf("error when trying to get id %s\n", id), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Search(index string, query elastic.Query) (*elastic.SearchResult, error) {
	ctx := context.Background()
	result, err := c.client.Search(index).
		Query(query).
		RestTotalHitsAsInt(true).
		Do(ctx)
	if err != nil {
		log2.Printf(fmt.Sprintf("error when trying to search documents in index %s\n", index), err)
		return nil, err
	}
	return result, nil
}
