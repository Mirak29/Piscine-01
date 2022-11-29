package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		file, _ := ioutil.ReadFile(os.Args[1])
		fmt.Printf(string(file))
	} else if len(os.Args) < 2 {
		fmt.Println("File name missing")
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	}
}
