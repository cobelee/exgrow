package testcode

import (
	"fmt"
)

const (
	Sunday = iota
	Monday
	Tuesday
	Wedenesday
	Thursday
	Friday
	Saturday
)

func ShowDay() {

	fmt.Println("Today is :%v", Sunday)
}
