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
		fmt.Printf("%d", scores[i])
	}

}

func sort(arrys []int, length int) {

	var minIndex int // save the index of the selected minimum

}
