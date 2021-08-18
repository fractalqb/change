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

package obsv

import (
	"fmt"
	"time"
)

var (
	_ Observer   = ObservingChan{}
	_ Observable = (*ObservableChan)(nil)
)

func ExampleObservingChan() {
	s := NewString("", "the-tag", 1)
	c := make(chan TaggedEvent)
	o := NewObservingChan(c, false)
	s.ObsRegister(0, nil, o)
	go func() {
		for te := range c {
			fmt.Println(te)
		}
		fmt.Println("Exit channel receiver")
	}()
	s.Set("John Doe", 0)
	time.Sleep(100 * time.Millisecond)
	o.Close()
	time.Sleep(100 * time.Millisecond)
	// Output:
	// {Chg string [] -> [John Doe] flags=1 the-tag}
	// Exit channel receiver
}

func ExampleObservableChan() {
	c := make(chan TaggedEvent)
	o := NewObservableChan(c)
	o.ObsRegister(0, "the-tag", &UpdateFunc{func(tag interface{}, e Event) {
		fmt.Println(tag, e)
	}})
	go o.Start()
	c <- TaggedEvent{Event: eventBase{}, Tag: "/dev/null"}
	o.Stop()
	// Output:
	// the-tag {{<nil> 0} /dev/null}
}
