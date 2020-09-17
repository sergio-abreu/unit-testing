package Assert

import . "github.com/golang/mock/gomock"

func Equal(expected, actual interface{}) {
	if Not(Eq(expected)).Matches(actual) {
		panic("not equal")
	}
}
