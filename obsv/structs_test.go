package obsv

import (
	"fmt"

	"git.fractalqb.de/fractalqb/change"
)

func ExampleObserveFields() {
	var person struct {
		ObservableBase
		Called String `obsv:"called,flag=0b100"`
		Given  String `obsv:",tag=given-tag,flag=1"`
	}
	fmt.Println("ObserveFields:", ObserveFields(&person, 0, true, change.AllFlags))
	person.ObsRegister(0, nil, UpdateFunc(func(tag interface{}, e Event) {
		fmt.Println(tag, e)
	}))
	chg := person.Called.Set("John", 2)
	chg |= person.Given.Set("Doe", 0)
	fmt.Println("flags", chg)
	// Output:
	// ObserveFields: <nil>
	// <nil> Field 'called': tag={Name:Called PkgPath: Type:obsv.String Tag:obsv:"called,flag=0b100" Offset:8 Index:[1] Anonymous:false} event=Chg string '' -> 'John' flags=2
	// <nil> Field 'Given': tag={Name:Given PkgPath: Type:obsv.String Tag:obsv:",tag=given-tag,flag=1" Offset:32 Index:[2] Anonymous:false} event=Chg string '' -> 'Doe' flags=1
	// flags 3
}
