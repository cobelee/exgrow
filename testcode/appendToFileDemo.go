package testcode

import (
	"fmt"
	"os"
	"time"
)

func AppendToFileDemo() {

	f, _ := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	cTime := time.Now().Format("2006-01-02 15:04:05")
	f.WriteString(cTime + "\n")
	f.Close()

	fmt.Println("Write file completed")

}
