package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
)

func main() {
	// Get the file name of self
	_, path, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Could not get the caller path")
	}
	fmt.Println(path)

	// Read the self
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// Print the self
	fmt.Println(string(data))
}
