package printarea

import (
	"LSP/before/rect"
	"fmt"
)

type PrintArea struct {

}


func (pA *PrintArea) Print (rf rect.RectForm) {
	rf.SetHeight(5)
	rf.SetWidth(12)
	fmt.Println(rf.Formula())
}

func CreatePrintArea () *PrintArea {
	return &PrintArea{}
}
