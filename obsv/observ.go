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
	"sort"

	"github.com/fractalqb/change"
)

// Veto keeps the information why a Set operation was stopped by which
// Obersver.  The Veto can be requested from each Observable until the
// next call to its Set method.
type Veto struct {
	stopper Observer
	err     error
}

// StoppedBy returns the Observer that stopped the Set operation with
// its Veto.
func (v *Veto) StoppedBy() Observer { return v.stopper }

func (v *Veto) Error() string { return "veto: " + v.err.Error() }
func (v *Veto) Unwrap() error { return v.err }

// Observer must be comparable to make Observable::ObsRegister and
// Observable::ObsUnregister work.
type Observer interface {
	// Check lets the observer inspect the change before it is executed. By
	// returning a non-nil error the observer can veto the change.
	// Check also can override the default change.Flags used for the Set
	// operation by returning chg != 0.
	Check(tag interface{}, e Event) (chg change.Flags, veto error)
	// Update notifies the observer about a change event.
	Update(tag interface{}, e Event)
}

// Use UpdateFunc to wrap an update function so that it can be used as
// an Observer. Note that the pointer to the UpdateFunc object will be
// used to implement equality. This is because Go functions are not
// comparable.
type UpdateFunc struct {
	F func(tag interface{}, e Event)
}

func (_ UpdateFunc) Check(tag interface{}, e Event) (change.Flags, error) {
	return 0, nil
}

func (uf *UpdateFunc) Update(tag interface{}, e Event) {
	uf.F(tag, e)
}

type Observable interface {
	ObsDefaults() (tag interface{}, chg change.Flags)
	SetObsDefaults(tag interface{}, chg change.Flags)
	ObsRegister(prio int, tag interface{}, o Observer)
	ObsUnregister(o Observer)
	// ObsLastVeto returns the Veto from the last Set call, if any.
	ObsLastVeto() *Veto
	ObsEach(do func(tag interface{}, o Observer))
}

type Event interface {
	Source() Observable
	Chg() change.Flags
}

type ValueChange interface {
	Event
	// OldValue returns the value befor the change.
	OldValue() interface{}
	// NewValue returns the current value, i.e. after the change.
	NewValue() interface{}
}

type stub struct {
	tag   interface{}
	chg   change.Flags
	obsls []obsLink // prio high → low
	veto  *Veto
}

func (s stub) change(chg change.Flags) change.Flags {
	if chg == 0 {
		return s.chg
	}
	return chg
}

func (s *stub) check(e Event) (change.Flags, *Veto) {
	for _, oln := range s.obsls {
		var tag interface{}
		if oln.tag != nil {
			tag = oln.tag
		} else {
			tag = s.tag
		}
		chg, veto := oln.obs.Check(tag, e)
		if veto != nil {
			if chg == 0 {
				return e.Chg(), &Veto{oln.obs, veto}
			}
			return chg, &Veto{oln.obs, veto}
		}
	}
	return e.Chg(), nil
}

func (s *stub) update(e Event) {
	for _, oln := range s.obsls {
		if oln.tag != nil {
			oln.obs.Update(oln.tag, e)
		} else {
			oln.obs.Update(s.tag, e)
		}
	}
}

type obsLink struct {
	prio int
	tag  interface{}
	obs  Observer
}

type eventBase struct {
	src Observable
	chg change.Flags
}

func (e eventBase) Source() Observable { return e.src }

func (e eventBase) Chg() change.Flags { return e.chg }

type ObservableBase struct {
	stub *stub
}

func (b ObservableBase) ObsDefaults() (tag interface{}, chg change.Flags) {
	if b.stub == nil {
		return nil, 0
	}
	return b.stub.tag, b.stub.chg
}

func (b *ObservableBase) SetObsDefaults(tag interface{}, chg change.Flags) {
	if b.stub == nil {
		b.stub = &stub{
			chg: chg,
			tag: tag,
		}
	} else {
		b.stub.chg = chg
		b.stub.tag = tag
	}
}

func (b *ObservableBase) ObsRegister(prio int, tag interface{}, o Observer) {
	if o == nil {
		return
	}
	if b.stub == nil {
		b.stub = new(stub)
	}
	b.ObsUnregister(o)
	obsls := b.stub.obsls
	ins := sort.Search(len(obsls), func(i int) bool {
		return obsls[i].prio < prio
	})
	oln := obsLink{prio, tag, o}
	obsls = append(obsls, oln)
	copy(obsls[ins+1:], obsls[ins:])
	obsls[ins] = oln
	b.stub.obsls = obsls
}

func (b *ObservableBase) ObsUnregister(o Observer) {
	if b.stub == nil {
		return
	}
	for i, oln := range b.stub.obsls {
		if oln.obs == o {
			obsls := b.stub.obsls
			copy(obsls[i:], obsls[i+1:])
			obsls = obsls[:len(obsls)-1]
			b.stub.obsls = obsls
			return
		}
	}
}

func (b *ObservableBase) ObsLastVeto() *Veto {
	if b.stub == nil {
		return nil
	}
	return b.stub.veto
}

func (b *ObservableBase) ObsEach(do func(tag interface{}, o Observer)) {
	if b.stub == nil {
		return
	}
	for _, oln := range b.stub.obsls {
		do(oln.tag, oln.obs)
	}
}
