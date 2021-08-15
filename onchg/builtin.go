package onchg

import (
	"git.fractalqb.de/fractalqb/change/chgv"
)

type StringChgFunc func(src *String, ov, nv string, pre bool) chgv.Chg

type StringChg chgv.Chg

func (chg StringChg) Change(src *String, ov, nv string, pre bool) chgv.Chg {
	return chgv.Chg(chg)
}

type String struct {
	v chgv.String
	f StringChgFunc
}

func NewString(init string, onChg StringChgFunc) String {
	return String{chgv.String(init), onChg}
}

func (s *String) Reset(v string, onChg StringChgFunc) (string, StringChgFunc) {
	ov, of := s.v, s.f
	(*s) = String{chgv.String(v), onChg}
	return ov.Get(), of
}

func (s String) Get() string { return s.v.Get() }

func (s *String) Set(v string) (chg chgv.Chg) {
	if s.f == nil {
		return s.v.Set(v, 1)
	}
	if chg = s.f(s, s.v.Get(), v, true); chg == 0 {
		return 0
	}
	old := s.v.Get()
	if chg = s.v.Set(v, chg); chg != 0 {
		s.f(s, old, v, false)
	}
	return chg
}
