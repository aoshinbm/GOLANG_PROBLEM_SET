package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//check first file file exist or not
	//Stat() function is used to return the file info structure describing the file
	if _, err := os.Stat("Cricket_Players_Stats.csv"); err == nil {
		fmt.Println("File exists...")
	} else {
		fmt.Println("File does not exist!!!!!")
	}

	fmt.Println("----------------------------------------------")

	// open file
	f, err := os.Open("Cricket_Players_Stats.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	for {
		playerData, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", playerData)
	}
}
