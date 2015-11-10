package main

import "fmt"

func insertionSort(a []int) {
	for j, v := range a {
		if j == 0 {
			continue
		}
		i := j - 1
		for i >= 0 && a[i] > v {
			a[i+1] = a[i]
			i = i - 1
		}
		a[i+1] = v
	}
}

func main() {
	a := []int{5, 2, 4, 1, 3, 6}
	insertionSort(a)
	fmt.Println(a)
}
