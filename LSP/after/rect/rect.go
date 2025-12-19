package rect

type Rect struct {
	Width int
	Height int
}

func (r *Rect) Formula () int {
	return r.Height * r.Width
}

func SetRect (h,w int) *Rect {
	return &Rect{
		Width: w,
		Height: h,
	}
}