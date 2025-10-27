package main

import (
	"fmt"
)

func swap(first *int, second *int) {
	var temp int = 0
	temp = *first
	*first = *second
	*second = temp
	fmt.Println(*first, *second)
}
func main() {
	a, b := 10, 20
	fmt.Println(a, b)
	swap(&a, &b) // call by pointer
	fmt.Println(a, b)
	//fmt.Printf("%0.3f\n", math.Sqrt(-9.3))
}

// if score >= 60 {
// 	fmt.Println(score, "Pass")
// } else {
// 	fmt.Println(score, "Fail")
// }

// shadowing

// var float64 float64 = 2.71
// var f float64 = 3.991
// fmt.Println(float64)
// fmt.Println(f)

// var fmt float64 = 2.71
// fmt.Println(fmt)
//}
