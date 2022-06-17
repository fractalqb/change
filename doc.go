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

// Package change helps to keep track of changed values in Go
// programs. It has three different concepts to do so that vary in
// memory overhead, time overhead and features: The simple Val, the Do
// with a single callback hook and the full-featured observable Obs.
// All implement a common interface Changeable and are implemented as
// generics.
//
// Compare the following plain-Go example with the examples from Val,
// On and Obs:
package change
