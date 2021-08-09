package chg8

import (
	"fmt"
)

func ExampleChanges() {
	type user struct {
		id   Int
		name String
	}
	const (
		IDChg = (1 << iota)
		NameChg
	)
	u := user{id: 4711, name: "John Doe"}
	chg := u.id.Set(4712, IDChg).
		Then(u.name.Set("John Doe", NameChg))
	fmt.Println(chg)
	chg = chg.Then(u.name.Set("Jack Doe", NameChg))
	fmt.Println(chg)
	// Output:
	// 1
	// 3
}
