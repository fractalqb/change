package obsv

import (
	"fmt"
	"time"

	"github.com/fractalqb/change"
	"github.com/fractalqb/change/chgv"
)

type Time struct {
	v chgv.Time
	ObservableBase
}

func NewTime(init time.Time, defaultTag interface{}, defaultChg change.Flags) *Time {
	res := &Time{v: chgv.Time(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

func (cv Time) Get() time.Time { return cv.v.Get() }

func (cv *Time) Set(v time.Time, chg change.Flags) change.Flags {
	if cv.stub == nil {
		return cv.v.Set(v, chg)
	}
	cv.stub.veto = nil
	e := TimeChg{
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
	cv.stub.update(e)
	return chg
}

type TimeChg struct {
	eventBase
	ov, nv time.Time
}

func (ic TimeChg) String() string {
	return fmt.Sprintf("Chg time.Time [%v] -> [%v] flags=%x",
		ic.ov, ic.nv,
		ic.chg)
}

func (ic TimeChg) OldValue() interface{} { return ic.ov }
func (ic TimeChg) NewValue() interface{} { return ic.nv }

type Duration struct {
	v chgv.Duration
	ObservableBase
}

func NewDuration(init time.Duration, defaultTag interface{}, defaultChg change.Flags) *Duration {
	res := &Duration{v: chgv.Duration(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

func (cv Duration) Get() time.Duration { return cv.v.Get() }

func (cv *Duration) Set(v time.Duration, chg change.Flags) change.Flags {
	if cv.stub == nil {
		return cv.v.Set(v, chg)
	}
	cv.stub.veto = nil
	e := DurationChg{
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
	cv.stub.update(e)
	return chg
}

type DurationChg struct {
	eventBase
	ov, nv time.Duration
}

func (ic DurationChg) String() string {
	return fmt.Sprintf("Chg time.Duration [%v] -> [%v] flags=%x",
		ic.ov, ic.nv,
		ic.chg)
}

func (ic DurationChg) OldValue() interface{} { return ic.ov }
func (ic DurationChg) NewValue() interface{} { return ic.nv }
