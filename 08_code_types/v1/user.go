package v1

import (
	"sergio/unit-testing/08_code_types/v1/bus"
	"sergio/unit-testing/08_code_types/v1/database"
	"strings"
)

type UserType int

const (
	Customer UserType = 1
	Employee          = 2
)

type User struct {
	userID int
	email  string
	group  UserType
}

func (user *User) ChangeEmail(userID int, newEmail string) error {
	data, err := database.GetUserById(userID)
	if err != nil {
		return err
	}
	user.userID = userID
	user.email = data["email"].(string)
	user.group = data["type"].(UserType)
	if user.email == newEmail {
		return nil
	}
	companyData, err := database.GetCompany()
	if err != nil {
		return err
	}
	emailDomain := strings.Split(newEmail, "@")[1]
	companyDomainName := companyData["companyDomainName"].(string)
	numberOfEmployees := companyData["numberOfEmployees"].(int)
	isEmailCorporate := companyDomainName == emailDomain
	newGroup := Customer
	if isEmailCorporate {
		newGroup = Employee
	}
	if user.group != newGroup {
		if newGroup == Customer {
			numberOfEmployees--
		} else {
			numberOfEmployees++
		}
		err = database.SaveCompany(numberOfEmployees)
		if err != nil {
			return err
		}
	}
	user.email = newEmail
	user.group = newGroup
	err = database.SaveUser(user)
	if err != nil {
		return err
	}
	return bus.SendEmailChangeMessage(userID, newEmail)
}
