package main

import (
	"fmt"
	"sync"
	"runtime"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	// go sayHello()
	// time.Sleep(100 * time.Millisecond)
	// var msg = "Hello"
	// //Go has the concept of closures
	// go func() {
	// 	fmt.Println(msg)
	// }()
	// //Go will not interrupt the main thread until the Sleep method, so the msg will be modified before the execution ends
	// msg = "Goodbye"

	// go func(msg string) {
	// 	fmt.Println(msg)//will print Goodbye
	// }(msg)

	// msg = "Guten Morgen"

	// time.Sleep(100 * time.Millisecond)


	// Waiting Groups make the main thread wait for the execution of a given number of goroutines.
	// var msg = "Hello"
	// wg.Add(1)
	// go func(msg string) {
	// 	fmt.Println(msg)
	// 	wg.Done()
	// }(msg)
	// msg = "Goodbye"
	// wg.Wait()

	// this prints the hello message out of order
	// for i := 0; i < 10; i++ {
	// 	wg.Add(2)
	// 	go sayHello()
	// 	go increment()
	// }
	// wg.Wait()

		runtime.GOMAXPROCS(100)
		for i := 0; i < 10; i++ {
			wg.Add(2)
			m.RLock()
			go sayHello()
			m.Lock()
			go increment()
		}
		wg.Wait()
		
		//Returns the number of threads we have available and can set the number of threads to we work with
		fmt.Printf("Threads: %vn", runtime.GOMAXPROCS(-1))


		//Best Practices with go routines

		// Don't create goroutines in libraries: Let the consumer control concurrency

		// When creating a goroutine, know how it will end: Avoid subtle memory leaks

		// Check for race conditions at compile time go run -race src/main.go
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	
	counter++
	m.Unlock()
	wg.Done()
}