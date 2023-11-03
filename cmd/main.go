package main

import (
	"fmt"
	"os"

	"digit_generator.com/internal"
)

func main() {
	solutions, err := internal.FindSolutions("9876543210")
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, solution := range solutions {
		fmt.Println(solution)
	}
}
