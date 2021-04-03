// Copyright 2021, The Go Authors. All rights reserved.
// Author: crochee
// Date: 2021/4/3

package message

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
)

// todo self done
func Setup(ctx context.Context) error {
	logger := watermill.NewStdLogger(true, true)
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		return err
	}

	// SignalsHandler will gracefully shutdown Router when SIGTERM is received.
	// You can also close the router by just calling `r.Close()`.
	//router.AddPlugin(plugin.SignalsHandler)

	// Router level middleware are executed for every message sent to the router
	router.AddMiddleware(
		// CorrelationID will copy the correlation id from the incoming message's metadata to the produced messages
		middleware.CorrelationID,

		// The handler function is retried if it returns an error.
		// After MaxRetries, the message is Nacked and it's up to the PubSub to resend it.
		middleware.Retry{
			MaxRetries:      3,
			InitialInterval: time.Millisecond * 100,
			Logger:          logger,
		}.Middleware,

		// Recoverer handles panics from handlers.
		// In this case, it passes them as errors to the Retry middleware.
		middleware.Recoverer,
	)

	// For simplicity, we are using the gochannel Pub/Sub here,
	// You can replace it with any Pub/Sub implementation, it will work the same.
	//pubSub := gochannel.NewGoChannel(gochannel.Config{}, logger)
	var publisher *amqp.Publisher
	if publisher, err = amqp.NewPublisher(amqp.Config{
		Connection: amqp.ConnectionConfig{
			AmqpURI: "amqp://admin:1234567@192.168.31.62:5672/",
		},
		Marshaler: amqp.DefaultMarshaler{NotPersistentDeliveryMode: true},
		Exchange: amqp.ExchangeConfig{
			GenerateName: func(topic string) string {
				switch topic {
				case "micro_topic":
					return "micro"
				default:
					return ""
				}
			},
			Type: "fanout",
		},
		QueueBind: amqp.QueueBindConfig{
			GenerateRoutingKey: func(topic string) string {
				return ""
			},
		},
		Publish: amqp.PublishConfig{
			GenerateRoutingKey: func(topic string) string {
				return ""
			},
		},
		TopologyBuilder: &amqp.DefaultTopologyBuilder{},
	}, logger); err != nil {
		return err
	}
	var subscriber *amqp.Subscriber
	if subscriber, err = amqp.NewSubscriber(amqp.Config{
		Connection: amqp.ConnectionConfig{
			AmqpURI: "amqp://admin:1234567@192.168.31.62:5672/",
		},
		Marshaler: amqp.DefaultMarshaler{NotPersistentDeliveryMode: true},
		Exchange: amqp.ExchangeConfig{
			GenerateName: func(topic string) string {
				switch topic {
				case "micro_topic":
					return "micro"
				default:
					return ""
				}
			},
			Type: "fanout",
		},
		QueueBind: amqp.QueueBindConfig{
			GenerateRoutingKey: func(topic string) string {
				return ""
			},
		},
		Queue: amqp.QueueConfig{
			GenerateName: func(topic string) string {
				return topic
			},
		},
		Consume: amqp.ConsumeConfig{
			Qos: amqp.QosConfig{
				PrefetchCount: 1,
			},
		},
		TopologyBuilder: &amqp.DefaultTopologyBuilder{},
	}, logger); err != nil {
		return err
	}
	// Producing some incoming messages in background
	go publishMessages(publisher)

	// AddHandler returns a handler which can be used to add handler level middleware
	//handler := router.AddHandler(
	//	"handler",        // handler name, must be unique
	//	"sub_read_topic", // topic from which we will read events
	//	subscriber,
	//	"pub_write_topic", // topic to which we will publish events
	//	publisher,
	//	HandlerFunc,
	//)

	// Handler level middleware is only executed for a specific handler
	// Such middleware can be added the same way the router level ones
	//handler.AddMiddleware(func(h message.HandlerFunc) message.HandlerFunc {
	//	return func(message *message.Message) ([]*message.Message, error) {
	//		log.Println("executing handler specific middleware for ", message.UUID)
	//		return h(message)
	//	}
	//})

	// just for debug, we are printing all messages received on `incoming_messages_topic`
	//router.AddNoPublisherHandler(
	//	"print_incoming_messages",
	//	"incoming_messages_topic",
	//	subscriber,
	//	printMessages,
	//)

	go func() {
		for {
			receiveChan, err := subscriber.Subscribe(ctx, "micro_topic")
			if err != nil {
				fmt.Println(err)
				return
			}
			msg := <-receiveChan
			fmt.Printf(
				"\n> Received message: %s\n> %s\n> metadata: %v\n\n",
				msg.UUID, string(msg.Payload), msg.Metadata,
			)
		}
	}()
	// just for debug, we are printing all events sent to `outgoing_messages_topic`
	//router.AddNoPublisherHandler(
	//	"print_outgoing_messages",
	//	"outgoing_messages_topic",
	//	subscriber,
	//	printMessages,
	//)

	// Now that all handlers are registered, we're running the Router.
	// Run is blocking while the router is running.
	return router.Run(ctx)
}

func publishMessages(publisher message.Publisher) {
	for {
		msg := message.NewMessage(watermill.NewUUID(), []byte("Hello, world!"))
		middleware.SetCorrelationID(watermill.NewUUID(), msg)

		log.Printf("sending message %s, correlation id: %s\n", msg.UUID, middleware.MessageCorrelationID(msg))

		if err := publisher.Publish("micro_topic", msg); err != nil {
			return
		}

		time.Sleep(time.Second)
	}
}

func printMessages(msg *message.Message) error {
	fmt.Printf(
		"\n> Received message: %s\n> %s\n> metadata: %v\n\n",
		msg.UUID, string(msg.Payload), msg.Metadata,
	)
	return nil
}

func HandlerFunc(msg *message.Message) ([]*message.Message, error) {
	log.Println("structHandler received message", msg.UUID)

	msg = message.NewMessage(watermill.NewUUID(), []byte("message produced by structHandler"))
	return message.Messages{msg}, nil
}
