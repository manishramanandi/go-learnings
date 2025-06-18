// maximum of integre sequence

package main

import "fmt"

func max(array []int, length int) int {

	for i := 0; i < length-1; i++ {

		if array[i] > array[i+1] {
			// swap
			var temp = array[i]
			array[i] = array[i+1]
			array[i+1] = temp
		}
	}
	var maxValue = array[length-1]
	return maxValue
}

func main() {
	var score = []int{30, 80, 30, 22, 11}
	var length = len(score)
	var maxValue = max(score, length)
	fmt.Println("max value", maxValue)
}
