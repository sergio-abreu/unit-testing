package database

import (
	"sergio/unit-testing/8_crm/v3/user"
)

type Database struct {}

func (Database) GetUserById(userId int) (map[string]interface{}, error) {
	return nil, nil
}

func (Database) GetCompany() (map[string]interface{}, error) {
	return nil, nil
}

func (Database) SaveCompany(company *user.Company) error {
	return nil
}

func (Database) SaveUser(user *user.User) error {
	return nil
}