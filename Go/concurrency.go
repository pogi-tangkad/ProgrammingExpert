// Date: 2024-01-15
// Author: Pogi Tangkad

package main

import (
	"fmt"
	"time"
)

// func run1() {
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("run1")
// }

// func run2() {
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("run2")
// }

// func run3() {
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("run3")
// }


func add(a int, b int, rv chan int) {
	time.Sleep(1 * time.Second)
	rv <- a + b
}

func main() {

	fmt.Println("Welcome to concurrency.")

	// go run1()
	// go run2()
	// go run3()
	// time.Sleep(5 * time.Second)

	returnChan := make(chan int)
	go add(1, 4, returnChan)
	rv := <-returnChan
	fmt.Println(rv)


	fmt.Println("done.")



}
