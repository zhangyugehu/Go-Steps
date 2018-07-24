package persist

import (
	"testing"
	"study/crawler/model"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"encoding/json"
	"study/crawler/engine"
)

func TestItemSaver(t *testing.T) {

	expected := engine.Item{
		Url:	"http://album.zhenai.com/u/1234567",
		Type:	"zhenai",
		Id:		"1234567",
		Payload:model.Profile{
			Age:		34,
			Height:		162,
			Weight:		57,
			Name:		"安静的雪",
		},
	}

	err := save(expected)
	if err != nil{
		panic(err)
	}

	// TODO using docker go client
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err!=nil{
		panic(err)
	}

	resp, err :=client.Get().
		Index("dating_profile").
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err!=nil{
		panic(err)
	}

	//t.Logf("%s\n", resp.Source)

	var actual engine.Item
	err = json.Unmarshal(
		[]byte(*resp.Source), &actual)

	if err!=nil{
		panic(err)
	}

	actualProfile,_ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile


	if actual != expected {
		t.Errorf("got %v; excepted %v",
			actual, expected)
	}

}
