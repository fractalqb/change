define(`onchgtype', `type $2 struct {
	v chgv.$2
	ObservableBase
}

func New$2(init $1, defaultTag interface{}, defaultChg change.Flags) *$2 {
	res := &$2{v: chgv.$2(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

func (cv $2) Get() $1 { return cv.v.Get() }

func (cv *$2) Set(v $1, chg change.Flags) change.Flags {
	if cv.stub == nil {
		return cv.v.Set(v, chg)
	}
	cv.stub.veto = nil
	e := $2Chg{
		eventBase: eventBase{
			src: cv,
			chg: cv.stub.change(chg),
		},
		ov: cv.v.Get(),
		nv: v,
	}
	if c, veto := cv.stub.check(e); veto != nil {
		cv.stub.veto = veto
		return 0
	} else if chg == 0 {
		chg = c
	}
	if chg = cv.v.Set(v, chg); chg == 0 {
		return 0
	}
	e.chg = chg
	cv.stub.notify(e)
	return chg
}

type $2Chg struct {
	eventBase
	ov, nv $1
}

func (ic $2Chg) String() string {
	return fmt.Sprintf("Chg $1 [%v] -> [%v] flags=%x",
		ic.ov, ic.nv,
		ic.chg)
}

func (ic $2Chg) OldValue() interface{} { return ic.ov }
func (ic $2Chg) NewValue() interface{} { return ic.nv }
')dnl
dnl
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

onchgtype(`bool', `Bool')
onchgtype(`uint8', `Uint8')
onchgtype(`uint16', `Uint16')
onchgtype(`uint32', `Uint32')
onchgtype(`uint64', `Uint64')
onchgtype(`int8', `Int8')
onchgtype(`int16', `Int16')
onchgtype(`int32', `Int32')
onchgtype(`int64', `Int64')
onchgtype(`float32', `Float32')
onchgtype(`float64', `Float64')
onchgtype(`complex64', `Complex64')
onchgtype(`complex128', `Complex128')
onchgtype(`byte', `Byte')
onchgtype(`rune', `Rune')
onchgtype(`uint', `Uint')
onchgtype(`int', `Int')
onchgtype(`uintptr', `UintPtr')
onchgtype(`string', `String')
