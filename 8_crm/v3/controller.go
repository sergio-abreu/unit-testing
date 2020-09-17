package v3

import (
	"sergio/unit-testing/8_crm/v3/bus"
	"sergio/unit-testing/8_crm/v3/company_factory"
	"sergio/unit-testing/8_crm/v3/database"
	"sergio/unit-testing/8_crm/v3/user_factory"
)

func NewUserController() UserController {
	return UserController{database: database.Database{}, bus: bus.MessageBus{}}
}

type UserController struct {
	database database.Database
	bus      bus.MessageBus
}

func (ctrl UserController) ChangeEmail(userID int, newEmail string) error {
	userData, err := ctrl.database.GetUserById(userID)
	if err != nil {
		return err
	}
	user := user_factory.Create(userData)
	companyData, err := ctrl.database.GetCompany()
	if err != nil {
		return err
	}
	company := company_factory.Create(companyData)
	err = user.ChangeEmail(
		newEmail, company)
	if err != nil {
		return err
	}
	err = ctrl.database.SaveCompany(company)
	if err != nil {
		return err
	}
	err = ctrl.database.SaveUser(user)
	if err != nil {
		return err
	}
	return ctrl.bus.SendEmailChangeMessage(userID, newEmail)
}
