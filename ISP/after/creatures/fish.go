package creatures

import "fmt"

type Fish struct {
	Name string
}

func (f *Fish) SwimMoving() {
	fmt.Println(f.Name, " Swimming....")
}
