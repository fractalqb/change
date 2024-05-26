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

// All implements [Able] with minimal overhead without checking for changes.
type All[T comparable] struct{ v T }

var _ Able[int] = (*All[int])(nil)

func NewAll[T comparable](init T) All[T] { return All[T]{v: init} }

// Get returns the current value.
func (cv All[T]) Get() T {
	return cv.v
}

// Set cv's value to v and always return chg.
func (cv *All[T]) Set(v T, chg Flags) Flags {
	cv.v = v
	return chg
}
