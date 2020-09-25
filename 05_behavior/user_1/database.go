package user_1

func GetUser1FromDatabase(userId int) *User1 {
	return &User1{Name: "Sérgio"}
}

func SaveUser1ToDatabase(user *User1) {
}
