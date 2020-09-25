package _3_refactor

import (
	. "github.com/onsi/gomega"
	"io/ioutil"
	"os"
	"testing"
)

func Test_MessageRenderer_uses_correct_sub_renderers(t *testing.T) {
	g := NewGomegaWithT(t)
	sut := NewMessageRenderer()

	renderers := sut.SubRenders

	g.Expect(renderers).Should(
		HaveLen(3))
	g.Expect(renderers[0]).Should(
		BeAssignableToTypeOf(HeaderRenderer{}))
	g.Expect(renderers[1]).Should(
		BeAssignableToTypeOf(BodyRenderer{}))
	g.Expect(renderers[2]).Should(
		BeAssignableToTypeOf(FootRenderer{}))
}

func Test_MessageRenderer_is_implemented_correctly(t *testing.T) {
	g := NewGomegaWithT(t)
	sourceFile, _ := os.OpenFile("message_renderer.go", os.O_RDONLY, os.ModeType)
	sourceCode, _ := ioutil.ReadAll(sourceFile)

	g.Expect(sourceCode, `package lesson_3

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

type BodyRenderer struct {}

func (b BodyRenderer) Render(message Message) string {
	return fmt.Sprintf("<b>%s</b>", message.Body)
}

type HeaderRenderer struct {}

func (b HeaderRenderer) Render(message Message) string {
	return fmt.Sprintf("<h1>%s</h1>", message.Header)
}

type FootRenderer struct {}

func (b FootRenderer) Render(message Message) string {
	return fmt.Sprintf("<i>%s</i>", message.Footer)
}
`, string(sourceCode))
}

func Test_Rendering_a_message(t *testing.T) {
	g := NewGomegaWithT(t)
	sut := NewMessageRenderer()
	message := Message{Header: "h", Body: "b", Footer: "f"}

	html := sut.Render(message)

	g.Expect(html).Should(
		Equal("<h1>h</h1><b>b</b><i>f</i>"))
}
