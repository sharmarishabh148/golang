package main

import "fmt"

type user struct {
	name     string
	username string
}

func main() {
	// Construct a map set to its zero value,
	// that can store user values based on a key of type string.
	// Trying to use this map will result in a runtime error (panic).
	//!var users map[string]user
	// Construct a map initialized using make,
	// that can store user values based on a key of type string.
	users := make(map[string]user)
	// Construct a map initialized using empty literal construction,
	// that can store user values based on a key of type string.
	// or users := map[string]user{}

	users["Roy"] = user{"Rob", "Roy"}
	users["Ford"] = user{"Henry", "Ford"}
	users["Mouse"] = user{"Mickey", "Mouse"}
	users["Jackson"] = user{"Michael", "Jackson"}
	for key, value := range users {
		fmt.Println(key, value)
	}

	user1, exists1 := users["Bill"]
	user2, exists2 := users["Ford"]
	fmt.Println("Bill:", exists1, user1)
	fmt.Println("Ford:", exists2, user2)

	delete(users, "Roy")

	//type slice []user
	//Users2 := make(map[slice]user)

}
