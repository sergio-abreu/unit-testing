package v1

import (
	"sergio/unit-testing/09_external_communication/v1/bus"
	"sergio/unit-testing/09_external_communication/v1/company_factory"
	"sergio/unit-testing/09_external_communication/v1/database"
	"sergio/unit-testing/09_external_communication/v1/user_factory"
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
	err = user.CanChangeEmail()
	if err != nil {
		return err
	}
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
	// Evento está sendo enviado mesmo que o email não tenha sido alterado
	return ctrl.bus.SendEmailChangeMessage(userID, newEmail)
}
