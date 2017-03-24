package main

/*
 * Heap sort - https://en.wikipedia.org/wiki/Heapsort
 */

import (
	"math/rand"  
	"time" 
	"fmt"
)

func sift(arr []int, i int, arrLen int) []int {
	done := false

	tmp := 0
	maxChild := 0

	for (i*2+1 < arrLen) && (!done) {
		if i*2+1 == arrLen-1 {
			maxChild = i*2 + 1
		} else if arr[i*2+1] > arr[i*2+2] {
			maxChild = i*2 + 1
		} else {
			maxChild = i*2 + 2
		}

		if arr[i] < arr[maxChild] {
			tmp = arr[i]
			arr[i] = arr[maxChild]
			arr[maxChild] = tmp
			i = maxChild
		} else {
			done = true
		}
	}

	return arr
}

func main() {
	arr := []int{}
	rand.Seed(time.Now().UnixNano())  
	for i := 0; i < 10; i++ {  
        	arr = append(arr, rand.Intn(100))
    	} 
	fmt.Println("Initial array is:", arr)
	fmt.Println("")

	i := 0
	tmp := 0

	for i = len(arr)/2 - 1; i >= 0; i-- {
		arr = sift(arr, i, len(arr))
	}

	for i = len(arr) - 1; i >= 1; i-- {
		tmp = arr[0]
		arr[0] = arr[i]
		arr[i] = tmp
		arr = sift(arr, 0, i)
	}

	fmt.Println("Sorted array is: ", arr)
}
