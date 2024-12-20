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
	"sort"
)

// Observable objects will notify registered [Observer] objects about state
// changes.  Observables use tags to let Observers easily distinguish the
// different subject they observe. In addition to change Events, Observables
// also use change Flags to descibe chages – much like Val and On values.
// Observers are notified in order of their registered priority, highest first.
// Observers with the same priority are notified in order of their registration.
type Observable interface {
	// ObsDefaults returns the default tag and change Flags of the observable.
	ObsDefaults() (tag any, chg Flags)

	// SetObsDefaults sets the default tag and change Flags of the observable.
	SetObsDefaults(tag any, chg Flags)

	// ObsRegister registres a new Observer with this Observable. If tag is not
	// nil, the Observer will be notified with this specific tag, not the
	// default tag.
	ObsRegister(prio int, tag any, o Observer)

	// ObsUnregister removes the Observer o from the observable if it is
	// registered. Otherwise ObsUnregister odes nothing.
	ObsUnregister(o Observer)

	// ObsLastVeto returns the Veto from the last Set call, if any.
	ObsLastVeto() *Veto

	// ObsEach calls do for all registered Observers in noitification order with
	// the same tag that would be used for notifications. ObsEach stops on first
	// non-nil error from do and returns that error.
	ObsEach(do func(tag any, o Observer) error) error
}

// Observers can be registered with Observables to be notified about state
// changes of the [Observable]. An Observer must be comparable to make
// Observable.ObsRegister() and Observable.ObsUnregister() work.
type Observer interface {
	// Check lets the observer inspect the hypothetical change Event e
	// before it is executed.  By returning a non-nil veto the
	// observer can block the change.
	Check(tag any, e Event) (veto error)

	// Update notifies the observer about a change Event e.
	Update(tag any, e Event)
}

// Use UpdateFunc to wrap an update function so that it can be used as
// an Observer. Note that the pointer to the UpdateFunc object will be
// used to implement equality. This is because Go functions are not
// comparable.
func UpdateFunc(f func(any, Event)) Observer { return &updFunc{f} }

// observers must be comparable
type updFunc struct {
	f func(tag any, e Event)
}

// Check will never block a change and always retuns 0 Flags, i.e. does not
// override default change Falgs.
func (*updFunc) Check(any, Event) error { return nil }

// Upate call the update function F.
func (uf *updFunc) Update(tag any, e Event) { uf.f(tag, e) }

// Event is the base interface for all state change events.
type Event interface {
	// Source returns the Observable that issued the Event.
	Source() Observable
	// Chg returns the change Flags for the state change.
	Chg() Flags
}

type eventBase struct {
	src Observable
	chg Flags
}

func (e eventBase) Source() Observable { return e.src }

func (e eventBase) Chg() Flags { return e.chg }

type ValueChange interface {
	Event

	// OldValue returns the value befor the change.
	OldValue() any

	// NewValue returns the current value, i.e. after the change.
	NewValue() any
}

type Changed[T comparable] struct {
	eventBase
	ov, nv T
}

// Implement ValueChange
func (ce Changed[T]) OldValue() any { return ce.ov }

// Implement ValueChange
func (ce Changed[T]) NewValue() any { return ce.nv }

// Veto keeps the information why an [Obs.Set] operation was stopped by which
// [Observer]. The Veto can be requested from each Observable until the next
// call to its Set method.
//
// Veto also implements an unwrappable Go error.
type Veto struct {
	stopper Observer
	err     error
}

// StoppedBy returns the Observer that stopped the Set operation with
// its Veto.
func (v *Veto) StoppedBy() Observer { return v.stopper }

func (v *Veto) Error() string { return "veto: " + v.err.Error() }
func (v *Veto) Unwrap() error { return v.err }

type ObservableBase struct {
	stub *stub
}

func (b ObservableBase) ObsDefaults() (tag any, chg Flags) {
	if b.stub == nil {
		return nil, 0
	}
	return b.stub.tag, b.stub.chg
}

func (b *ObservableBase) SetObsDefaults(tag any, chg Flags) {
	b.replaceStub(nil)
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

func (b *ObservableBase) ObsRegister(prio int, tag any, o Observer) {
	if o == nil {
		return
	}
	b.replaceStub(nil)
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

func (b *ObservableBase) ObsEach(do func(tag any, o Observer) error) error {
	if b.stub == nil {
		return nil
	}
	for _, oln := range b.stub.obsls {
		if err := do(oln.tag, oln.obs); err != nil {
			return err
		}
	}
	return nil
}

func (b *ObservableBase) replaceStub(s *stub) {
	if s != nil {
		if b.stub != nil {
			*s = *b.stub
		} else {
			*s = stub{}
		}
		b.stub = s
	}
}

// Obs implements Changeable as a full featured Observable.
type Obs[T comparable] struct {
	v Val[T]
	ObservableBase
}

var (
	_ Observable = (*Obs[int])(nil)
	_ Able[int]  = (*Obs[int])(nil)
)

func NewObs[T comparable](init T, defaultTag any, defaultChg Flags, os ...Observer) Obs[T] {
	res := Obs[T]{v: Val[T]{init}}
	res.SetObsDefaults(defaultTag, defaultChg)
	for _, o := range os {
		res.ObsRegister(0, nil, o)
	}
	return res
}

func (ov Obs[T]) Get() T { return ov.v.Get() }

func (ov *Obs[T]) Set(v T, chg Flags) Flags {
	if ov.stub == nil {
		return ov.v.Set(v, chg)
	}
	return ov.slowSet(v, chg)
}

func (ov *Obs[T]) slowSet(v T, chg Flags) Flags {
	ov.stub.veto = nil
	e := Changed[T]{
		eventBase: eventBase{
			src: ov,
			chg: ov.stub.change(chg),
		},
		ov: ov.v.Get(),
		nv: v,
	}
	if veto := ov.stub.check(e); veto != nil {
		ov.stub.veto = veto
		return 0
	}
	if ov.v.Set(v, 1) == 0 {
		return 0
	}
	ov.stub.update(e)
	return e.chg
}

type stub struct {
	tag   any
	chg   Flags
	obsls []obsLink // prio high → low
	veto  *Veto
}

func (s *stub) change(chg Flags) Flags {
	if chg == 0 {
		return s.chg
	}
	return chg
}

func (s *stub) check(e Event) *Veto {
	for _, oln := range s.obsls {
		var tag any
		if oln.tag != nil {
			tag = oln.tag
		} else {
			tag = s.tag
		}
		veto := oln.obs.Check(tag, e)
		if veto != nil {
			return &Veto{oln.obs, veto}
		}
	}
	return nil
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
	tag  any
	obs  Observer
}
