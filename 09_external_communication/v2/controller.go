package v2

import (
	"sergio/unit-testing/09_external_communication/v2/bus"
	"sergio/unit-testing/09_external_communication/v2/company_factory"
	"sergio/unit-testing/09_external_communication/v2/database"
	"sergio/unit-testing/09_external_communication/v2/user_factory"
)

func NewUserController(database database.Database, bus bus.IMessageBus) UserController {
	return UserController{database: database, bus: bus}
}

type UserController struct {
	database database.Database
	bus      bus.IMessageBus
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
	for _, event := range user.Events() {
		err = ctrl.bus.SendEmailChangeMessage(event.UserID, event.NewEmail)
		if err != nil {
			return err
		}
	}
	return nil
}
