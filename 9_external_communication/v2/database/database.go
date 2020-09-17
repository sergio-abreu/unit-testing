package database

import (
	"sergio/unit-testing/9_external_communication/v2/user"
)

func NewDatabase() Database {
	return Database{user: make(map[string]interface{}), company: make(map[string]interface{})}
}

type Database struct {
	user map[string]interface{}
	company map[string]interface{}
}

func (d Database) GetUserById(userId int) (map[string]interface{}, error) {
	return d.user, nil
}

func (d Database) GetCompany() (map[string]interface{}, error) {
	return d.company, nil
}

func (d Database) SaveCompany(company *user.Company) error {
	d.company["domainName"] = company.DomainName()
	d.company["numberOfEmployees"] = company.NumberOfEmployees()
	return nil
}

func (d Database) SaveUser(user *user.User) error {
	d.user["id"] = 1
	d.user["email"] = user.Email()
	d.user["type"] = user.Group()
	d.user["isEmailConfirmed"] = user.IsEmailConfirmed()
	return nil
}