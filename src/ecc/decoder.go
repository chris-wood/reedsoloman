package main

import "./field"
import "os"
import "fmt"
import "io/ioutil"
import "log"

func main() {
	fmt.Println("Decoder...")
	args := os.Args[1:]

	fname = args[0]
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		// TODO: give data to the decoder
		fmt.Print(data)
	} else {
		log.Fatal("Error reading the file %s", string(err))
	}
}
