package creatures

import "fmt"


type Human struct {
	Name string
}

func (h *Human) Walk() {
	fmt.Println(h.Name, " Walking....")
}

func (h *Human) Swim() {
	fmt.Println(h.Name, " Swimming....")
}

func (h *Human) Spread() {
	
}