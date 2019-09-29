package world

type IPosChangeListenser interface {
	onPosChanged(oldPos Point2D, newPos Point2D)
}
