package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	for _, v := range arguments[1:] {
		if v == "01" || v == "galaxy" || v == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
