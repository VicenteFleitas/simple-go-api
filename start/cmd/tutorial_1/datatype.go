package main

import "fmt"

func main() {
	var intNum int = 32767
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1/intNum2);
	fmt.Println(intNum1%intNum2);

	var myString string = `Hello world!`;
	fmt.Println(myString);

	var myBoolean bool = false;
	fmt.Println(myBoolean);

	var var1, var2 = 1, 2;
	fmt.Println(var1 + var2);

	const myConst string = "const value";
	fmt.Println(myConst);
}