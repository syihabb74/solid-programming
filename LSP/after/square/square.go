package square


type Square struct {
	Size int
}

func (r *Square) Formula () int {
	return r.Size * r.Size
}

func SetSqr (s int) *Square {
	return &Square{
		Size: s,
	}
}