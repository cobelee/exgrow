package localdb

import (
	"bufio"
	"fmt"
	"strings"

	// "fmt"
	"io"

	// "io/ioutil"
	"os"
)

func ReadFile() {

	ReadLine("data/sz002623.csv", processLine)
}

func ReadLine(filePath string, hookfn func([]byte)) error {
	f, e := os.Open(filePath)
	if e != nil {
		return e
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	for i := 0; i < 5; i++ {
		line, err := reader.ReadBytes('\n')

		hookfn(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}

func processLine(line []byte) {
	var dayData []string
	dayData = strings.Split(string(line), ",")
	fmt.Println(dayData)
	fmt.Println(len(dayData))
}
