// Define a one dimensional array of student scores
// package algorithms

import "fmt"

func main() {
	var scores = []int{90, 75, 80, 60, 20}

	var length = len(scores)
	for i := 0; i < length; i++ {

		fmt.Println(scores[i])
	}
}
