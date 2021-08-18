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

// Package chgv implements values that track changes with a simple
// flag set of type 'change.Flags'. The flag(s) passed to any Set call
// are retuned if the underlying value changed. Otherwise 0 is
// returned. Note that if one passes flag=0 to Set the returned value
// will always be 0, which make it rather uninteresting. However this
// will not affect the actual value change.
//
// While these changeable values are rather bare bones they come
// without memory overhead – unlike most observable libraries.
//
// Set Methods
//
// TODO: describe principle of Set methods.
package chgv
