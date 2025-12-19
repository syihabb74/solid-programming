package printarea

import (
	"fmt"
)

type Form interface {
	Formula() int
}

type PrintArea struct {

}


func (pA *PrintArea) Print (f Form) {
	fmt.Println(f.Formula())
}

func CreatePrintArea () *PrintArea {
	return &PrintArea{}
}
