package __mocks

type IEmailGateway interface {
	SendGreetingsEmail(email string)
}

func NewController(emailGateway IEmailGateway) Controller {
	return Controller{emailGateway: emailGateway}
}

type Controller struct {
	emailGateway IEmailGateway
}

func (c Controller) GreetUser(email string) {
	c.emailGateway.SendGreetingsEmail(email)
}