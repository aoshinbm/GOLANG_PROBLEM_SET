package main

import "fmt"

// Defining a struct type
type Address struct {
	Name         string
	MobileNo     string
	Alt_MobileNo string
	Addres       string
	City         string
}

func main() {

	fmt.Println("Creating Address Book of MyFriends:")
	friend1 := Address{"Vaishali", "123789456", "9897745687", "2nd Tree Street", "Mumbai"}
	friend2 := Address{"Roma", "654789123", "", "Nehru Street", "Virar"}
	friend3 := Address{"Rajesh", "369852147", "852147963", "Gandhi road", "Pune"}
	friend4 := Address{"John", "1597536482", "6842351759", "Gokuldham", "Nasik"}
	friend5 := Address{"Bindu", "9987456123", "", "MIDC", "Nagpur"}
	friend6 := Address{"Alvin", "9988445566", "7748596321", "2nd Tree Street", "Mumbai"}
	friend7 := Address{"Tijo", "8845679123", "", "Burgundy Road", "Mumbai"}
	friend8 := Address{"Jaison", "9897745687", "123789456", "2nd Tree Street", "Daman"}
	friend9 := Address{"Alinda", "9986754123", "7715987455", "Oberoi Streets", "Mumbai"}
	friend10 := Address{"Christy", "9222248796", "9048756321", "Goregaon", "Virar"}
	frnz := []Address{
		friend1,
		friend2,
		friend3,
		friend4,
		friend5,
		friend6,
		friend7,
		friend8,
		friend9,
		friend10,
	}
	/*	for _, frn := range frnz {
			fmt.Println(frn)
		}
	*/
	for _, frnd := range frnz {
		fmt.Printf("%s ==> %s \n", frnd.Name, frnd.MobileNo)
	}
	fmt.Println("")

	for _, frnd := range frnz {
		if frnd != (Address{}) {
			fmt.Printf("\n%s ==> %s , %s", frnd.Name, frnd.MobileNo, frnd.Alt_MobileNo)
		}
	}
	fmt.Println("")
	// Declaring a map
	var friends map[Address]int

	// Checking if the map is empty or not
	if friends == nil {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	fmt.Println("\n")
	// Declaring and initialising
	// using map literals
	sample := map[Address]int{friend1: 1, friend2: 2, friend3: 3, friend4: 4, friend5: 5, friend6: 6,
		friend7: 7, friend8: 8, friend9: 9, friend10: 10}
	fmt.Println(sample)
	fmt.Println("\n")

	type Address struct {
		collegeFriend string
	}

}
