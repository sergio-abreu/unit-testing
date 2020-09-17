package user

import (
	. "github.com/onsi/gomega"
	"testing"
)

func Test_Changing_email_from_non_corporate_to_corporate(t *testing.T) {
	g := NewGomegaWithT(t)
	company := NewCompany("ntopus.com.br", 1)
	sut := NewUser(1, "sergio.vaz.abreu@gmail.com", Customer, false)

	err := sut.ChangeEmail("sergioabreu@ntopus.com.br", company)

	g.Expect(err).Should(
		Not(HaveOccurred()))
	g.Expect(company.numberOfEmployees).Should(
		BeEquivalentTo(2))
	g.Expect(sut.email).Should(
		Equal("sergioabreu@ntopus.com.br"))
	g.Expect(sut.group).Should(
		BeEquivalentTo(Employee))
}

func Test_Changing_email_from_corporate_to_non_corporate(t *testing.T) {}
func Test_Changing_email_without_changing_user_type(t *testing.T) {}
func Test_Changing_email_to_the_same_one(t *testing.T) {}