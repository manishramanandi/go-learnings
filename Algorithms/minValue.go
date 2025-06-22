// searching the minimam of integer sequence

package main

import "fmt"

func min(arrays []int, length int) int {

	var minIndex = 0
	for j := 0; j < length; j++ {
		if arrays[minIndex] > arrays[j] {
			minIndex = j
		}
	}
	return arrays[minIndex]
}

func main() {
	var scores = []int{40, 30, 70, 44, 22}
	var length = len(scores)
	var minValue = min(scores, length)
	fmt.Printf("min value = %d\n", minValue)

}
