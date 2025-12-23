package before

type CreatureMove interface {
	Walk()
	Swim()
	Spread()
}

type Creature struct {

}

func (c *Creature) MoveWalk (cM CreatureMove) {

	cM.Walk()
}

func (c *Creature) MoveSwim (cM CreatureMove) {

	cM.Swim()
}

func (c *Creature) MoveSpread (cM CreatureMove) {

	cM.Spread()
}


