package chg16

type Chg uint16

func (cs Chg) Then(chg Chg) Chg { return cs | chg }
