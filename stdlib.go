package chgv

import (
	"time"
)

type Time time.Time

func (t Time) Get() time.Time { return time.Time(t) }

func (t *Time) Set(v time.Time, chg Chg) Chg {
	if time.Time(*t).Equal(v) {
		return 0
	}
	*t = Time(v)
	return chg
}

func (t Time) String() string { return time.Time(t).String() }

type Duration time.Duration

func (d Duration) Get() time.Duration { return time.Duration(d) }

func (d *Duration) Set(v time.Duration, chg Chg) Chg {
	if time.Duration(*d) == v {
		return 0
	}
	*d = Duration(v)
	return chg
}

func (d Duration) String() string { return time.Duration(d).String() }
