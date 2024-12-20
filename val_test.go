// Copyright 2021 Marcus Perlick
// This file is part of Go module github.com/fractalqb/change
//
// change is free software: you can redistribute it and/or modify it
// under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// change is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with change.  If not, see <http://www.gnu.org/licenses/>.

package change

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestVal_sizes(t *testing.T) {
	var (
		i  int16
		iv Val[int16]
		s  struct {
			a, b, c Val[int16]
		}
		iSz = unsafe.Sizeof(i)
	)
	if sz := unsafe.Sizeof(iv); sz != iSz {
		t.Errorf("Unexpected size %d of %T, want %d", sz, iv, iSz)
	}
	if sz := unsafe.Sizeof(s); sz != 3*iSz {
		t.Errorf("Unexpected size %d of %T, want %d", sz, s, 3*iSz)
	}
}

func ExampleVal() {
	user := struct {
		Name   Val[string]
		Logins Val[int]
	}{
		Name:   NewVal("John Doe"),
		Logins: NewVal(0),
	}
	chg := user.Name.Set("John Doe", 1) // 1 = (1<<0) = 0b01
	chg |= user.Logins.Set(1, 2)        // 2 = (1<<1) = 0b10

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

func TestValInt(t *testing.T) {
	t.Run("zero value", func(t *testing.T) {
		var iv Val[int]
		if iv.Get() != 0 {
			t.Error("Int does not initialize to zero value")
		}
	})
	t.Run("init", func(t *testing.T) {
		iv := NewVal(4)
		if v := iv.Get(); v != 4 {
			t.Errorf("Int initialization failed. Got %d want 4", v)
		}
	})
	t.Run("unchanged set", func(t *testing.T) {
		iv := NewVal(4)
		chg := iv.Set(4, 1)
		if v := iv.Get(); v != 4 {
			t.Errorf("Got unexpeted value %d, want 4", v)
		}
		if chg != 0 {
			t.Errorf("Unexpected change flags 0x%x, want 0", chg)
		}
	})
	t.Run("changing set", func(t *testing.T) {
		iv := NewVal(4)
		chg := iv.Set(0, 1)
		if v := iv.Get(); v != 0 {
			t.Errorf("Got unexpeted value %d, want 0", v)
		}
		if chg != 1 {
			t.Errorf("Unexpected change flags 0x%x, want 1", chg)
		}
	})
}
