package v2

import (
	"strings"
)

type UserType int

const (
	Customer UserType = 1
	Employee          = 2
)

func NewUser(userID int, email string, group UserType) *User {
	return &User{userID: userID, email: email, group: group}
}

type User struct {
	userID int
	email  string
	group  UserType
}

func (user *User) ChangeEmail(newEmail, companyDomainName string, numberOfEmployees int) int {
	if user.email == newEmail {
		return numberOfEmployees
	}
	emailDomain := strings.Split(newEmail, "@")[1]
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
	}
	user.email = newEmail
	user.group = newGroup
	return numberOfEmployees
}
