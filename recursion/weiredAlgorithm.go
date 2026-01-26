package main

import "fmt"

func weired(n int) {

	fmt.Printf("%d ", n)
	if n == 1 {
		return
	}

	if n%2 == 0 {
		n = n / 2
	} else {
		n = n*3 + 1
	}

	weired(n)
}
