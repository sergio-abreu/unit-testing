package bus

type MessageBus struct{}

func (MessageBus) SendEmailChangeMessage(userID int, email string) error {
	return nil
}
