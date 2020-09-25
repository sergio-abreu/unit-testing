package user_factory

import (
	"sergio/unit-testing/08_code_types/v3/user"
)

func Create(data map[string]interface{}) *user.User {
	id := data["id"].(int)
	email := data["email"].(string)
	userType := data["type"].(user.UserType)
	return user.NewUser(id, email, userType)
}
