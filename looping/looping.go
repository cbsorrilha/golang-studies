package main

import (
	"fmt"
)

func main() {
	// i++ is a statement, not an expression. It doesn't resolve into a variable
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	j := 0 
	//The golang equivalent to a while loop
	for ; j < 5 ; {
		fmt.Println(j)
		j++
	}

	k := 0

	for {
		fmt.Println(k)
		k++
		if k == 5 {
			//exiting programatically
			break
		}
	}

	s := []int{1,2,3}
	for k, v := range s {
		fmt.Println("k", k)
		fmt.Println("v", v)
	}
}