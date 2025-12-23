package after

type CreatureMove interface {
	Move()
}

type Creature struct {

}


func (c *Creature) Moving(cM CreatureMove) {
	cM.Move()
}

