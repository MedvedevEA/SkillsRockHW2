package main

import "fmt"

func printNumber(ptrToNumber interface{}) {

	value, ok := ptrToNumber.(*int)

	if !ok {
		fmt.Println("not *int")
	} else if value != nil {
		fmt.Println(*value)
	} else {
		fmt.Println("nil")
	}
}
func main() {
	//корректное значени
	v := 10
	printNumber(&v) //10
	//nil значение
	var v2 *int
	printNumber(v2) //nil
	//значение не *int
	s := "10"
	printNumber(&s) // not *int
	//интерфейс равный nil
	var v3 interface{}
	fmt.Println(v3 == nil) //true
	printNumber(v3)        // not *int
	//интерфейс не равный nil
	var v4 interface{} = v2
	fmt.Println(v4 == nil) //false
	printNumber(v4)        //nil
}
