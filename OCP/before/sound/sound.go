package sound

import (
	_ "ocp/before/animal/bird"
	"ocp/before/animal/snake"
)



type PrintSound struct {

}


func (s *PrintSound) Print(snake *snake.Snake) {
	snake.Sound()
}

