package world

type IWorldChangeListenser interface {
	OnPosChanged(oldPos Point2D, newPos Point2D)
	OnCreatureEaten(creature ICreature, food IFood)
	OnLifeDead(life Life)
	OnNewLifeBorn(creature ICreature)
}
