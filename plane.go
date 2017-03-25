package plane

import (
	"github.com/stojg/vector"
)

// Plane3 Struct
type Plane3 struct {
	normal   *vector.Vector3
	constant float64
}

// NewPlane3 creates a new Plane
func NewPlane3(normal *vector.Vector3, constant float64) *Plane3 {
	p := &Plane3{}
	p.normal = normal
	p.constant = constant
	return p
}

// Set Plane from Normal and Constant
func (p *Plane3) Set(normal *vector.Vector3, constant float64) {
	p.normal = normal
	p.constant = constant
}

// DistanceToPoint calculates the Distance from the Plane to a given Point
func (p *Plane3) DistanceToPoint(point *vector.Vector3) float64 {
	return p.normal.Dot(point) + p.constant
}
