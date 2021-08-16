package chgv

import "fmt"

func Example() {
	user := struct {
		Name   String
		Logins Int
	}{
		Name:   "John Doe",
		Logins: 0,
	}
	chg := user.Name.Set("John Doe", 1) // = (1<<0) = 0b01
	chg |= user.Logins.Set(1, 2)        // = (1<<1) = 0b10
	fmt.Printf("Changes: 0b%b\n", chg)
	if chg&1 == 0 {
		fmt.Println("Name did not change")
	} else {
		fmt.Println("Name changed")
	}
	// Output:
	// Changes: 0b10
	// Name did not change
}
