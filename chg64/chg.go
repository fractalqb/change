package chg64

type Chg uint64

func (cs Chg) Then(chg Chg) Chg { return cs | chg }
