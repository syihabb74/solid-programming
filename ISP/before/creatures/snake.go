package creatures

import "fmt"
type Snake struct {
	Name string
}

func (s *Snake) Walk() {
	
}

func (s *Snake) Swim() {

}

func (s *Snake) Spread() {
	fmt.Println(s.Name,"Spreading....")
}