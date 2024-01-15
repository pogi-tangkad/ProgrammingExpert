// Date: 2024-01-15
// Author: Pogi Tangkad

package main

import (
	"fmt"
)

// type Name interface {
// 	GetName() string
// 	//LastName(x string) string
// }

// type Person struct {
// 	firstName string
// 	lastName string
// }

// func (p Person) GetName() string {
// 	return p.firstName + " " + p.lastName
// }

// type Employee struct {
// 	name string
// }

// func (e Employee) GetName() string {
// 	return e.name
// }

func main() {

		
	// names := []Name{Employee{"Joe"}, Person{"Tim", "Ruscica"}}

	// fmt.Println(names)
	// for _, name := range names {
	// 	fmt.Println(name.GetName())
	// }

	fmt.Println("no error")
	// defer fmt.Println("defer")
	// fmt.Println("after defer")

	// panic("i failed")
	// fmt.Println("goodbye")
	
	// defer handlePanic()
	// fmt.Println("start")
	// doPanic()
	// fmt.Println("end")

	// result, err := divide(1, 0)
	// fmt.Println(result, err)

	// arr1 :=	[]int{1, 2, 3, 4, 5}
	// arr2 := []int{1, 2, 3, 4, 5}
	// xyz(&arr1, arr2)
	// fmt.Println(arr1, arr2)

	// jkl(arr2)
	// fmt.Println(arr2)



}

// func jkl(arr []int) {
// 	arr[1] = 8
// }

// func xyz(a1 *[]int, a2 []int) {
// 	(*a1)[0] = 10
// 	a2 = *a1
// }

// func divide(a int, b int) (int, error) {
// 	if b == 0 {
// 		return 0, errors.New("do not divide by zero, silly")
// 	}
// 	return a / b, nil
// }

// func doPanic() {
// 	panic("fail")
// }

// func handlePanic() {
// 	r := recover()
// 	if r != nil {
// 		fmt.Println(r)
// 	}
// }