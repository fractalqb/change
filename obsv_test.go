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

import (
	"fmt"
	"testing"
)

var _ Changeable[int] = (*Obs[int])(nil)

func ExampleObs_string() {
	str := NewObs("", "example string", 4711)
	str.ObsRegister(0, nil, &UpdateFunc{func(tag interface{}, e Event) {
		chg := e.(Changed[string])
		fmt.Printf("Tag: %+v\n", tag)
		fmt.Printf("Event: '%s'→'%s': %d\n",
			chg.OldValue(),
			chg.NewValue(),
			e.Chg())
	}})
	fmt.Println(str.Set("Observe this change!", 0), str.ObsLastVeto())
	// Output:
	// Tag: example string
	// Event: ''→'Observe this change!': 4711
	// 4711 <nil>
}

func ExampleObs_ObsRegister() {
	obs := NewObs("", nil, 0)
	updFunc := func(tag interface{}, _ Event) { fmt.Println(tag) }
	obs.ObsRegister(0, "A", &UpdateFunc{updFunc})
	obs.ObsRegister(1, "B", &UpdateFunc{updFunc})
	obs.ObsRegister(1, "C", &UpdateFunc{updFunc})
	obs.ObsRegister(2, "D", &UpdateFunc{updFunc})
	obs.Set("foo", 0)
	// Output:
	// D
	// B
	// C
	// A
}

func TestObs_NewObs_withObserver(t *testing.T) {
	count := 0
	obs := NewObs("", nil, 1, &UpdateFunc{func(any, Event) {
		count++
	}})
	obs.Set("foo", 1)
	if count != 1 {
		t.Errorf("change event observed %d times", count)
	}
}
