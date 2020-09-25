package user

import (
	. "github.com/onsi/gomega"
	"testing"
)

type testCase struct {
	companyDomainName string
	email             string
	isEmailCorporate  bool
}

func Test_Differentiates_a_corporate_email_from_non_corporate(t *testing.T) {
	for _, _case := range []testCase{
		{"ntopus.com.br", "sergioabreu@ntopus.com.br", true},
		{"ntopus.com.br", "sergio.vaz.abreu@gmail.com", false},
	} {
		t.Run("Differentiates a corporate email from non corporate", func(t *testing.T) {
			g := NewGomegaWithT(t)
			sut := NewCompany(_case.companyDomainName, 0)

			isEmailCorporate := sut.IsEmailCorporate(_case.email)

			g.Expect(isEmailCorporate).Should(
				Equal(_case.isEmailCorporate))
		})
	}
}

func Test_Company_cannot_have_negative_number_of_employees(t *testing.T) {
	g := NewGomegaWithT(t)
	sut := NewCompany("ntopus.com.br", 0)

	err := sut.ChangeNumberOfEmployees(-1)

	g.Expect(err).Should(
		MatchError(InvalidNumberOfEmployeesErr))
}
