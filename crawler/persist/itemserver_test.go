package persist

import (
	"learn_go/crawler/engine"
	"encoding/json"
	"learn_go/crawler/model"
	"context"
	"github.com/olivere/elastic/v7"
	"testing"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url: "",
		Type: "zhenai",
		Id: "",
		Payload: model.Profile{
			Age:        27,
			Height:     165,
			Weight:     70,
			Income:     "月收入:1.2-2万",
			Gender:     "男士",
			Name:       "张韶勇",
			Xinzuo:     "射手座(11.22-12.21)",
			Occupation: "工程师",
			Marriage:   "未婚",
			House:      "已购房",
			Hokou:      "籍贯:重庆",
			Education:  "大专",
			Car:        "已买车",
		},
	}

	err := save(expected)

	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index("dating_test").Type(expected.Type).Id(expected.Id).Do(context.Background())

	var actual engine.Item
	err = json.Unmarshal([]byte(resp.Source), &actual)

	actualProfile, _ := model.FromJSONObj(actual.Payload)
	actual.Payload = actualProfile

	if err != nil {
		panic(err)
	}

	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}

}