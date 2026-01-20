package vector

import (
	"fmt"
	"math"

	"github.com/sakiphan/qsim-core/units"
)

// Vector3 represents a 3D vector with physical units.
// Each component is a unit-safe Value, ensuring dimensional consistency.
type Vector3 struct {
	X, Y, Z units.Value
}

// New creates a new Vector3 with the specified components.
// All components must have the same dimension.
//
// Example:
//
//	v := vector.New(
//	    units.Meter(1.0).Value,
//	    units.Meter(2.0).Value,
//	    units.Meter(3.0).Value,
//	)
func New(x, y, z units.Value) (Vector3, error) {
	// Check dimensional consistency
	if x.Dim() != y.Dim() || x.Dim() != z.Dim() {
		return Vector3{}, fmt.Errorf("vector components must have same dimension: x=%s, y=%s, z=%s",
			x.Dim(), y.Dim(), z.Dim())
	}
	return Vector3{X: x, Y: y, Z: z}, nil
}

// NewPosition creates a position vector with Length components.
//
// Example:
//
//	r := vector.NewPosition(
//	    units.Meter(1.0),
//	    units.Meter(2.0),
//	    units.Meter(3.0),
//	)
func NewPosition(x, y, z units.Length) Vector3 {
	return Vector3{X: x.Value, Y: y.Value, Z: z.Value}
}

// NewVelocity creates a velocity vector with Velocity components.
func NewVelocity(vx, vy, vz units.Velocity) Vector3 {
	return Vector3{X: vx.Value, Y: vy.Value, Z: vz.Value}
}

// NewAcceleration creates an acceleration vector with Acceleration components.
func NewAcceleration(ax, ay, az units.Acceleration) Vector3 {
	return Vector3{X: ax.Value, Y: ay.Value, Z: az.Value}
}

// NewForce creates a force vector with Force components.
func NewForce(fx, fy, fz units.Force) Vector3 {
	return Vector3{X: fx.Value, Y: fy.Value, Z: fz.Value}
}

// Zero creates a zero vector with the specified dimension.
func Zero(dim units.Dimension) Vector3 {
	return Vector3{
		X: units.NewValue(0, dim),
		Y: units.NewValue(0, dim),
		Z: units.NewValue(0, dim),
	}
}

// UnitX returns a unit vector in the X direction with the specified dimension.
func UnitX(dim units.Dimension) Vector3 {
	return Vector3{
		X: units.NewValue(1, dim),
		Y: units.NewValue(0, dim),
		Z: units.NewValue(0, dim),
	}
}

// UnitY returns a unit vector in the Y direction with the specified dimension.
func UnitY(dim units.Dimension) Vector3 {
	return Vector3{
		X: units.NewValue(0, dim),
		Y: units.NewValue(1, dim),
		Z: units.NewValue(0, dim),
	}
}

// UnitZ returns a unit vector in the Z direction with the specified dimension.
func UnitZ(dim units.Dimension) Vector3 {
	return Vector3{
		X: units.NewValue(0, dim),
		Y: units.NewValue(1, dim),
		Z: units.NewValue(0, dim),
	}
}

// String returns a human-readable representation of the vector.
func (v Vector3) String() string {
	return fmt.Sprintf("(%v, %v, %v)", v.X, v.Y, v.Z)
}

// Dim returns the dimension of the vector components.
func (v Vector3) Dim() units.Dimension {
	return v.X.Dim()
}

// Add returns the sum of two vectors. Vectors must have the same dimension.
//
// Example:
//
//	r1 := vector.NewPosition(units.Meter(1), units.Meter(2), units.Meter(3))
//	r2 := vector.NewPosition(units.Meter(4), units.Meter(5), units.Meter(6))
//	r3, _ := r1.Add(r2) // (5, 7, 9) m
func (v Vector3) Add(other Vector3) (Vector3, error) {
	x, err := v.X.Add(other.X)
	if err != nil {
		return Vector3{}, err
	}
	y, err := v.Y.Add(other.Y)
	if err != nil {
		return Vector3{}, err
	}
	z, err := v.Z.Add(other.Z)
	if err != nil {
		return Vector3{}, err
	}
	return Vector3{X: x, Y: y, Z: z}, nil
}

// Subtract returns the difference of two vectors. Vectors must have the same dimension.
//
// Example:
//
//	displacement := r2.Subtract(r1) // Displacement from r1 to r2
func (v Vector3) Subtract(other Vector3) (Vector3, error) {
	x, err := v.X.Subtract(other.X)
	if err != nil {
		return Vector3{}, err
	}
	y, err := v.Y.Subtract(other.Y)
	if err != nil {
		return Vector3{}, err
	}
	z, err := v.Z.Subtract(other.Z)
	if err != nil {
		return Vector3{}, err
	}
	return Vector3{X: x, Y: y, Z: z}, nil
}

// Scale multiplies the vector by a dimensionless scalar.
//
// Example:
//
//	v2 := v1.Scale(2.0) // Double the vector
func (v Vector3) Scale(scalar float64) Vector3 {
	return Vector3{
		X: v.X.Scale(scalar),
		Y: v.Y.Scale(scalar),
		Z: v.Z.Scale(scalar),
	}
}

// Negate returns the negation of the vector (-v).
func (v Vector3) Negate() Vector3 {
	return Vector3{
		X: v.X.Negate(),
		Y: v.Y.Negate(),
		Z: v.Z.Negate(),
	}
}

// Dot returns the dot product of two vectors.
// Result has dimension equal to the product of component dimensions.
//
// Example:
//
//	// Work: W = F · d
//	force := vector.NewForce(units.Newton(10), ...)
//	displacement := vector.NewPosition(units.Meter(5), ...)
//	work := force.Dot(displacement) // Result has dimension [L²MT⁻²] (energy)
func (v Vector3) Dot(other Vector3) units.Value {
	xx := v.X.Multiply(other.X)
	yy := v.Y.Multiply(other.Y)
	zz := v.Z.Multiply(other.Z)

	result, _ := xx.Add(yy)
	result, _ = result.Add(zz)
	return result
}

// Cross returns the cross product of two vectors.
// Result has dimension equal to the product of component dimensions.
//
// The cross product is defined as:
//
//	v × w = (v_y*w_z - v_z*w_y, v_z*w_x - v_x*w_z, v_x*w_y - v_y*w_x)
//
// Example:
//
//	// Torque: τ = r × F
//	position := vector.NewPosition(units.Meter(1), units.Meter(0), units.Meter(0))
//	force := vector.NewForce(units.Newton(0), units.Newton(10), units.Newton(0))
//	torque := position.Cross(force) // Result has dimension [L²MT⁻²]
func (v Vector3) Cross(other Vector3) Vector3 {
	// v × w = (v_y*w_z - v_z*w_y, v_z*w_x - v_x*w_z, v_x*w_y - v_y*w_x)
	xVal, _ := v.Y.Multiply(other.Z).Subtract(v.Z.Multiply(other.Y))
	yVal, _ := v.Z.Multiply(other.X).Subtract(v.X.Multiply(other.Z))
	zVal, _ := v.X.Multiply(other.Y).Subtract(v.Y.Multiply(other.X))

	// These subtractions should never fail since dimensions match from multiply
	return Vector3{X: xVal, Y: yVal, Z: zVal}
}

// MagnitudeSquared returns the squared magnitude of the vector (v · v).
// This avoids the square root and preserves exact arithmetic.
//
// Example:
//
//	v := vector.NewVelocity(units.MeterPerSecond(3), units.MeterPerSecond(4), units.MeterPerSecond(0))
//	vSquared := v.MagnitudeSquared() // 25 m²/s²
func (v Vector3) MagnitudeSquared() units.Value {
	return v.Dot(v)
}

// Magnitude returns the magnitude (length) of the vector: |v| = √(v · v).
// Returns an error if the dimension cannot be square-rooted (odd exponents).
//
// Example:
//
//	v := vector.NewVelocity(units.MeterPerSecond(3), units.MeterPerSecond(4), units.MeterPerSecond(0))
//	speed, _ := v.Magnitude() // 5 m/s
func (v Vector3) Magnitude() (units.Value, error) {
	magSquared := v.MagnitudeSquared()
	return magSquared.Sqrt()
}

// Normalize returns a unit vector in the same direction.
// Only works for dimensionless vectors or when you want a direction vector.
//
// Example:
//
//	direction := v.Normalize() // Returns dimensionless direction vector
func (v Vector3) Normalize() (Vector3, error) {
	mag, err := v.Magnitude()
	if err != nil {
		return Vector3{}, err
	}

	if mag.Val() == 0 {
		return Vector3{}, fmt.Errorf("cannot normalize zero vector")
	}

	// Divide each component by magnitude to get dimensionless direction
	return Vector3{
		X: v.X.Divide(mag),
		Y: v.Y.Divide(mag),
		Z: v.Z.Divide(mag),
	}, nil
}

// ProjectOnto projects this vector onto another vector.
// Returns the component of v in the direction of other.
//
// Projection formula: proj_w(v) = (v · w / |w|²) * w
//
// Example:
//
//	// Project velocity onto a direction
//	velocityAlongAxis, _ := velocity.ProjectOnto(axis)
func (v Vector3) ProjectOnto(other Vector3) (Vector3, error) {
	dotProduct := v.Dot(other)
	otherMagSquared := other.MagnitudeSquared()

	if otherMagSquared.Val() == 0 {
		return Vector3{}, fmt.Errorf("cannot project onto zero vector")
	}

	// (v · w) / |w|²
	scalar := dotProduct.Divide(otherMagSquared)

	// Multiply each component of w by scalar
	return Vector3{
		X: other.X.Multiply(scalar),
		Y: other.Y.Multiply(scalar),
		Z: other.Z.Multiply(scalar),
	}, nil
}

// AngleBetween returns the angle (in radians) between two vectors.
// Result is dimensionless.
//
// Formula: cos(θ) = (v · w) / (|v| |w|)
//
// Example:
//
//	angle := v1.AngleBetween(v2) // Returns angle in radians
func (v Vector3) AngleBetween(other Vector3) (float64, error) {
	dotProduct := v.Dot(other)
	magV, err := v.Magnitude()
	if err != nil {
		return 0, err
	}
	magOther, err := other.Magnitude()
	if err != nil {
		return 0, err
	}

	magProduct := magV.Multiply(magOther)
	if magProduct.Val() == 0 {
		return 0, fmt.Errorf("cannot compute angle with zero vector")
	}

	cosTheta := dotProduct.Divide(magProduct).Val()

	// Clamp to [-1, 1] to handle numerical errors
	if cosTheta > 1.0 {
		cosTheta = 1.0
	}
	if cosTheta < -1.0 {
		cosTheta = -1.0
	}

	return math.Acos(cosTheta), nil
}

// IsZero returns true if all components are zero.
func (v Vector3) IsZero() bool {
	return v.X.Val() == 0 && v.Y.Val() == 0 && v.Z.Val() == 0
}

// IsParallel returns true if vectors are parallel (including antiparallel).
// Uses cross product: v × w = 0 if parallel.
func (v Vector3) IsParallel(other Vector3, tolerance float64) bool {
	cross := v.Cross(other)
	magSquared := cross.MagnitudeSquared()
	return math.Abs(magSquared.Val()) < tolerance
}

// IsPerpendicular returns true if vectors are perpendicular.
// Uses dot product: v · w = 0 if perpendicular.
func (v Vector3) IsPerpendicular(other Vector3, tolerance float64) bool {
	dot := v.Dot(other)
	return math.Abs(dot.Val()) < tolerance
}

// Components returns the X, Y, Z components as a slice.
func (v Vector3) Components() []units.Value {
	return []units.Value{v.X, v.Y, v.Z}
}

// ToArray returns the vector components as a float64 array (in SI base units).
func (v Vector3) ToArray() [3]float64 {
	return [3]float64{v.X.Val(), v.Y.Val(), v.Z.Val()}
}
