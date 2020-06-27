package main

import (
	"fmt"
)

// func main() {
// 	var w Writer = ConsoleWriter{}
// 	w.Write([]byte("Hello Go!!"))
// }

// //interfaces don't describe data. Interfaces describe behaviour.
// type Writer interface {
// 	Write([]byte) (int, error)
// }

// // Go doesnt have a explicit implementation
// type ConsoleWriter struct {}

// func (cw ConsoleWriter) Write(data []byte) (int, error) {
// 	n, err := fmt.Println(string(data))
// 	return n, err
// }

func main() {
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i<10;i++ {
		fmt.Println(inc.Increment())
	}
}

//interfaces can be of whatever custom type
type Incrementer  interface{
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}