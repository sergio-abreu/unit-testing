package user

import "errors"

type UserType int

var CannotChangeConfirmedEmailErr = errors.New("can't change a confirmed email")

const (
	Customer UserType = 1
	Employee UserType = 2
)

type EmailChangedEvent struct {
	UserID   int
	NewEmail string
}

func NewUser(userID int, email string, group UserType, isEmailConfirmed bool) *User {
	return &User{userID: userID, email: email, group: group, isEmailConfirmed: isEmailConfirmed}
}

type User struct {
	userID           int
	email            string
	group            UserType
	isEmailConfirmed bool
	events           []EmailChangedEvent
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
	user.events = append(user.events, EmailChangedEvent{UserID: user.userID, NewEmail: newEmail})
	return nil
}

func (user *User) Events() []EmailChangedEvent {
	return user.events
}

func (user *User) IsEmailConfirmed() bool {
	return user.isEmailConfirmed
}

func (user *User) Group() UserType {
	return user.group
}

func (user *User) Email() string {
	return user.email
}

func (user *User) UserID() int {
	return user.userID
}
