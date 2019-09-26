package testcode

import (
	"fmt"
)

type s struct {
	name string
	age  int
}

func TestInterface() {
	var s1 s
	s1 = s{"sainna", 42}
	GetValue(s1)
	fmt.Println(s1)

}

func GetValue(ifv interface{}) {
	var s2 = s{"cobe", 43}

	ifv = s2
}
