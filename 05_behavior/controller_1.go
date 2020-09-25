package _5_behavior

import (
	. "sergio/unit-testing/05_behavior/user_1"
)

type UserController1 struct{}

func (u UserController1) RenameUser(userId int, newName string) {
	user := GetUser1FromDatabase(userId)

	normalizedName := user.NormalizeName(newName)
	user.Name = normalizedName

	SaveUser1ToDatabase(user)
}
