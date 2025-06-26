// Add a score 75 at the end of the one dimensional array scores

package main

import "fmt"

func main()  {
	
	var scores = []int{20, 30, 50, 40, 70, 10}
	var length = len(scores)

	var tempArray = make([]int, length+1) // create a new array

	for i:=0; i<length; i++{
		tempArray[i] = scores[i]
	}
	tempArray[length] = 75

	scores = tempArray
	for i:=0;i<length+1;i++{
	fmt.Printf("%d ", scores[i])	
	}
}
