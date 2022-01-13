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

// Val implements values that track changes with a simple flag-set of
// type Flags. The flag(s) passed to any Set method call are retuned
// if the underlying value changed. Otherwise the passed value is not
// assigned and 0 is returned. Note that if one passes flag=0 to Set
// the returned value will always be 0, which makes it rather
// uninteresting. However this will not affect the actual value
// change.
//
// While these changeable values are rather bare bones they come
// without memory overhead – unlike most observable libraries.
type Val[T comparable] struct{ v T }

func NewVal[T comparable](init T) Val[T] { return Val[T]{v: init} }

// Get returns the current value.
func (cv Val[T]) Get() T {
	return cv.v
}

// Set checks whether v is equal to the current value and returns 0 if
// it matches. Otherwise v is set as cv's value and chg is returned.
func (cv *Val[T]) Set(v T, chg Flags) Flags {
	if cv.v == v {
		return 0
	}
	cv.v = v
	return chg
}
