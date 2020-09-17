package user_1

import "strings"

type User1 struct {
	Name string
}

func (u *User1) NormalizeName(name string) string {
	result := strings.Trim(name, "")
	if len(result) > 50 {
		result = result[:50]
	}
	return result
}
