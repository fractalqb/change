package chgv

type Chg uint

func (cs Chg) Then(chg Chg) Chg { return cs | chg }
