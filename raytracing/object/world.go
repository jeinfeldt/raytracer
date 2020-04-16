package object

type (
	// World represents a list of hittable objects
	World []Hittable
)

// NewWorld is a factory method to create a new world
func NewWorld() World {
	return World{}
}

// Add adds a hittable to the world
func (w *World) Add(h Hittable) {
	*w = append(*w, h)
}

// Clear removes hittables from the world
func (w *World) Clear() {
	*w = NewWorld()
}

// Hit indicates whether or not an object inside the list has been hit
func (w *World) Hit(r Ray, tmin float64, tmax float64, record *HitRecord) bool {
	localRecord := &HitRecord{}
	hitAnything := false
	closest := tmax
	// loop over objects in list and check which was hit
	for _, ele := range *w {
		// guard if element was not hit by ray, continue
		if !ele.Hit(r, tmin, closest, localRecord) {
			continue
		}
		hitAnything = true
		closest = localRecord.Time
		// change pointer address
		*record = *localRecord
	}

	return hitAnything
}
