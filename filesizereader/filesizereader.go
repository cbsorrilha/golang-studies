package main

import (
	"fmt"
	"flag"
	"os"
)

const (
	_ = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	filePath := flag.String("file-path", "", "")

	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
			// handle the error here
			panic("ERROR OPENING THE FILE")
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
			panic("ERROR GETING THE FILE SIZE")
	}

	fileSize := stat.Size()

	fmt.Println("File is:")
	fmt.Printf("Raw: %v\n", fileSize)
	fmt.Printf("%v KB\n", (float64(fileSize)/KB))
	fmt.Printf("%v MB\n", (float64(fileSize)/MB))
	fmt.Printf("%v GB\n", (float64(fileSize)/GB))
	fmt.Printf("%v TB\n", (float64(fileSize)/TB))
	fmt.Printf("%v PB\n", (float64(fileSize)/PB))
	fmt.Printf("%v EB\n", (float64(fileSize)/EB))
	fmt.Printf("%v ZB\n", (float64(fileSize)/ZB))
	fmt.Printf("%v YB\n", (float64(fileSize)/YB))

}