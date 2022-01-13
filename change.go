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

// Flags is used by all Set methods to return to the caller if a Set
// operation did change a value. Generally Flags==0 means nothing
// changed.  The flags from different Set calls can efficiently be
// combined with the '|' binary or operator. However there is room for
// no more than 64 different flags.
//
// The specific flags for a call to Set are generally provided by the
// caller. It depends on the sub package what it means to pass 0 flags
// to a Set method. E.g. there may be defaults that kick in if one
// passes 0 flags to the actual call.
type Flags uint64

// AllFlags is a Flags set with all flags set – just to have said
// that.
const AllFlags = ^Flags(0)

// Any tests if any of bits is set in c.
func (c Flags) Any(bits Flags) bool { return c&bits != 0 }

// All tests if all bits are set in c.
func (c Flags) All(bits Flags) bool { return c&bits == bits }

// Map remaps bits of c to a new combination of bits.
//
// If the number of bits elements is odd then the last element is
// added to res if c is not zero. The even number of leading elements
// is treated as a list of pairs. The first element of a pair is
// checked to have any bits common with c.  If so, the second element
// of the pair is added to res.
func (c Flags) Map(bits ...Flags) (res Flags) {
	l := len(bits) & ^1
	for i := 0; i < l; i++ {
		b := bits[i]
		i++
		if c.Any(b) {
			res |= bits[i]
		}
	}
	if l != len(bits) && c != 0 {
		res |= bits[l]
	}
	return res
}

type Changeable[T comparable] interface {
	Get() T
	Set(v T, chg Flags) Flags
	//Bypass() *T
}
