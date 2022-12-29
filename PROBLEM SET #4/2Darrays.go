package main

import "fmt"

func main() {
	/* an array with 5 rows and 2 columns*/
	var a = [5][2]int{}
	a[0][0] = 10
	a[1][0] = 12
	a[2][0] = 15
	a[3][0] = 19
	a[4][0] = 24

	var i, j int

	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			a[i][1] = a[i][0] * 2
		}
	}

	/* output each array element's value */
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}

}
