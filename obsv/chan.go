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
	"sync/atomic"

	"github.com/fractalqb/change"
)

type TaggedEvent struct {
	Event
	Tag interface{}
}

// An ObservingChan wraps a Go channel so that the channel's input
// becomes an Observer, i.e. the ObservingChan can register with
// Observables and then will send tagged update Events to the channel.
type ObservingChan struct {
	// If CheckFunc is not nil it will be called in the
	// ObservingChan's Check method.
	CheckFunc func(tag interface{}, e Event) (change.Flags, error)
	c         chan<- TaggedEvent
	drop      bool
	dropCount uint64
}

func NewObservingChan(c chan<- TaggedEvent, dropping bool) *ObservingChan {
	return &ObservingChan{c: c, drop: dropping}
}

func (oc ObservingChan) Close() {
	close(oc.c)
}

func (oc ObservingChan) DropCount() uint64 {
	return atomic.LoadUint64(&oc.dropCount)
}

func (oc ObservingChan) Check(tag interface{}, e Event) (change.Flags, error) {
	if oc.CheckFunc == nil {
		return 0, nil
	}
	return oc.CheckFunc(tag, e)
}

func (oc ObservingChan) Update(tag interface{}, e Event) {
	te := TaggedEvent{Event: e, Tag: tag}
	if oc.drop {
		select {
		case oc.c <- te:
		default:
			atomic.AddUint64(&oc.dropCount, 1)
		}
	} else {
		oc.c <- te
	}
}

// An ObservableChan wraps a Go channel of TaggedEvents so that the
// channel's output becomes Observable.
type ObservableChan struct {
	c    <-chan TaggedEvent
	stop chan struct{}
	ObservableBase
}

func NewObservableChan(c <-chan TaggedEvent) *ObservableChan {
	return &ObservableChan{c: c}
}

func (oc *ObservableChan) Start() {
	oc.stop = make(chan struct{})
LOOP:
	for {
		select {
		case e := <-oc.c:
			oc.forward(e)
		case <-oc.stop:
			break LOOP
		}
	}
	close(oc.stop)
}

func (oc *ObservableChan) Stop() {
	if oc.stop != nil {
		oc.stop <- struct{}{}
		<-oc.stop
	}
}

func (oc *ObservableChan) forward(te TaggedEvent) {
	if oc.stub == nil {
		return
	}
	oc.stub.veto = nil
	_, veto := oc.stub.check(te) // TODO what to do with chg?
	if veto != nil {
		oc.stub.veto = veto
		return
	}
	oc.stub.update(te)
}
