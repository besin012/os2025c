package main

import "fmt"

func main() {
	subjects := make([]string, 3)
	subjects[0] = "GO"
	subjects[2] = "python"

	for _, subject := range subjects {
		fmt.Println((subject))
	}
}
