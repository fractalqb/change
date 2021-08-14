package chgv

import (
	"fmt"
	"testing"
)

func ExampleChg() {
	type user struct {
		id   Int
		name String
	}
	const (
		IDChg = (1 << iota)
		NameChg
	)
	u := user{id: 4711, name: "John Doe"}
	chg := u.id.Set(4712, IDChg)
	chg |= u.name.Set("John Doe", NameChg)
	fmt.Println(chg)
	chg |= u.name.Set("Jack Doe", NameChg)
	fmt.Println(chg)
	// Output:
	// 1
	// 3
}

func ExampleChg_Map() {
	fmt.Println(Chg(0xf0).Map(4))          // 0xf0 != 0 → 4
	fmt.Println(Chg(0xf0).Map(0x0f, 2, 1)) // 0xf0 & 0x0f == 0 → 0 and 0x0f != 0 → 1 ⇒ 0|1 = 1
	fmt.Println(Chg(0x18).Map(0x0f, 2, 1)) // 0x18 & 0x0f != 0 → 2 and 0x0f != 0 → 1 ⇒ 2|1 = 3
	fmt.Println(Chg(0x18).Map(0x0f, 2))    // 0x18 & 0x0f != 0 → 2
	// Output:
	// 4
	// 1
	// 3
	// 2
}

func TestChg_Map(t *testing.T) {
	t.Run("any to 1", func(t *testing.T) {
		mapped := Chg(0).Map(1)
		if mapped != 0 {
			t.Error("mapped 0 to some bit")
		}
		mapped = Chg(32).Map(1)
		if mapped != 1 {
			t.Errorf("mapped to unexpected value %x, not 1", mapped)
		}
	})
	t.Run("pattern to 1", func(t *testing.T) {
		mapped := Chg(0).Map(0xf0, 0x2)
		if mapped != 0 {
			t.Error("mapped 0 to some bit")
		}
		mapped = Chg(0x0f).Map(0xf0, 2)
		if mapped != 0 {
			t.Errorf("mapped disjoint to unexpected value %x, not 0", mapped)
		}
		mapped = Chg(0x18).Map(0xf0, 2)
		if mapped != 2 {
			t.Errorf("mapped to unexpected value %x, not 2", mapped)
		}
	})
	t.Run("mask and any", func(t *testing.T) {
		mapped := Chg(0).Map(0xf0, 2, 1)
		if mapped != 0 {
			t.Error("mapped 0 to some bit")
		}
		mapped = Chg(0x0f).Map(0xf0, 2, 1)
		if mapped != 1 {
			t.Errorf("mapped disjoint to unexpected value %x, not 1", mapped)
		}
		mapped = Chg(0x18).Map(0xf0, 2, 1)
		if mapped != 3 {
			t.Errorf("mapped to unexpected value %x, not 3", mapped)
		}
	})
}
