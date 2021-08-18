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

package onchg

import (
	"testing"

	"github.com/fractalqb/change"
)

var (
	_ change.Bool       = (*Bool)(nil)
	_ change.Uint8      = (*Uint8)(nil)
	_ change.Uint16     = (*Uint16)(nil)
	_ change.Uint32     = (*Uint32)(nil)
	_ change.Uint64     = (*Uint64)(nil)
	_ change.Int8       = (*Int8)(nil)
	_ change.Int16      = (*Int16)(nil)
	_ change.Int32      = (*Int32)(nil)
	_ change.Int64      = (*Int64)(nil)
	_ change.Float32    = (*Float32)(nil)
	_ change.Float64    = (*Float64)(nil)
	_ change.Complex64  = (*Complex64)(nil)
	_ change.Complex128 = (*Complex128)(nil)
	_ change.Byte       = (*Byte)(nil)
	_ change.Rune       = (*Rune)(nil)
	_ change.Uint       = (*Uint)(nil)
	_ change.Int        = (*Int)(nil)
	_ change.UintPtr    = (*UintPtr)(nil)
	_ change.String     = (*String)(nil)
)

func TestInt(t *testing.T) {
	t.Run("zero value", func(t *testing.T) {
		var iv Int
		if iv.Get() != 0 {
			t.Error("Int does not initialize to zero value")
		}
	})
	t.Run("without hook", func(t *testing.T) {
		iv := *NewInt(4, nil)
		if v := iv.Get(); v != 4 {
			t.Errorf("Int initialization failed. Got %d want 4", v)
		}
		if chg := iv.Set(4, 1); chg != 0 {
			t.Errorf("Unexpected change flags 0x%x, want 0", chg)
		}
		if v := iv.Get(); v != 4 {
			t.Errorf("Unexpected Int value: %d want 4", v)
		}
		if chg := iv.Set(3, 2); chg != 2 {
			t.Errorf("Invalid change flags 0x%x, want 2", chg)
		}
		if v := iv.Get(); v != 3 {
			t.Errorf("Unexpected Int value: %d want 3", v)
		}
	})
	t.Run("hook veto", func(t *testing.T) {
		iv := *NewInt(4, func(src *Int, ov, nv int, check bool) change.Flags {
			if !check {
				t.Error("Hook was called despite veto")
			}
			return 0 // Always veto
		})
		chg := iv.Set(3, 1)
		if chg != 0 {
			t.Errorf("Unexpected change flags 0x%x, want 0", chg)
		}
		if v := iv.Get(); v != 4 {
			t.Errorf("Unexpected Int value: %d want 4", v)
		}
	})
	t.Run("hook flags", func(t *testing.T) {
		iv := *NewInt(4, ChgFlag(1).Int)
		if v := iv.Get(); v != 4 {
			t.Errorf("Int initialization failed. Got %d want 4", v)
		}
		if chg := iv.Set(4, 0); chg != 0 {
			t.Errorf("Unexpected change flags 0x%x, want 0", chg)
		}
		if chg := iv.Set(4, 2); chg != 0 {
			t.Errorf("Unexpected change flags 0x%x, want 0", chg)
		}
		if v := iv.Get(); v != 4 {
			t.Errorf("Unexpected Int value: %d want 4", v)
		}
		if chg := iv.Set(3, 2); chg != 2 {
			t.Errorf("Invalid change flags 0x%x, want 2", chg)
		}
		if v := iv.Get(); v != 3 {
			t.Errorf("Unexpected Int value: %d want 3", v)
		}
		if chg := iv.Set(2, 0); chg != 1 {
			t.Errorf("Invalid change flags 0x%x, want 1", chg)
		}
		if v := iv.Get(); v != 2 {
			t.Errorf("Unexpected Int value: %d want 2", v)
		}
	})
}
