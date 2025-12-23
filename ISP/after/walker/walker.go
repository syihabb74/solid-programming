package walker

type WalkMove interface {
	WalkMoving()
}

type Walker struct {
	Walk WalkMove
}

func (w *Walker) Move() {
	w.Walk.WalkMoving()
}