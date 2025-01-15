package main

import "fmt"

func main() {
	dictionary := map[string]string{
		"Go":     "A programing language created by Google.",
		"Gopher": "A software engenier who build with Go.",
		"Golang": "Another name for Go",
	}
	fmt.Println(dictionary)
	fmt.Println(dictionary["Gopher"])
	dictionary["Map"] = "An unordered data structure with key-value pairs"
	fmt.Println(dictionary["Map"])
	dictionary["Gopher"] = "The fuzzy mascot for Go."
	fmt.Println(dictionary["Gopher"] + "\n")

	for word, def := range dictionary {
		fmt.Printf("the definition of %s, is %s\n", word, def)
	}
	fmt.Printf(">-----------------------------------------------------------------------<\n")
	delete(dictionary, "Map")
	for word, def := range dictionary {
		fmt.Printf("the definition of %s, is %s\n", word, def)
	}
}
