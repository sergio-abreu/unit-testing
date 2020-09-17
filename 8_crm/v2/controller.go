package v2

import (
	"sergio/unit-testing/8_crm/v2/bus"
	"sergio/unit-testing/8_crm/v2/database"
)

func NewUserController() UserController {
	return UserController{database: database.Database{}, bus: bus.MessageBus{}}
}

type UserController struct {
	database database.Database
	bus      bus.MessageBus
}

func (ctrl UserController) ChangeEmail(userID int, newEmail string) error {
	data, err := ctrl.database.GetUserById(userID)
	if err != nil {
		return err
	}
	id := data["id"].(int)
	email := data["email"].(string)
	userType := data["type"].(UserType)
	user := NewUser(id, email, userType)
	companyData, err := ctrl.database.GetCompany()
	if err != nil {
		return err
	}
	companyDomainName := companyData["companyDomainName"].(string)
	numberOfEmployees := companyData["numberOfEmployees"].(int)
	newNumberOfEmployees := user.ChangeEmail(
		newEmail, companyDomainName, numberOfEmployees)
	err = ctrl.database.SaveCompany(newNumberOfEmployees)
	if err != nil {
		return err
	}
	err = ctrl.database.SaveUser(user)
	if err != nil {
		return err
	}
	return ctrl.bus.SendEmailChangeMessage(userID, newEmail)
}
