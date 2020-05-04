package persist

import (
	"errors"
	"learn_go/crawler/engine"
	"fmt"
	"context"
	"log"
	"github.com/olivere/elastic/v7"
)

// ItemSaver func
func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item saver: got item #%d: %v", itemCount, item)
			itemCount++
			err := save(item)
			if err != nil {
				log.Printf("Item Saver: error "+
					"saving item %v: %v",
					item, err)
			}
		}
	}()

	return out
}

func save(item engine.Item) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)

	if err != nil {
		return err
	}

	if item.Type == "" {
		return errors.New("Must supply Type")
	}

	indexService := client.Index().
		Index("dating_profit").
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err = indexService.Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
