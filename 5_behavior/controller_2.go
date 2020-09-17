package __behavior

import (
	. "sergio/unit-testing/5_behavior/user_2"
)

type UserController2 struct {}

func (u UserController2) RenameUser(userId int, newName string) {
	user := GetUser2FromDatabase(userId)

	user.Rename(newName)

	SaveUser2ToDatabase(user)
}

