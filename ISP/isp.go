package main

import (
	"ISP/after"
	"ISP/after/creatures"
	"ISP/after/swimmer"
	"ISP/after/walker"
	_ "ISP/before"
	_ "ISP/before/creatures"
)




func main () {
	// -------BEFORE------------
	// human := &creatures.Human{Name: "John"}
	// creature := &before.Creature{}
	// creature.MoveWalk(human)
	// creature.MoveSwim(human)
	// creature.MoveSpread(human) // here is located of violation of isp principle because some class force to implement
	// method that the class not need and not use
	// -------BEFORE------------

	// -------AFTER------------
	human := &creatures.Human{Name: "Santi"}
	walker := &walker.Walker{Walk: human}
	swimmer1 := &swimmer.Swimmer{Swim: human}
	creature := &after.Creature{}
	
	fish := &creatures.Fish{Name: "Blob"}
	swimmer2 := &swimmer.Swimmer{Swim: fish}

	creature.Moving(walker)
	creature.Moving(swimmer1)
	creature.Moving(swimmer2)
	// -------AFTER------------


}