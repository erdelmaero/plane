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

// SetComponents sets the Planes components
func (p *Plane3) SetComponents(x, y, z, w float64) {
	p.normal = vector.NewVector3(x, y, z)
	p.constant = w
}

// SetFromNormalAndCoplanarPoint sets the Plane from normal vector and onepoint containted by the plane
func (p *Plane3) SetFromNormalAndCoplanarPoint(normal *vector.Vector3, point *vector.Vector3) {
	p.normal = normal
	p.constant = -point.Dot(p.normal)
}

// SetFromCoplanarPoints sets the Plane from three Coplanar Points
func (p *Plane3) SetFromCoplanarPoints(a, b, c *vector.Vector3) {
	v1 := c.Sub(b)
	v2 := a.Sub(b)
	normal := v1.NewCross(v2).Normalize()
	p.SetFromNormalAndCoplanarPoint(normal, a)
}
