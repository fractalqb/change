define(`onchgtype', `// $2 is the common interface of change-detecable $1 values.
type $2 interface {
	Get() $1
	Set(v $1, chg Flags) Flags
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

package change

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
