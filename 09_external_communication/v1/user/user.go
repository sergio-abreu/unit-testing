package user

import "errors"

type UserType int

var CannotChangeConfirmedEmailErr = errors.New("can't change a confirmed email")

const (
	Customer UserType = 1
	Employee          = 2
)

func NewUser(userID int, email string, group UserType, isEmailConfirmed bool) *User {
	return &User{userID: userID, email: email, group: group, isEmailConfirmed: isEmailConfirmed}
}

type User struct {
	userID           int
	email            string
	group            UserType
	isEmailConfirmed bool
}

func (user *User) CanChangeEmail() error {
	if user.isEmailConfirmed {
		return CannotChangeConfirmedEmailErr
	}
	return nil
}

func (user *User) ChangeEmail(newEmail string, company *Company) error {
	err := user.CanChangeEmail()
	if err != nil {
		return err
	}
	if user.email == newEmail {
		return nil
	}
	newGroup := Customer
	if company.IsEmailCorporate(newEmail) {
		newGroup = Employee
	}
	if user.group != newGroup {
		delta := -1
		if newGroup == Employee {
			delta = 1
		}
		err := company.ChangeNumberOfEmployees(delta)
		if err != nil {
			return err
		}
	}
	user.email = newEmail
	user.group = newGroup
	return nil
}
