package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	fmt.Println("Creating a map Of INDIAN STATES")
	//	var stateCapitals = make(map[string]string)
	stateCapital := make(map[string]string)

	stateCapital["MAHARASHTRA"] = "MUMBAI"
	stateCapital["TAMILNADU"] = "Chennai"
	stateCapital["KERALA"] = "TRIVANDRUM"
	stateCapital["GOA"] = "Panaji"
	stateCapital["KARNATAKA"] = "BENGALURU"
	stateCapital["J&K"] = "SRINAGAR"
	stateCapital["RAJASTHAN"] = "JAIPUR"
	stateCapital["MADHYA_PRADESH"] = "Bhopal"
	stateCapital["BIHAR"] = "Patna"
	stateCapital["PUNJAB"] = "Chandigarh"
	for state, capital := range stateCapital {
		fmt.Printf("Capital of %s is %s \n", state, capital)
	}
	fmt.Println("***************************")

	fmt.Println("Converting into json")
	bytesStructSlice, _ := json.MarshalIndent(stateCapital, " ", "\t")
	json.MarshalIndent(stateCapital, " ", "\t")
	fmt.Println(string(bytesStructSlice))

	fmt.Println("***************************")

}
