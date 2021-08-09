package chg32

type Chg uint32

func (cs Chg) Then(chg Chg) Chg { return cs | chg }
