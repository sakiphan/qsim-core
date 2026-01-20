// Package vector provides unit-safe 3D vector operations for physics calculations.
//
// Vectors can represent physical quantities like position, velocity, force, etc.,
// with automatic dimensional tracking through the units package.
//
// Example usage:
//
//	import (
//	    "github.com/sakiphan/qsim-core/math/vector"
//	    "github.com/sakiphan/qsim-core/units"
//	)
//
//	// Position vectors
//	r1 := vector.NewPosition(units.Meter(1), units.Meter(2), units.Meter(3))
//	r2 := vector.NewPosition(units.Meter(4), units.Meter(5), units.Meter(6))
//
//	// Vector addition
//	r3, _ := r1.Add(r2) // (5, 7, 9) m
//
//	// Dot product: work = F · d
//	force := vector.NewForce(units.Newton(10), units.Newton(0), units.Newton(0))
//	displacement := vector.NewPosition(units.Meter(5), units.Meter(0), units.Meter(0))
//	work := force.Dot(displacement) // 50 J
//
//	// Cross product: torque = r × F
//	torque := r1.Cross(force) // Returns angular momentum dimension
//
// All vector operations preserve dimensional consistency and prevent
// incompatible operations at compile time.
package vector
