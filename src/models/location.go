package volunteer/model

import "math"

type Location struct {
	name string
	x float
	y float
}

func (loc1 Location) DistanceTo(loc2 Location) float {
	deltaX := loc1.x - loc2.x
	deltaY := loc1.y - loc.y
	return math.Sqrt(deltaX * deltaX + deltaY * deltaY)
}

func (loc Location) Rename(newName string) {
  loc.name = newName
}

