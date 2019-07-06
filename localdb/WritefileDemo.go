package localdb

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

/*
	Two method of creating file.
	Just test program during Dev.
*/
func WriteFileDemo() {

	// Method 1:  Creating file by ioutil.WriteFile()
	data := []byte("Hello World!\n")
	err := ioutil.WriteFile("data1.txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("The first file is created . ")

	//  --------------------------------------------
	// Method 2: Creating file by bufio & fmt.Fprintf()
	f, _ := os.Create("./stockCode1")
	defer f.Close()

	w := bufio.NewWriter(f)
	lineStr := "This is a new world."
	_, err3 := fmt.Fprintf(w, lineStr)
	if err3 != nil {
		fmt.Println(err3)
	}

	w.Flush()
	fmt.Println("The second file writted")

}
