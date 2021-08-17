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
package chgv

import (
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
