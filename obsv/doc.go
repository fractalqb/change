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

// Package obsv implements changeable values that follow the GoF
// Observer pattern. However, it is not a strict implementation but
// has been enriched with additional features:
//
//  - The Set method of changeable values is compatible with the Set
//    methods of the other sub packages, i.e. it handles
//    'change.Flags'. This also makes it easy for the caller to
//    detect changes, just in case you need it.
//
//  - Observers are added with a priority. This determines the order
//    in which operations of the observers are called, high priorities
//    first; within a priority in the order of registration.
//
//  - Observers can check a Set operation and decide to veto it
//    before the actual change is made.
//
//  - Observers Check methods can provide Falgs. This is to free the
//    caller of Set methods from dealing with Flags when not needed.
//
//  - Value-specific and observer-specific tags can be defined so
//    that they will be passed to the observer when a change occurs.
//    Observer-specific tags override vaule-specific tags.
package obsv
