package change

import (
	"fmt"
	"testing"
)

func ExampleFlags_Map() {
	fmt.Println(Flags(0xf0).Map(4))          // 0xf0 != 0 → 4
	fmt.Println(Flags(0xf0).Map(0x0f, 2, 1)) // 0xf0 & 0x0f == 0 → 0 and 0x0f != 0 → 1 ⇒ 0|1 = 1
	fmt.Println(Flags(0x18).Map(0x0f, 2, 1)) // 0x18 & 0x0f != 0 → 2 and 0x0f != 0 → 1 ⇒ 2|1 = 3
	fmt.Println(Flags(0x18).Map(0x0f, 2))    // 0x18 & 0x0f != 0 → 2
	// Output:
	// 4
	// 1
	// 3
	// 2
}

func TestFlags_Map(t *testing.T) {
	t.Run("any to 1", func(t *testing.T) {
		mapped := Flags(0).Map(1)
		if mapped != 0 {
			t.Error("mapped 0 to some bit")
		}
		mapped = Flags(32).Map(1)
		if mapped != 1 {
			t.Errorf("mapped to unexpected value %x, not 1", mapped)
		}
	})
	t.Run("pattern to 1", func(t *testing.T) {
		mapped := Flags(0).Map(0xf0, 0x2)
		if mapped != 0 {
			t.Error("mapped 0 to some bit")
		}
		mapped = Flags(0x0f).Map(0xf0, 2)
		if mapped != 0 {
			t.Errorf("mapped disjoint to unexpected value %x, not 0", mapped)
		}
		mapped = Flags(0x18).Map(0xf0, 2)
		if mapped != 2 {
			t.Errorf("mapped to unexpected value %x, not 2", mapped)
		}
	})
	t.Run("mask and any", func(t *testing.T) {
		mapped := Flags(0).Map(0xf0, 2, 1)
		if mapped != 0 {
			t.Error("mapped 0 to some bit")
		}
		mapped = Flags(0x0f).Map(0xf0, 2, 1)
		if mapped != 1 {
			t.Errorf("mapped disjoint to unexpected value %x, not 1", mapped)
		}
		mapped = Flags(0x18).Map(0xf0, 2, 1)
		if mapped != 3 {
			t.Errorf("mapped to unexpected value %x, not 3", mapped)
		}
	})
}
