package main

import (
	"fmt"
	"time"
)

func main() {
	time := time.Now()

	fmt.Println("DATE : ", time.Format("12-02-2022"))
	fmt.Println("Time : ", time.Format("10:04:10"))
}
