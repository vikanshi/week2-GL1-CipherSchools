package main

import (
	"log"
	"os"
	"fmt"
	"io/ioutil"
	
)

type Something struct {
	Name    string    `json:"name"`
	Married bool      `json:"married"`
	Age     float64   `json:"age"`
	Address []Address `json:"address"`
	Marks   []int     `json:"marks"`
}
type Address struct {
	Street int    `json:"street"`
	City   string `json:"city"`
}

func main () {
	jsonFile, err := os.Open("something.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	jsonByteValues, err  := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonByteValues))
	// var something Something
	// json.Unmarshal(jsonByteValues, &something)
	// log.Println(something)

	// newJsonByteValues, err := json.Marshal(something)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// os.WriteFile("some.json", newJsonByteValues)
}