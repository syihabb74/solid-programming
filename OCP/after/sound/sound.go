package sound

type Sound interface {
	Sound()
}

type PrintSound struct {

}


func (s *PrintSound) Print(a Sound) {
	a.Sound()
}
