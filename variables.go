package main

import (
	"fmt"
)

type Doctor struct {
	number int
	actorName string
	companions []string
}

func main() {
	aDoctor := Doctor{
		number: 3,
		actorName: "Jon Pertwee",
		companions: []string {
			"Liz Shaw",
			"Jo Grant",
		},
	}

	fmt.Println(aDoctor)
}