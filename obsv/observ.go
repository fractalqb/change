package obsv

import (
	"sort"

	"git.fractalqb.de/fractalqb/change"
)

type Veto struct {
	stopper Observer
	err     error
}

func (v *Veto) StoppedBy() Observer { return v.stopper }
func (v *Veto) Error() string       { return "veto: " + v.err.Error() }
func (v *Veto) Unwrap() error       { return v.err }

type Observer interface {
	// Check lets the observer inspect the change before it is executed. By
	// returning a non-nil error the observer can prohibit the change.
	// Check also can override the default change.Flags used for the Set
	// operation by returning Flags != 0.
	Check(tag interface{}, e Event) (change.Flags, error)
	// Update notifies the observer about a change event.
	Update(tag interface{}, e Event)
}

type UpdateFunc func(tag interface{}, e Event)

func (_ UpdateFunc) Check(tag interface{}, e Event) (change.Flags, error) {
	return 0, nil
}

func (f UpdateFunc) Update(tag interface{}, e Event) {
	f(tag, e)
}

type Observable interface {
	ObsDefaults() (tag interface{}, chg change.Flags)
	SetObsDefaults(tag interface{}, chg change.Flags)
	ObsRegister(prio int, tag interface{}, o Observer)
	ObsUnregister(o Observer)
	ObsLastVeto() *Veto
	ObsEach(do func(tag interface{}, o Observer))
}

type Event interface {
	Source() Observable
	Chg() change.Flags
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

func (s *stub) notify(e Event) {
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
