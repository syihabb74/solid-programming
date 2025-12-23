package creatures

import "fmt"

type Snake struct {
	Name string
}

func (s *Snake) SpreadMoving() {
	fmt.Println(s.Name, " Spreading....")
}
