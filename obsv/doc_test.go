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

package obsv

import (
	"fmt"

	"github.com/fractalqb/change"
)

func Example() {
	user := struct {
		ObservableBase
		Name   String `obsv:"called,flag=0b100"`
		Logins Int    `obsv:",tag=given-tag,flag=1"`
	}{
		Name:   *NewString("John Doe", nil, 0),
		Logins: *NewInt(0, nil, 0),
	}
	ObserveFields(&user, 0, false, change.AllFlags)
	user.ObsRegister(0, nil, &UpdateFunc{func(tag interface{}, e Event) {
		fmt.Println("Tag:", tag, "Event:", e)
	}})

	chg := user.Name.Set("John Doe", 0)
	chg |= user.Logins.Set(1, 2)

	fmt.Printf("Changes: 0b%b\n", chg)
	if chg&1 == 0 {
		fmt.Println("Name did not change")
	} else {
		fmt.Println("Name changed")
	}
	// Output:
	// Tag: <nil> Event: Field 'Logins': tag=given-tag event=Chg int [0] -> [1] flags=2
	// Changes: 0b10
	// Name did not change
}
