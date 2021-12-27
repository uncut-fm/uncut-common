package pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
)

func NewPublisher(pubsubClient *pubsub.Client, topicName string) *Publisher {
	tp := pubsubClient.Topic(topicName)
	return &Publisher{cl: tp}
}

type PublisherInterface interface {
	Publish(context.Context, interface{}) error
}

// Publisher represents google pubsub activity publisher
type Publisher struct {
	cl *pubsub.Topic
}

// Publish publishes event to activity queue
func (d *Publisher) Publish(ctx context.Context, message interface{}) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}
	_, err = d.cl.Publish(ctx, &pubsub.Message{
		Data: body,
	}).Get(ctx)

	return err
}
