package user_2

import "strings"

type User2 struct {
	name string
}

func (u *User2) normalizeName(name string) string {
	result := strings.Trim(name, "")
	if len(result) > 50 {
		result = result[:50]
	}
	return result
}

func (u *User2) Rename(newName string) {
	u.name = u.normalizeName(newName)
}
