package persist

import (
	"fmt"
	"context"
	"log"
	"github.com/olivere/elastic/v7"
)

// ItemSaver func
func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item saver: got item #%d: %v", itemCount, item)
			itemCount ++
			save(item)
		}
	}()

	return out
}


func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)

	if err != nil {
		return "", err
	}

	resp, err := client.Index().
				Index("dating_profit").
				Type("zhenai").
				BodyJson(item).Do(context.Background())

	if err != nil {
		return "", err
	}

	fmt.Printf("%+v", resp)

	return resp.Id, nil
}