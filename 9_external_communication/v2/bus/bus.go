package bus

import "fmt"

type IMessageBus interface {
	SendEmailChangeMessage(userID int, email string) error
}

type MessageBus struct {}

func (MessageBus) SendEmailChangeMessage(userID int, email string) error {
	return nil
}

type Nats interface {
	Publish(subject, content string) error
}

func NewBroker(nats Nats) Broker {
	return Broker{nats: nats}
}

type Broker struct {
	nats Nats
}

func (b Broker) SendEmailChangeMessage(userID int, email string) error {
	return b.nats.Publish("email-changed", fmt.Sprintf(`{"id":%d,"email":"%s"}`, userID, email))
}