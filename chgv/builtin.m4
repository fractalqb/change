define(`onchgtype', `// $2 tracks changes of $1 values with caller-provided Flags.
// See also chgv package documentation.
type $2 $1

// Get returns the current $1 value.
func (cv $2) Get() $1 { return $1(cv) }

// Set the value of cv to v. See also section Set Methods.
func (cv *$2) Set(v $1, chg change.Flags) change.Flags {
	if $1(*cv) == v {
		return 0
	}
	*cv = $2(v)
	return chg
}
')dnl
dnl
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

import "github.com/fractalqb/change"

onchgtype(`bool', `Bool')
onchgtype(`uint8', `Uint8')
onchgtype(`uint16', `Uint16')
onchgtype(`uint32', `Uint32')
onchgtype(`uint64', `Uint64')
onchgtype(`int8', `Int8')
onchgtype(`int16', `Int16')
onchgtype(`int32', `Int32')
onchgtype(`int64', `Int64')
onchgtype(`float32', `Float32')
onchgtype(`float64', `Float64')
onchgtype(`complex64', `Complex64')
onchgtype(`complex128', `Complex128')
onchgtype(`byte', `Byte')
onchgtype(`rune', `Rune')
onchgtype(`uint', `Uint')
onchgtype(`int', `Int')
onchgtype(`uintptr', `UintPtr')
onchgtype(`string', `String')
