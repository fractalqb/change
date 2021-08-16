package onchg

import (
	"fmt"

	"git.fractalqb.de/fractalqb/change"
)

func Example() {
	user := struct {
		Name   String
		Logins Int
	}{
		Name: *NewString("John Doe", ChgFlag(1).String),
		Logins: *NewInt(0, func(i *Int, o, n int, pre bool) change.Flags {
			if !pre {
				fmt.Printf("changed logins from %d to %d\n", o, n)
			}
			return 2
		}),
	}
	chg := user.Name.Set("John Doe", 0)
	chg |= user.Logins.Set(1, 0)
	fmt.Printf("Changes: 0b%b\n", chg)
	if chg&1 == 0 {
		fmt.Println("Name did not change")
	} else {
		fmt.Println("Name changed")
	}
	// Output:
	// changed logins from 0 to 1
	// Changes: 0b10
	// Name did not change
}
