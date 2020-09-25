package user

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

func (user *User) ChangeEmail(newEmail string, company *Company) error {
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
