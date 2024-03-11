package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type contactInfo struct {
	Name  string
	Email string
}

type purchaseInfo struct {
	Name   string
	Price  float32
	Amount int
}

func main() {
	var contacts []contactInfo = loadJSON[contactInfo]("./contactInfo.json")
	fmt.Printf("%+v\n", contacts)

	var purchases []purchaseInfo = loadJSON[purchaseInfo]("./purchaseInfo.json")
	fmt.Printf("%+v\n", purchases)
}

func loadJSON[T contactInfo | purchaseInfo](filePath string) []T {
	var data, _ = ioutil.ReadFile(filePath)

	var loaded []T
	json.Unmarshal(data, &loaded)

	return loaded
}
