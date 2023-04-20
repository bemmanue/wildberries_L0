package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/bemmanue/wildberries_L0/internal/logger"
	"github.com/bemmanue/wildberries_L0/internal/model"
	"github.com/nats-io/stan.go"
	"log"
	"math/rand"
	"os"
)

func init() {
	log.SetFlags(0)
	log.SetOutput(new(logger.Writer))
}

func main() {
	n := flag.Int("n", 1, "count of messages")
	flag.Parse()

	// connect to nats streaming
	sc, err := stan.Connect("test-cluster", "publisher", stan.NatsURL(stan.DefaultNatsURL))
	if err != nil {
		log.Fatalln(err)
	}
	defer sc.Close()

	// read json model
	modelJSON, err := os.ReadFile("model.json")

	// convert model into order struct
	var order model.Order
	if err := json.Unmarshal(modelJSON, &order); err != nil {
		log.Fatalln(err)
	}

	// generate data with unique order_uid and publish to nats chanel
	for i := 0; i < *n; i++ {
		order.OrderUID = fmt.Sprintf("%xtest", rand.Int())

		data, err := json.Marshal(order)
		if err != nil {
			log.Fatalln(err)
		}

		if err := sc.Publish("subject", data); err != nil {
			log.Fatalln(err)
		}

		log.Printf("published order with order_uid=%s\n", order.OrderUID)
	}
}
