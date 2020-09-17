package user_factory

import (
	"sergio/unit-testing/9_external_communication/v2/user"
)

func Create(data map[string]interface{}) *user.User {
	id := data["id"].(int)
	email := data["email"].(string)
	userType := data["type"].(user.UserType)
	isEmailConfirmed := data["isEmailConfirmed"].(bool)
	return user.NewUser(id, email, userType, isEmailConfirmed)
}