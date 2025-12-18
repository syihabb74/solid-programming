package main

import (
	"ocp/after/animal/bird"
	"ocp/after/animal/snake"
	"ocp/after/sound"
	_ "ocp/before/animal/bird"
	_ "ocp/before/animal/snake"
	_ "ocp/before/sound"
)


func main () {
	// b := &bird.Bird{}
	// s := &sound.PrintSound{}
	// sn := &snake.Snake{}
	// // s.Print(b) // we have to modify Sound struct method every new class impelement sound
	// s.Print(sn)



	b := &bird.Bird{};
	s := &snake.Snake{};
	sound := &sound.PrintSound{}

	sound.Print(b)
	sound.Print(s)

}