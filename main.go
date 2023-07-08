package main

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://elasticsearch.draft.com",
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("failed at set es client. err is %s", err.Error())
		return
	}
	res, err := es.Ping()
	if err != nil {
		log.Fatalf("failed at es ping. err is %s", err.Error())
		return
	}
	defer res.Body.Close()
	log.Printf("res is %s", res.String())
	log.Printf("status is %s", res.Status())
}
