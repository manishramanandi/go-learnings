//	Sorts an array by repeatedly finding the minimum elemetn from unsorted part
//
// and putting it at the begginnig
package main

import "fmt"

func main() {
	var scores = []int{30, 10, 55, 33, 11}
	var length = len(scores)

	sort(scores, length)

	for i := 0; i < length; i++ {
		fmt.Printf("%d\n", scores[i])
	}

}

func sort(arrays []int, length int) {

	var minIndex int // save the index of the selected minimum

	for i:=0; i<length-1; i++ {
		minIndex = i 
		// save the minimum value of each loop as the first elment
		var minValue = arrays[minIndex]
	for j:=0; j<length-1; j++ {
		if minValue > arrays[j+1]{
			minValue = arrays[j+1]
			minIndex = j+1
		}
	}

	// if minIndex changed current minimum is exachange with minIndex
	if i != minIndex {
		var temp = arrays[i]
		arrays[i] = arrays[minIndex]
		arrays[minIndex] = temp
	}
	}

}
