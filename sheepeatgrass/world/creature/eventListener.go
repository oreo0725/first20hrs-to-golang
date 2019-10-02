package creature

import (
	"zentest.io/sheepeatgrass/world/geo"
)

type IWorldChangeListenser interface {
	OnPosChanged(oldPos geo.Point2D, newPos geo.Point2D)
	OnCreatureEaten(creature ICreature, food IFood)
	OnLifeDead(creature ICreature)
	OnNewLifeBorn(creature ICreature)
}
