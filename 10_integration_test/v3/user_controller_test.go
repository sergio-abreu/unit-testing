package v3_test

import (
	"fmt"
	. "github.com/onsi/gomega"
	v2 "sergio/unit-testing/09_external_communication/v2"
	"sergio/unit-testing/09_external_communication/v2/bus"
	"sergio/unit-testing/09_external_communication/v2/company_factory"
	"sergio/unit-testing/09_external_communication/v2/database"
	"sergio/unit-testing/09_external_communication/v2/user"
	"sergio/unit-testing/09_external_communication/v2/user_factory"
	"testing"
)

func Test_Changing_email_from_corporate_to_non_corporate(t *testing.T) {
	g := NewGomegaWithT(t)
	db := database.NewDatabase()
	CreateUser(
		"sergioabreu@ntopus.com.br", user.Employee, db)
	CreateCompany(
		"ntopus.com.br", 1, db)
	spy := NewNatsSpyWithT(t)
	broker := bus.NewBroker(spy)
	sut := v2.NewUserController(db, broker)

	err := sut.ChangeEmail(1, "sergio.vaz.abreu@gmail.com")

	g.Expect(err).Should(
		Not(HaveOccurred()))
	userData, err := db.GetUserById(1)
	g.Expect(err).Should(
		Not(HaveOccurred()))
	userFromDB := user_factory.Create(userData)
	g.Expect(userFromDB.Email()).Should(
		Equal("sergio.vaz.abreu@gmail.com"))
	g.Expect(userFromDB.Group()).Should(
		Equal(user.Customer))
	companyData, err := db.GetCompany()
	g.Expect(err).Should(
		Not(HaveOccurred()))
	companyFromDB := company_factory.Create(companyData)
	g.Expect(companyFromDB.NumberOfEmployees()).Should(
		Equal(0))
	spy.ShouldSendNumberOfMessages(1).
		WithEmailChangedMessage(1, "sergio.vaz.abreu@gmail.com")
}

func CreateUser(email string, group user.UserType, db database.Database) {
	_ = db.SaveUser(user.NewUser(1, email, group, false))
}

func CreateCompany(domainName string, numbersOfEmployees int, db database.Database) {
	_ = db.SaveCompany(user.NewCompany(domainName, numbersOfEmployees))
}

func NewNatsSpyWithT(t *testing.T) *NatsSpy {
	return &NatsSpy{g: NewGomegaWithT(t), calls: make(map[string][]string)}
}

type NatsSpy struct {
	g     *GomegaWithT
	calls map[string][]string
}

func (spy *NatsSpy) Publish(subject, content string) error {
	spy.calls[subject] = append(spy.calls[subject], content)
	return nil
}

func (spy *NatsSpy) ShouldSendNumberOfMessages(number int) *NatsSpy {
	spy.g.Expect(spy.calls["email-changed"]).Should(
		HaveLen(number))
	return spy
}

func (spy *NatsSpy) WithEmailChangedMessage(userId int, newEmail string) *NatsSpy {
	spy.g.Expect(spy.calls["email-changed"]).Should(
		ContainElement(fmt.Sprintf(`{"id":%d,"email":"%s"}`, userId, newEmail)))
	return spy
}
