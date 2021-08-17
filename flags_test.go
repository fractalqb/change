// Copyright 2021 Marcus Perlick
// This file is part of Go module github.com/fractalqb/change
//
// Foobar is free software: you can redistribute it and/or modify it
// under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// Foobar is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
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
