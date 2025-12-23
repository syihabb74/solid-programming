package creatures

import "fmt"
type Fish struct {
	Name string
}

func (f *Fish) Walk() {

}

func (f *Fish) Swim() {
	fmt.Println(f.Name, " Swimming....")
}

func (f *Fish) Spread() {

}