package user

import (
	"errors"
	"strings"
)

var InvalidNumberOfEmployeesErr = errors.New("invalid number of employees")

func NewCompany(domainName string, numberOfEmployees int) *Company {
	return &Company{numberOfEmployees: numberOfEmployees, domainName: domainName}
}

type Company struct {
	domainName        string
	numberOfEmployees int
}

func (company *Company) ChangeNumberOfEmployees(delta int) error {
	if company.numberOfEmployees+delta < 0 {
		return InvalidNumberOfEmployeesErr
	}
	company.numberOfEmployees += delta
	return nil
}

func (company *Company) IsEmailCorporate(email string) bool {
	emailDomain := strings.Split(email, "@")[1]
	return emailDomain == company.domainName
}

func (company *Company) DomainName() string {
	return company.domainName
}

func (company *Company) NumberOfEmployees() int {
	return company.numberOfEmployees
}
