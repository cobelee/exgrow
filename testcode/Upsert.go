package testcode

import (
	"exgrow/localdb"
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Name     string
	Sex      bool
	Phone    string
	Age      int
	Money    float64
	Birthday time.Time
}

func TestUpsert() {
	t, _ := time.Parse("2006-01-02", "1977-08-20")
	fmt.Println(t)
	u := User{"cobe1", true, "13958202267", 40, 6789.12, time.Date(1977, time.August, 21, 0, 0, 0, 0, time.UTC)}
	localdb.Upsert("Test", "User", bson.M{"birthday": t, "name": "cobe1"}, u)
	fmt.Println("更新完成\n")
}
