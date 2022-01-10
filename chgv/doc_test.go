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

package chgv

import "fmt"

func Example() {
	user := struct {
		Name   String
		Logins Int
	}{
		Name:   "John Doe",
		Logins: 0,
	}

	chg := user.Name.Set("John Doe", 1) // = (1<<0) = 0b01
	chg |= user.Logins.Set(1, 2)        // = (1<<1) = 0b10

	fmt.Printf("Changes: 0b%b\n", chg)
	if chg&1 == 0 {
		fmt.Println("Name did not change")
	} else {
		fmt.Println("Name changed")
	}
	// Output:
	// Changes: 0b10
	// Name did not change
}
