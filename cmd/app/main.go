package main

import "log"

func main() {
	err := Run()
	if err != nil {
		log.Fatal(err)
	}
}
