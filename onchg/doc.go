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

// Package onchg implements values with a hook function that handles
// value changes. First the hook is called before a change is made
// with check=true. With check==true the hook decides if it blocks the
// set operation. To block the set operation, the hook returns
// 0. Otherwise the hook provides the change flags for the Set
// method. The flags from the hook are overridden if non-zero flags
// are passed to the value's Set method. If the hook is nil all Set
// methods of this package behave exactly like their chgv-package
// counterparts.
//
// Set Methods
//
// TODO: describe principle of Set methods.
//
// Hook Functions
//
// TODO: describe principle of hook functions.
package onchg
