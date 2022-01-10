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
	"reflect"
	"testing"
)

func TestObservableBase_ObsRegister_prio(t *testing.T) {
	onUpdate := func(tag interface{}, e Event) {
		t.Log("Update", tag, e)
	}
	var obl ObservableBase
	obl.ObsRegister(-1, 1, &UpdateFunc{onUpdate})
	obl.ObsRegister(0, 2, &UpdateFunc{onUpdate})
	obl.ObsRegister(1, 3, &UpdateFunc{onUpdate})
	obl.ObsRegister(-1, 4, &UpdateFunc{onUpdate})
	obl.ObsRegister(0, 5, &UpdateFunc{onUpdate})
	obl.ObsRegister(1, 6, &UpdateFunc{onUpdate})
	var tags []int
	obl.ObsEach(func(tag interface{}, o Observer) {
		tags = append(tags, tag.(int))
	})
	if !reflect.DeepEqual(tags, []int{3, 6, 2, 5, 1, 4}) {
		t.Errorf("Wrong observer order: %v", tags)
	}
}
