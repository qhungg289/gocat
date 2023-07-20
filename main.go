package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("Please provide correct path to file. Example: gocat sample.txt")
	}

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
	fmt.Println("---")
	fmt.Print(string(b))

	if err = file.Close(); err != nil {
		log.Fatal(err)
	}
}
