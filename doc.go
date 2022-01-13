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

// Package change has the common parts for the sub packages that help
// to track changes of values in an application. The sub packages
// implement different trade offs between resource consumption –
// memory and time – against features richness.  The change-detectable
// value types of the sub packages all implement their respective
// interface in this package which makes migration between the sub
// packages more easy.
//
// Compare the following example with the examples from the sub
// packages:
package change
