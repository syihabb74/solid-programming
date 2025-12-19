package square

type Square struct {
	Size int
}


func (s *Square) Formula () int {
	return s.Size * s.Size
}

func (s *Square) SetHeight(h int) {
	s.Size = h
}

func (s *Square) SetWidth(w int) {
	s.Size = w
}

