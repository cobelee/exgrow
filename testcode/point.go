package testcode

import (
	"encoding/json"
	"fmt"
)

func TestPoint() {
	// var u = user{"cobe", "1395202267"}
	var m, n map[string]user
	m = make(map[string]user)
	m["1"] = user{"cobe", "13958202267"}
	m["2"] = user{"sainna", "13968321211"}

	var i int = 3
	fmt.Printf("The value of i %v\n", i)
	fmt.Printf("The address of i %p\n", &i)

	var pi *int
	fmt.Printf("The address of pi is %p\n", &pi)
	fmt.Printf("The pi value is %v\n", pi)
	pi = &i
	fmt.Printf("The address of pi is %p\n", &pi)
	fmt.Printf("The pi value is %v\n", pi)

	fmt.Printf("-----------------------\n")
	fmt.Printf("The value of m %v\n", m)
	fmt.Printf("The address of m %p\n", m)
	fmt.Printf("The address of &m %p\n", &m)
	fmt.Printf("-----------------------\n")

	var pm *map[string]user
	fmt.Printf("The address of pm is %p\n", &pm)
	fmt.Printf("The value of pm is %p\n", pm)
	pm = &m
	fmt.Printf("pm = &m\n")
	fmt.Printf("The address of pm is %p\n", &pm)
	fmt.Printf("The value of pm is %p\n", pm)
	fmt.Printf("--------------------------\n")

	strJson, _ := json.Marshal(m)
	fmt.Printf("The json string of m is %v\n", string(strJson))
	fmt.Printf("%v\n", m)

	if n == nil {
		fmt.Printf("The n is nil\n")
	}

	fmt.Printf("The value of n is %v\n", n)
	fmt.Printf("The address of n is %p\n", &n)
	json.Unmarshal(strJson, &n)
	fmt.Printf("json.Unmarshal\n")
	fmt.Printf("The address of n is %p\n", &n)
	fmt.Printf("The address of n is %p\n", n)

	if n != nil {
		fmt.Printf("The n is unnil\n")
	}

	fmt.Printf("%v\n", n)
}
