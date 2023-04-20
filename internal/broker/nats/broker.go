package nats

import (
	"encoding/json"
	"github.com/bemmanue/wildberries_L0/internal/cache"
	"github.com/bemmanue/wildberries_L0/internal/cache/mapcache"
	"github.com/bemmanue/wildberries_L0/internal/model"
	"github.com/bemmanue/wildberries_L0/internal/store"
	"github.com/go-playground/validator/v10"
	"github.com/nats-io/stan.go"
	"log"
)

type Broker struct {
	conn      stan.Conn
	store     store.Store
	cache     cache.Cache
	validator *validator.Validate
}

func New(store store.Store, cache cache.Cache) (*Broker, error) {
	conn, err := stan.Connect("test-cluster", "subscriber", stan.NatsURL(stan.DefaultNatsURL))
	if err != nil {
		return nil, err
	}

	b := &Broker{
		conn:      conn,
		store:     store,
		cache:     cache,
		validator: validator.New(),
	}

	return b, nil
}

func (b *Broker) Subscribe() error {
	_, err := b.conn.Subscribe("subject", func(msg *stan.Msg) {
		if err := msg.Ack(); err != nil {
			log.Println(err)
			return
		}

		var order model.Order
		if err := json.Unmarshal(msg.Data, &order); err != nil {
			log.Println(err)
			return
		}

		if err := b.validator.Struct(order); err != nil {
			log.Println(err)
			return
		}

		data, err := json.Marshal(order)
		if err != nil {
			log.Println(err)
			return
		}

		orderJSON := model.OrderJSON{
			OrderUID: order.OrderUID,
			Data:     data,
		}

		if err := b.store.Order().Create(&orderJSON); err != nil {
			log.Println(err)
			return
		}

		log.Printf("order with order_uid=%s stored to database\n", order.OrderUID)

		if err := b.cache.Order().Create(&orderJSON); err != nil {
			b.cache, _ = mapcache.New(b.store)
		}

		log.Printf("order with order_uid=%s stored to cache\n", order.OrderUID)

	}, stan.SetManualAckMode())

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
