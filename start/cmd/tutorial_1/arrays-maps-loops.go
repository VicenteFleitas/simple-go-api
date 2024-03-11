package main

import "fmt"

func main() {
	var intArr [3]int32 = [3]int32{11, 12, 13}
	fmt.Println(intArr)

	var intSlice []int32 = []int32{4, 5, 6}
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v \n", len(intSlice), cap(intSlice))

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])
	var age, ok = myMap2["Jason"]
	// delete(myMap2, "Adam")
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}
	// iterate map
	for name, age := range myMap2 {
		fmt.Printf("Name: %v - Age: %v\n", name, age)
	}
	// iterate arr
	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}
	// old school
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
