package onchg

import (
	"fmt"

	"git.fractalqb.de/fractalqb/change/chgv"
)

func ExampleString() {
	s := NewString("", func(_ *String, ov, nv string, pre bool) chgv.Chg {
		if pre {
			if nv == "" {
				return 0
			}
		} else {
			fmt.Printf("Changed from '%s' to '%s'\n", ov, nv)
		}
		return 1
	})
	fmt.Println("Chg:", s.Set("Embrace change"))
	fmt.Println("Chg:", s.Set(""))
	// Output:
	// Changed from '' to 'Embrace change'
	// Chg: 1
	// Chg: 0
}
