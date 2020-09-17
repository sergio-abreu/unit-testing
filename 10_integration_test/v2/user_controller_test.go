package v2_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	v2 "sergio/unit-testing/9_external_communication/v2"
	"sergio/unit-testing/9_external_communication/v2/bus"
	"sergio/unit-testing/9_external_communication/v2/bus/mocks"
	"sergio/unit-testing/9_external_communication/v2/company_factory"
	"sergio/unit-testing/9_external_communication/v2/database"
	"sergio/unit-testing/9_external_communication/v2/user"
	"sergio/unit-testing/9_external_communication/v2/user_factory"
	"testing"
)

func Test_Changing_email_from_corporate_to_non_corporate(t *testing.T) {
	g := NewGomegaWithT(t)
	ctrl := gomock.NewController(t)
	db := database.NewDatabase()
	CreateUser(
		"sergioabreu@ntopus.com.br", user.Employee, db)
	CreateCompany(
		"ntopus.com.br", 1, db)
	mock := mocks.NewMockNats(ctrl)
	broker := bus.NewBroker(mock)
	sut := v2.NewUserController(db, broker)

	mock.EXPECT().
	Publish("email-changed", `{"id":1,"email":"sergio.vaz.abreu@gmail.com"}`).
	Times(1)

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
	g.Expect(userFromDB.Email()).Should(
		Equal("sergio.vaz.abreu@gmail.com"))
	g.Expect(userFromDB.Group()).Should(
		Equal(user.Customer))
	g.Expect(companyFromDB.NumberOfEmployees()).Should(
		Equal(0))
}

func CreateUser(email string, group user.UserType, db database.Database) {
	_ = db.SaveUser(user.NewUser(1, email, group, false))
}

func CreateCompany(domainName string, numbersOfEmployees int, db database.Database) {
	_ = db.SaveCompany(user.NewCompany(domainName, numbersOfEmployees))
}