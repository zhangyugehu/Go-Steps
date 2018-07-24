package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"study/crawler/engine"
	"errors"
)

func ItemSaver() (chan engine.Item, error) {

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err!=nil{
		return nil, err
	}
	out := make(chan engine.Item)
	itemCount := 0
	go func() {
		item := <- out
		log.Printf("Got item #%d by saver: %v", itemCount, item)
		itemCount++
	}()

	return out, nil
}

func save(client *elastic.Client, item engine.Item) error {

	if item.Type == ""{
		return errors.New("must supply Type")
	}

	indexService := client.Index().
		Index("dating_profile").
		Type(item.Type).
		BodyJson(item)

	if item.Id != ""{
		indexService.Id(item.Id)
	}

	_, err := indexService.Do(context.Background())

	if err != nil{
		return err
	}

	return nil

}