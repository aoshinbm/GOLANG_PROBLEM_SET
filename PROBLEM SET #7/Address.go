package main

import (
	"encoding/json"
	"fmt"
)

type Friends struct {
	Name, Address, City      string
	Mobile_No, Alt_Mobile_No int
}

func main() {
	fmt.Println("Creating Address Book of MyFriends:")
	friend1 := Friends{"Vaishali", "2nd Tree Street", "Mumbai", 123789456, 9897745687}
	friend2 := Friends{"Roma", "Nehru Street", "Virar", 654789123}
	friend3 := Friends{"Rajesh", "Gandhi road", "Pune", 369852147, 852147963}
	friend4 := Friends{"John", "Gokuldham", "Nasik", 1597536482, 6842351759}
	friend5 := Friends{"Bindu", "MIDC", "Nagpur", 9987456123}
	friend6 := Friends{"Alvin", "2nd Tree Street", "Mumbai", 9988445566, 7748596321}
	friend7 := Friends{"Tijo", "Burgundy Road", "Mumbai", 8845679123}
	friend8 := Friends{"Jaison", "2nd Tree Street", "Daman", 9897745687, 123789456}
	friend9 := Friends{"Alinda", "Oberoi Streets", "Mumbai", 9986754123, 7715987455}
	friend10 := Friends{"Christy", "Goregaon", "Virar", 9222248796, 9048756321}
	frnz := []Friends{
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
	for _, frn := range frnz {
		fmt.Println(frn)
	}
	fmt.Println("")

	fmt.Printf("\nConverting slice of structs to JSON \n")
	bytesStructSlice, _ := json.MarshalIndent(Friends, " ", "\t")
	json.MarshalIndent(Friends, " ", "\t")
	fmt.Println(string(bytesStructSlice))

}
