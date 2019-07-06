package localdb

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Log the sync result to log.txt
func WriteLog(s string) {

	f, _ := os.OpenFile("localdb/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()

	w := bufio.NewWriter(f)
	cTime := time.Now().Format("2006-01-02 15:04:05")
	lineStr := fmt.Sprintf("%s > %s\n", cTime, s)
	w.WriteString(lineStr)
	w.Flush()
	fmt.Printf("Write a log to log.txt. Content: %s\n", s)
}
