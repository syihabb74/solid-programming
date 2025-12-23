package creatures

import "fmt"

type Human struct {
	Name string
}

func (f *Human) WalkMoving() {
	fmt.Println(f.Name, " Walking....")
}

func (f *Human) SwimMoving() {
	fmt.Println(f.Name, " Swimming....")
}
