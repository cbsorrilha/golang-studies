package main

import (
	"fmt"
	// "io/ioutil"
	"log"
	// "net/http"
)


func main() {
	// fmt.Println("start")
	// // it's not like asynchronous in js. It only defer the execution of the statement to the end of the func,
	// // but before it actually returns.
	// // defer is executed in LIFO (last in, first out) order
	// defer fmt.Println("middle")
	// fmt.Println("end")

	// defer fmt.Println("start")
	// defer fmt.Println("middle")
	// defer fmt.Println("end")

	// -=======================================================-
	// res, err := http.Get("https://www.google.com/robos.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // free the resources when the function "ends"
	// defer res.Body.Close()
	// robots, err := ioutil.ReadAll(res.Body)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", robots)

	// -=======================================================-

	// a := "start"
	// // it evaluate earlier, using the value passed to the function when the defer is called
	// defer fmt.Println(a)
	// a = "end"

	// -=======================================================-

	// a, b := 1, 0
	// ans := a / b
	// fmt.Println(ans)

	// fmt.Println("start")
	// defer fmt.Println("defered")
	// panic("something bad happened")
	// fmt.Println("end")

	// -=======================================================-

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	// 	w.Write([]byte("Hello Go!"))
	// })
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// -=======================================================-
	// fmt.Println("start")
	// // defer works even if the function panics to free resources
	// // defer fmt.Println("defered")
	// defer func() {
	// 	if err := recover(); err!=nil {
	// 		log.Println("error:", err)
	// 	}
	// }()
	// panic("something bad happened")
	// fmt.Println("end")
	// -=======================================================-

	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		// with recover, the function don't stop the execution flow
		if err := recover(); err!=nil {
			log.Println("error:", err)
			//panicking again when failed to recover
			panic(err)
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}