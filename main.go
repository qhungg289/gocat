package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	file, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}

	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	b := make([]byte, stat.Size())
	_, err = file.Read(b)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Name: %v, Size: %v bytes, Modified: %v\n", stat.Name(), stat.Size(), stat.ModTime())
	fmt.Println(string(b))

	if err = file.Close(); err != nil {
		log.Fatal(err)
	}
}
