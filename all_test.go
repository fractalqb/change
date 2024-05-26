package change

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestAll_sizes(t *testing.T) {
	var (
		i  int16
		iv All[int16]
		s  struct {
			a, b, c All[int16]
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

func ExampleAll() {
	user := struct {
		Name   All[string]
		Logins All[int]
	}{
		Name:   NewAll("John Doe"),
		Logins: NewAll(0),
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
	// Changes: 0b11
	// Name changed
}

func TestAllInt(t *testing.T) {
	t.Run("zero value", func(t *testing.T) {
		var iv All[int]
		if iv.Get() != 0 {
			t.Error("Int does not initialize to zero value")
		}
	})
	t.Run("init", func(t *testing.T) {
		iv := NewAll(4)
		if v := iv.Get(); v != 4 {
			t.Errorf("Int initialization failed. Got %d want 4", v)
		}
	})
	t.Run("unchanged set", func(t *testing.T) {
		iv := NewAll(4)
		chg := iv.Set(4, 1)
		if v := iv.Get(); v != 4 {
			t.Errorf("Got unexpeted value %d, want 4", v)
		}
		if chg == 0 {
			t.Error("No change")
		}
	})
	t.Run("changing set", func(t *testing.T) {
		iv := NewAll(4)
		chg := iv.Set(0, 1)
		if v := iv.Get(); v != 0 {
			t.Errorf("Got unexpeted value %d, want 0", v)
		}
		if chg != 1 {
			t.Errorf("Unexpected change flags 0x%x, want 1", chg)
		}
	})
}
