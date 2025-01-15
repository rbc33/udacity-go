package main

import "fmt"

func main() {
	athletes := []string{"Stephen", "Klay", "Harrison", "Draymond", "Andrew"}

	for i, name := range athletes {
		fmt.Printf("i = %d, name = %s\n", i, name)
	}

	for _, name := range athletes {
		fmt.Printf("name = %s\n", name)
	}
}
