package main

import "fmt"

func main() {
	var pi float64 = 3.142
	var d int = 24
	r := int(d) / 2
	fmt.Println("Radius of circle is", r)

	//area of circle
	area := pi * float64(r) * float64(r)
	fmt.Println(area)

	//area of circle
	perimeter := 2 * pi * float64(r)
	fmt.Println(perimeter)
}

/*import "fmt"

const PI float64 = 3.14 // global constant

func main() {
    var radius float64 = 5.154
    var area float64

    area = PI * radius * radius
    fmt.Printf("Radius: %.2f \nPI:%.1f \n", radius, PI)
    fmt.Printf("Area of Circle is : %f", area)
    fmt.Println("Thank You")
*/
