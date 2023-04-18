// Created by: Jakub Malhotra
// Created on: April 2023
//
// This program calculates the area of a triangle

package main

import "fmt"

func main() {
	// This function calculates the area of a triangle
	var base int
	var height int
	var area int

	fmt.Println("This program finds the area of a triangle.")
	fmt.Println()
	fmt.Print("Enter the base (cm): ")
	fmt.Scanln(&base)
	fmt.Print("Enter the height (cm): ")
	fmt.Scanln(&height)

	area = (base * height) /2

	fmt.Println()
	fmt.Println("The area is:", area, "cm².")
	fmt.Println("\nDone.")
}
