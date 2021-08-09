package chgv

import (
	"fmt"
)

func ExampleString() {
	s := String("Hello, change")
	chg := s.Set("Hello, change", 1)
	fmt.Println(chg)
	chg |= s.Set("Hello, change!", 1)
	fmt.Println(chg)
	chg |= s.Set("Change", 2)
	fmt.Println(chg, s.Get())
	// Output:
	// 0
	// 1
	// 3 Change
}
