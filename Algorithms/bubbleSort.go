package main

import "fmt"

func main() {

	var values = []int{30, 10, 40, 80, 22}
	var length = len(values)
	sort(values, length)
	for i := 0; i < length; i++ {
		fmt.Printf("|%d|", values[i])
	}
}

func sort(arrays []int, length int) {
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1; j++ {
			if arrays[j] > arrays[j+1] { //swap
				var flag = arrays[j]
				arrays[j] = arrays[j+1]
				arrays[j+1] = flag
			}
		}
	}
}
