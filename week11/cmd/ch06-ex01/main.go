package main

import "fmt"

func main() {
	subjects := [4]string{"GO", "Javascript", "Python", "Linux"}
	subjectsSlice := subjects[:3] //slicing
	subjects[0] = "Java"
	subjectsSlice = append(subjectsSlice, "GO") // list.append("GO")
	//subjectsSlice = append(subjectsSlice, "GO", "DB")
	for _, subject := range subjects {
		fmt.Println((subject))
	}
	fmt.Println("=================")
	for i := 0; i < len(subjectsSlice); i++ {
		fmt.Println(subjectsSlice[i])
	}
}
