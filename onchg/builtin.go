package onchg

import (
	"git.fractalqb.de/fractalqb/change"
	"git.fractalqb.de/fractalqb/change/chgv"
)

type ChgFlag change.Flags

func (chg ChgFlag) Int(src *Int, ov, nv int, pre bool) change.Flags {
	return change.Flags(chg)
}

func (chg ChgFlag) String(src *String, ov, nv string, pre bool) change.Flags {
	return change.Flags(chg)
}

type Int struct {
	v chgv.Int
	f IntChgFunc
}

func NewInt(init int, onChg IntChgFunc) *Int {
	return &Int{chgv.Int(init), onChg}
}

type IntChgFunc func(src *Int, ov, nv int, pre bool) change.Flags

func (chgi *Int) Reset(v int, onChg IntChgFunc) (int, IntChgFunc) {
	ov, of := chgi.v, chgi.f
	(*chgi) = Int{chgv.Int(v), onChg}
	return ov.Get(), of
}

func (chgi Int) Get() int { return chgi.v.Get() }

func (chgi *Int) Set(v int, chg change.Flags) change.Flags {
	if chgi.f == nil {
		return chgi.v.Set(v, chg)
	}
	if c := chgi.f(chgi, chgi.v.Get(), v, true); c == 0 {
		return 0
	} else if chg == 0 {
		chg = c
	}
	old := chgi.v.Get()
	if chg = chgi.v.Set(v, chg); chg != 0 {
		chgi.f(chgi, old, v, false)
	}
	return chg
}

type String struct {
	v chgv.String
	f StringChgFunc
}

func NewString(init string, onChg StringChgFunc) *String {
	return &String{chgv.String(init), onChg}
}

type StringChgFunc func(src *String, ov, nv string, pre bool) change.Flags

func (s *String) Reset(v string, onChg StringChgFunc) (string, StringChgFunc) {
	ov, of := s.v, s.f
	(*s) = String{chgv.String(v), onChg}
	return ov.Get(), of
}

func (s String) Get() string { return s.v.Get() }

func (s *String) Set(v string, chg change.Flags) change.Flags {
	if s.f == nil {
		return s.v.Set(v, chg)
	}
	if c := s.f(s, s.v.Get(), v, true); c == 0 {
		return 0
	} else if chg == 0 {
		chg = c
	}
	old := s.v.Get()
	if chg = s.v.Set(v, chg); chg != 0 {
		s.f(s, old, v, false)
	}
	return chg
}
