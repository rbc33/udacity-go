package main

import "fmt"

func main() {
	var product string = "T-shirt"
	var cost int = 20

	fmt.Printf("The product is: %v", product)
	fmt.Printf("\nThe price is: %v", cost)
	fmt.Printf("\nThe types are: product:%T, price%T", product, cost)

	cost = 15

	fmt.Printf("\nThe product is: %v", product)
	fmt.Printf("\nThe price is: %v", cost)
	fmt.Printf("\nThe types are: product:%T, price%T", product, cost)

}
