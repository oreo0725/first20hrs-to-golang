package world

type IWorldChangeListenser interface {
	onPosChanged(oldPos Point2D, newPos Point2D)
	onCreatureEaten(creature ICreature, food IFood)
}
