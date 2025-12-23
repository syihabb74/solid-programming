package swimmer

type SwimMove interface {
	SwimMoving()
}

type Swimmer struct {
	Swim SwimMove
}

func (s *Swimmer) Move() {
	s.Swim.SwimMoving()
}