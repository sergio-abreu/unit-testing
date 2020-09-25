package _3_refactor

import "fmt"

type Message struct {
	Header string
	Body   string
	Footer string
}

type IRenderer interface {
	Render(message Message) string
}

func NewMessageRenderer() MessageRenderer {
	return MessageRenderer{SubRenders: []IRenderer{
		HeaderRenderer{},
		BodyRenderer{},
		FootRenderer{},
	}}
}

type MessageRenderer struct {
	SubRenders []IRenderer
}

func (m MessageRenderer) Render(message Message) (html string) {
	for _, render := range m.SubRenders {
		html += render.Render(message)
	}
	return
}

type BodyRenderer struct{}

func (b BodyRenderer) Render(message Message) string {
	return fmt.Sprintf("<b>%s</b>", message.Body)
}

type HeaderRenderer struct{}

func (b HeaderRenderer) Render(message Message) string {
	return fmt.Sprintf("<h1>%s</h1>", message.Header)
}

type FootRenderer struct{}

func (b FootRenderer) Render(message Message) string {
	return fmt.Sprintf("<i>%s</i>", message.Footer)
}
