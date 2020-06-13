package main

import (
	"fmt"
	"reflect"
)
// Maps

// func main() {
	// statePopulations := make(map[string]int)
	// statePopulations = map[string]int{
	// 		"California": 39123456,
	// 		"Texas": 4652123,
	// 		"Florida": 4564646,
	// 		"New York": 789456,
	// 		"Pennsylvania": 3456456,
	// 		"Illinois": 7954123,
	// 		"Ohio": 45123,
	// }

// 	statePopulations["Georgia"] = 10310371

// 	// delete(statePopulations, "Georgia")

// 	_, ok := statePopulations["Ohio"]

// 	fmt.Println(ok)

// 	fmt.Println(len(statePopulations))

// 	fmt.Println(statePopulations)

// }

// Structs

type Doctor struct {
	number int
	actorName string
	companions []string
}

type Animal struct {
	Name string `required;max:"100"`
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly bool
}

func main() {
	// aDoctor := Doctor{
	// 	number: 3,
	// 	actorName: "Jon Pertwee",
	// 	companions: []string {
	// 		"Liz Shaw",
	// 		"Jo Grant",
	// 		"Jane Smith",
	// 	},
	// }

	//Creating a anonymous struct
	// aDoctor := struct{name string}{name: "John Pertwee"}
	// anotherDoctor := &aDoctor
	// anotherDoctor.name = "Tom Baker"
	// fmt.Println(aDoctor)
	// fmt.Println(anotherDoctor)

	// Embedding
	// Is good to altering a library. Something close to the decorator pattern.

	// b := Bird{}
	// b.Name = "Emu"
	// b.Origin = "Australia"
	// b.SpeedKPH = 48
	// b.CanFly = false
	// b := Bird {
	// 	Animal: Animal{Name: "Emu", Origin: "Australia"},
	// 	SpeedKPH: 30,
	// 	CanFly: false,
	// }
	// fmt.Println(b.Name)

	// Tags
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}