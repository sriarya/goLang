package main

import "fmt"

//	func printState() string {
//		return " Welcome to California, it's a lovely place"
//	}

type laptopSize float64

func state() {

	var device laptopSize = 34.5

	device.getSizeOfLaptop()
}

func (this laptopSize) getSizeOfLaptop() {
	fmt.Println(this)
}
