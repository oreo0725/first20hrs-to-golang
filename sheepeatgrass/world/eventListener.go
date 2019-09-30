package world

type IWorldChangeListenser interface {
	OnPosChanged(oldPos Point2D, newPos Point2D)
	OnCreatureEaten(creature ICreature, food IFood)
	OnLifeDead(creature ICreature)
	OnNewLifeBorn(creature ICreature)
}
