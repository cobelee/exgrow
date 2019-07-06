package testcode

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name  string
	Phone string
}

func testJson() {
	// var u = user{"cobe", "1395202267"}
	var m, n map[string]user
	m = make(map[string]user)
	m["1"] = user{"cobe", "13958202267"}
	m["2"] = user{"sainna", "13968321211"}

	strJson, _ := json.Marshal(m)
	fmt.Printf("The json string of m is %v\n", string(strJson))
	fmt.Printf("%v\n", m)

	if n == nil {
		fmt.Printf("The n is nil\n")
	}

	fmt.Printf("The address of n is %p\n", n)
	json.Unmarshal(strJson, &n)
	fmt.Printf("The address of n is %p\n", n)

	if n != nil {
		fmt.Printf("The n is unnil\n")
	}

	fmt.Printf("%v\n", n)
}
