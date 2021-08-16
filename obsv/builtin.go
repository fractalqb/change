package obsv

import (
	"fmt"

	"git.fractalqb.de/fractalqb/change"
	"git.fractalqb.de/fractalqb/change/chgv"
)

type ValueChange interface {
	Event
	OldValue() interface{}
	NewValue() interface{}
}

type Int struct {
	v chgv.Int
	ObservableBase
}

func NewInt(init int, defaultTag interface{}, defaultChg change.Flags) *Int {
	res := &Int{v: chgv.Int(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

func (chgi Int) Get() int { return chgi.v.Get() }

func (chgi *Int) Set(v int, chg change.Flags) change.Flags {
	if chgi.stub == nil {
		return chgi.v.Set(v, chg)
	}
	chgi.stub.veto = nil
	e := IntChg{
		eventBase: eventBase{
			src: chgi,
			chg: chgi.stub.change(chg),
		},
		ov: chgi.v.Get(),
		nv: v,
	}
	if c, veto := chgi.stub.check(e); veto != nil {
		chgi.stub.veto = veto
		return 0
	} else if chg == 0 {
		chg = c
	}
	if chg = chgi.v.Set(v, chg); chg == 0 {
		return 0
	}
	e.chg = chg
	chgi.stub.notify(e)
	return chg
}

type IntChg struct {
	eventBase
	ov, nv int
}

func (ic IntChg) String() string {
	return fmt.Sprintf("Chg int %d -> %d flags=%x",
		ic.ov, ic.nv,
		ic.chg)
}

func (ic IntChg) OldValue() interface{} { return ic.ov }
func (ic IntChg) NewValue() interface{} { return ic.nv }

type String struct {
	v chgv.String
	ObservableBase
}

func NewString(init string, defaultTag interface{}, defaultChg change.Flags) *String {
	res := &String{v: chgv.String(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

func (str String) Get() string { return str.v.Get() }

func (str *String) Set(v string, chg change.Flags) change.Flags {
	if str.stub == nil {
		return str.v.Set(v, chg)
	}
	str.stub.veto = nil
	e := StringChg{
		eventBase: eventBase{
			src: str,
			chg: str.stub.change(chg),
		},
		ov: str.v.Get(),
		nv: v,
	}
	if c, veto := str.stub.check(e); veto != nil {
		str.stub.veto = veto
		return 0
	} else if chg == 0 {
		chg = c
	}
	if chg = str.v.Set(v, chg); chg == 0 {
		return 0
	}
	e.chg = chg
	str.stub.notify(e)
	return chg
}

type StringChg struct {
	eventBase
	ov, nv string
}

func (sc StringChg) String() string {
	return fmt.Sprintf("Chg string '%s' -> '%s' flags=%x",
		sc.ov, sc.nv,
		sc.chg)
}

func (sc StringChg) OldValue() interface{} { return sc.ov }
func (sc StringChg) NewValue() interface{} { return sc.nv }
