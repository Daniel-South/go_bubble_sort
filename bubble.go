/**********************************************************************************
File:	bubble.go
Title:	Bubble Sort For Integer Arrays
Author:	Daniel R. South, April 2024

Note: 	The bubble sort is more of an exercise than a practical tool. Other sorting
	algorithms are more efficient. The Go language has a sorting function
	built into the slices package.

	That said, someone recently challenged me to code a bubble sort without
	referring to any documentaion.
	
	I enjoy a challenge, so here you go!

	I added a verbose option that shows step by step how the numbers are
	"bubbled up" within the array. For faster results, pass false to defeat
	the verbose option.
**********************************************************************************/

package main

import (
	"fmt"
	"time"
	"math/rand/v2"
)

func main() {
	p := fmt.Println
	p("*** Bubble Sort Demo in Go ***")

	a1 := [] int { 5, 3, 2, 4, 1 }
	bubbleIntArr( a1, true )
	p("\n==> Sorted array: ", a1)

	a2 := [] int { 10, 20, 30, 40 }
	bubbleIntArr( a2, true )
	p("\n==> Sorted array: ", a2)

	a3 := [] int { 123, 99, 66, 11, 0, -22 }
	bubbleIntArr( a3, true )
	p("\n==> Sorted array: ", a3)


	p("\nGenerate an array of random numbers...")
	a4 := genRandIntArray(40, 200)
	p(a4)
	bubbleIntArr( a4, false)
	p("\n==> Sorted array: ", a4)

}

func bubbleIntArr( arr [] int, verbose bool ) {
	p := fmt.Println
	pf := fmt.Printf

	p("\n*** Array before sorting:", arr)
	p("*** Starting Sort At", time.Now())

	var len_minus_one = len(arr) - 1
	var max_index = len_minus_one

	// The outer loop needs to run a maximum of one fewer times
	// than the number of elements in the array.

	for i := 0; i < len_minus_one; i++ {

		var switch_happened = false
		if verbose == true {
			p("Outer loop:", i)
		}

		// The inner loop will run len(arr) - 1 times on the first pass.
		// One each pass it can go one fewer time, because the largest
		// numbers will be moved to the highest positions in the early
		// loops.

		for j := 0; j < max_index; j++ {

			if arr[j] > arr[j+1] {
				if verbose == true {
					pf("(i,j) = (%d %d) %d moves past %d ", i, j, arr[j], arr[j+1])
				}
				var temp_val    = arr[j]
				arr[j]          = arr[j+1]
				arr[j+1]        = temp_val
				switch_happened = true
				if verbose == true {
					p(arr)
				}
			}
			
		}

		// If no switches occured in the inner loop, the array is now sorted.
		// We can break out of the outer loop.

		if switch_happened == false {
			break
		}

		max_index = max_index - 1
	}

	p("*** Finished Sort At", time.Now())
	p("*** Array after sorting:", arr)
}

func genRandIntArray(num_vals int, max_val int) []int {
	var int_arr [] int
	for i := 0; i < num_vals; i++ {
		rand_val := rand.IntN(max_val) - max_val/2
		int_arr = append(int_arr, rand_val)
	}

	return int_arr
}
