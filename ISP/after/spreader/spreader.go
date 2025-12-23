package spreader

type SpreadMove interface {
	SpreadMoving()
}

type Spreader struct {
	Spread SpreadMove
}


func (s *Spreader) Move () {
	s.Spread.SpreadMoving()
}

