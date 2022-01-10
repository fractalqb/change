// Copyright 2021 Marcus Perlick
// This file is part of Go module github.com/fractalqb/change
//
// change is free software: you can redistribute it and/or modify it
// under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// change is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with change.  If not, see <http://www.gnu.org/licenses/>.

package chgv

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
	t.Run("init", func(t *testing.T) {
		iv := Int(4)
		if v := iv.Get(); v != 4 {
			t.Errorf("Int initialization failed. Got %d want 4", v)
		}
	})
	t.Run("unchanged set", func(t *testing.T) {
		iv := Int(4)
		chg := iv.Set(4, 1)
		if v := iv.Get(); v != 4 {
			t.Errorf("Got unexpeted value %d, want 4", v)
		}
		if chg != 0 {
			t.Errorf("Unexpected change flags 0x%x, want 0", chg)
		}
	})
	t.Run("changing set", func(t *testing.T) {
		iv := Int(4)
		chg := iv.Set(0, 1)
		if v := iv.Get(); v != 0 {
			t.Errorf("Got unexpeted value %d, want 0", v)
		}
		if chg != 1 {
			t.Errorf("Unexpected change flags 0x%x, want 1", chg)
		}
	})
}
