package gcp_pubsub

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
)

type PubSub interface {
	Publish(topic, message string) error
	Subscribre(topic_name, subscription_name string, subscription_function SubscriptionFunc) error
}

type SubscriptionFunc func(message string) error

type pubSub struct {
	client       *pubsub.Client
	topic        *pubsub.Topic
	subscription *pubsub.Subscription
}

func NewPubSub() PubSub {
	client := createClient()

	return &pubSub{
		client:       client,
		topic:        &pubsub.Topic{},
		subscription: &pubsub.Subscription{},
	}
}

func (p *pubSub) Publish(topic_name, message string) error {
	err := p.createTopic(topic_name)
	if err != nil {
		return err
	}

	p.topic.Publish(context.Background(), &pubsub.Message{
		Data: []byte(message),
	})

	return nil
}

func (p *pubSub) Subscribre(topic_name, subscription_name string, subscription_function SubscriptionFunc) error {
	err := p.createTopic(topic_name)
	if err != nil {
		return err
	}

	err = p.createSubscription(subscription_name, topic_name)
	if err != nil {
		return err
	}

	p.subscription.Receive(context.Background(), func(ctx context.Context, msg *pubsub.Message) {
		err := subscription_function(string(msg.Data))
		if err != nil {
			return
		}

		msg.Ack()
	})

	return nil
}

func (p *pubSub) createTopic(topic_name string) error {
	topic, err := createTopicIfNotExists(p.client, topic_name)
	if err != nil {
		return err
	}

	p.topic = topic
	return err
}

func (p *pubSub) createSubscription(subscription_name, topic_name string) error {
	err := p.createTopic(topic_name)
	if err != nil {
		return err
	}

	sub, err := getOrCreateSubscription(context.Background(), p.client, subscription_name, p.topic)
	if err != nil {
		return err
	}

	p.subscription = sub
	return nil
}

func createClient() *pubsub.Client {
	ctx := context.Background()
	proj := os.Getenv("PUBSUB_PROJECT_ID")
	client, err := pubsub.NewClient(ctx, proj)
	if err != nil {
		log.Fatalf("Could not create pubsub Client: %v", err)
	}

	return client
}

func createTopicIfNotExists(c *pubsub.Client, topic string) (*pubsub.Topic, error) {
	ctx := context.Background()
	t := c.Topic(topic)
	ok, err := t.Exists(ctx)
	if err != nil {
		return nil, err
	}
	if ok {
		return t, nil
	}
	t, err = c.CreateTopic(ctx, topic)
	if err != nil {
		log.Fatalf("Failed to create the topic: %v", err)
	}
	return t, nil
}

func getOrCreateSubscription(ctx context.Context, client *pubsub.Client, subID string, topic *pubsub.Topic) (*pubsub.Subscription, error) {
	cfg := &pubsub.SubscriptionConfig{
		Topic: topic,
	}

	sub := client.Subscription(subID)
	ok, err := sub.Exists(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to check if subscription exists: %v", err)
	}
	if !ok {
		sub, err = client.CreateSubscription(ctx, subID, *cfg)
		if err != nil {
			return nil, fmt.Errorf("failed to create subscription (%q): %v", "REGISTER_SUCESS", err)
		}
	}
	return sub, nil
}
