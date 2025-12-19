package rect


type Rect struct {
	Width int
	Height int
}

func (r *Rect) Formula () int {
	return r.Height * r.Width
}

func (r *Rect) SetWidth(w int) {
	r.Width = w
}

func (r *Rect) SetHeight(h int) {
	r.Height = h
}