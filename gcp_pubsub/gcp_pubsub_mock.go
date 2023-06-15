package gcp_pubsub

import (
	"github.com/stretchr/testify/mock"
)

type PubSubMock struct {
	mock.Mock
}

func NewPubSubMock() *PubSubMock {
	return &PubSubMock{}
}

func (m *PubSubMock) Publish(topic_name string, message interface{}) error {
	args := m.Called(topic_name, message)
	return args.Error(0)
}

func (m *PubSubMock) Subscribre(topic_name, subscription_name string, subscription_function SubscriptionFunc) error {
	return m.Called(topic_name, subscription_name, subscription_function).Error(0)
}
