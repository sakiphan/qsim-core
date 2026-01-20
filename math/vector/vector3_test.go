package vector

import (
	"math"
	"testing"

	"github.com/sakiphan/qsim-core/units"
)

func almostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) < tolerance
}

// -----------------------------------------------------------------------------
// Constructor Tests
// -----------------------------------------------------------------------------

func TestNew(t *testing.T) {
	// Valid: all same dimension
	v, err := New(
		units.Meter(1).Value,
		units.Meter(2).Value,
		units.Meter(3).Value,
	)
	if err != nil {
		t.Errorf("New() with same dimensions failed: %v", err)
	}
	if v.X.Val() != 1 || v.Y.Val() != 2 || v.Z.Val() != 3 {
		t.Errorf("New() values incorrect: got (%v, %v, %v)", v.X.Val(), v.Y.Val(), v.Z.Val())
	}

	// Invalid: different dimensions
	_, err = New(
		units.Meter(1).Value,
		units.Second(2).Value,
		units.Meter(3).Value,
	)
	if err == nil {
		t.Error("New() should fail with different dimensions")
	}
}

func TestNewPosition(t *testing.T) {
	v := NewPosition(
		units.Meter(1),
		units.Meter(2),
		units.Meter(3),
	)

	expected := units.Dimension{L: 1}
	if v.Dim() != expected {
		t.Errorf("NewPosition dimension = %v, want %v", v.Dim(), expected)
	}
}

func TestNewVelocity(t *testing.T) {
	v := NewVelocity(
		units.MeterPerSecond(10),
		units.MeterPerSecond(20),
		units.MeterPerSecond(30),
	)

	expected := units.Dimension{L: 1, T: -1}
	if v.Dim() != expected {
		t.Errorf("NewVelocity dimension = %v, want %v", v.Dim(), expected)
	}
}

func TestZero(t *testing.T) {
	dim := units.Dimension{L: 1}
	v := Zero(dim)

	if !v.IsZero() {
		t.Error("Zero() should create zero vector")
	}
	if v.Dim() != dim {
		t.Errorf("Zero() dimension = %v, want %v", v.Dim(), dim)
	}
}

// -----------------------------------------------------------------------------
// Basic Operations Tests
// -----------------------------------------------------------------------------

func TestAdd(t *testing.T) {
	v1 := NewPosition(units.Meter(1), units.Meter(2), units.Meter(3))
	v2 := NewPosition(units.Meter(4), units.Meter(5), units.Meter(6))

	v3, err := v1.Add(v2)
	if err != nil {
		t.Errorf("Add() failed: %v", err)
	}

	if v3.X.Val() != 5 || v3.Y.Val() != 7 || v3.Z.Val() != 9 {
		t.Errorf("Add() = (%v, %v, %v), want (5, 7, 9)", v3.X.Val(), v3.Y.Val(), v3.Z.Val())
	}

	// Test incompatible dimensions
	v4 := NewVelocity(units.MeterPerSecond(1), units.MeterPerSecond(2), units.MeterPerSecond(3))
	_, err = v1.Add(v4)
	if err == nil {
		t.Error("Add() should fail with incompatible dimensions")
	}
}

func TestSubtract(t *testing.T) {
	v1 := NewPosition(units.Meter(5), units.Meter(7), units.Meter(9))
	v2 := NewPosition(units.Meter(1), units.Meter(2), units.Meter(3))

	v3, err := v1.Subtract(v2)
	if err != nil {
		t.Errorf("Subtract() failed: %v", err)
	}

	if v3.X.Val() != 4 || v3.Y.Val() != 5 || v3.Z.Val() != 6 {
		t.Errorf("Subtract() = (%v, %v, %v), want (4, 5, 6)", v3.X.Val(), v3.Y.Val(), v3.Z.Val())
	}
}

func TestScale(t *testing.T) {
	v1 := NewPosition(units.Meter(1), units.Meter(2), units.Meter(3))
	v2 := v1.Scale(2.0)

	if v2.X.Val() != 2 || v2.Y.Val() != 4 || v2.Z.Val() != 6 {
		t.Errorf("Scale(2) = (%v, %v, %v), want (2, 4, 6)", v2.X.Val(), v2.Y.Val(), v2.Z.Val())
	}
}

func TestNegate(t *testing.T) {
	v1 := NewPosition(units.Meter(1), units.Meter(-2), units.Meter(3))
	v2 := v1.Negate()

	if v2.X.Val() != -1 || v2.Y.Val() != 2 || v2.Z.Val() != -3 {
		t.Errorf("Negate() = (%v, %v, %v), want (-1, 2, -3)", v2.X.Val(), v2.Y.Val(), v2.Z.Val())
	}
}

// -----------------------------------------------------------------------------
// Dot Product Tests
// -----------------------------------------------------------------------------

func TestDot(t *testing.T) {
	tests := []struct {
		name     string
		v1       Vector3
		v2       Vector3
		expected float64
	}{
		{
			name:     "orthogonal vectors",
			v1:       NewPosition(units.Meter(1), units.Meter(0), units.Meter(0)),
			v2:       NewPosition(units.Meter(0), units.Meter(1), units.Meter(0)),
			expected: 0.0,
		},
		{
			name:     "parallel vectors",
			v1:       NewPosition(units.Meter(2), units.Meter(0), units.Meter(0)),
			v2:       NewPosition(units.Meter(3), units.Meter(0), units.Meter(0)),
			expected: 6.0, // 2*3 = 6 m²
		},
		{
			name:     "general case",
			v1:       NewPosition(units.Meter(1), units.Meter(2), units.Meter(3)),
			v2:       NewPosition(units.Meter(4), units.Meter(5), units.Meter(6)),
			expected: 32.0, // 1*4 + 2*5 + 3*6 = 32 m²
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dot := tt.v1.Dot(tt.v2)
			if !almostEqual(dot.Val(), tt.expected, 1e-10) {
				t.Errorf("Dot() = %v, want %v", dot.Val(), tt.expected)
			}
		})
	}
}

func TestDotWork(t *testing.T) {
	// Work = Force · displacement
	force := NewForce(
		units.Newton(10),
		units.Newton(0),
		units.Newton(0),
	)
	displacement := NewPosition(
		units.Meter(5),
		units.Meter(0),
		units.Meter(0),
	)

	work := force.Dot(displacement)

	// Expected dimension: [L²MT⁻²] (energy)
	expectedDim := units.Dimension{L: 2, M: 1, T: -2}
	if work.Dim() != expectedDim {
		t.Errorf("Work dimension = %v, want %v (energy)", work.Dim(), expectedDim)
	}

	// Expected value: 10 N × 5 m = 50 J
	if !almostEqual(work.Val(), 50.0, 1e-10) {
		t.Errorf("Work = %v J, want 50 J", work.Val())
	}
}

// -----------------------------------------------------------------------------
// Cross Product Tests
// -----------------------------------------------------------------------------

func TestCross(t *testing.T) {
	// i × j = k
	i := NewPosition(units.Meter(1), units.Meter(0), units.Meter(0))
	j := NewPosition(units.Meter(0), units.Meter(1), units.Meter(0))
	k := i.Cross(j)

	if !almostEqual(k.X.Val(), 0, 1e-10) || !almostEqual(k.Y.Val(), 0, 1e-10) || !almostEqual(k.Z.Val(), 1, 1e-10) {
		t.Errorf("i × j = (%v, %v, %v), want (0, 0, 1)", k.X.Val(), k.Y.Val(), k.Z.Val())
	}

	// j × k = i
	k2 := NewPosition(units.Meter(0), units.Meter(0), units.Meter(1))
	i2 := j.Cross(k2)
	if !almostEqual(i2.X.Val(), 1, 1e-10) || !almostEqual(i2.Y.Val(), 0, 1e-10) || !almostEqual(i2.Z.Val(), 0, 1e-10) {
		t.Errorf("j × k = (%v, %v, %v), want (1, 0, 0)", i2.X.Val(), i2.Y.Val(), i2.Z.Val())
	}

	// Anticommutative: v × w = -(w × v)
	v := NewPosition(units.Meter(1), units.Meter(2), units.Meter(3))
	w := NewPosition(units.Meter(4), units.Meter(5), units.Meter(6))
	vw := v.Cross(w)
	wv := w.Cross(v)
	neg_wv := wv.Negate()

	if !almostEqual(vw.X.Val(), neg_wv.X.Val(), 1e-10) ||
		!almostEqual(vw.Y.Val(), neg_wv.Y.Val(), 1e-10) ||
		!almostEqual(vw.Z.Val(), neg_wv.Z.Val(), 1e-10) {
		t.Error("Cross product should be anticommutative")
	}
}

func TestCrossTorque(t *testing.T) {
	// Torque: τ = r × F
	r := NewPosition(
		units.Meter(1),
		units.Meter(0),
		units.Meter(0),
	)
	F := NewForce(
		units.Newton(0),
		units.Newton(10),
		units.Newton(0),
	)

	torque := r.Cross(F)

	// Expected dimension: [L²MT⁻²] (energy/angular momentum)
	expectedDim := units.Dimension{L: 2, M: 1, T: -2}
	if torque.Dim() != expectedDim {
		t.Errorf("Torque dimension = %v, want %v", torque.Dim(), expectedDim)
	}

	// r × F = (1, 0, 0) × (0, 10, 0) = (0, 0, 10)
	if !almostEqual(torque.X.Val(), 0, 1e-10) ||
		!almostEqual(torque.Y.Val(), 0, 1e-10) ||
		!almostEqual(torque.Z.Val(), 10, 1e-10) {
		t.Errorf("Torque = (%v, %v, %v), want (0, 0, 10)", torque.X.Val(), torque.Y.Val(), torque.Z.Val())
	}
}

// -----------------------------------------------------------------------------
// Magnitude Tests
// -----------------------------------------------------------------------------

func TestMagnitude(t *testing.T) {
	// 3-4-5 triangle
	v := NewPosition(units.Meter(3), units.Meter(4), units.Meter(0))
	mag, err := v.Magnitude()
	if err != nil {
		t.Errorf("Magnitude() failed: %v", err)
	}

	if !almostEqual(mag.Val(), 5.0, 1e-10) {
		t.Errorf("Magnitude() = %v, want 5", mag.Val())
	}

	// Dimension should be same as components
	if mag.Dim() != v.Dim() {
		t.Errorf("Magnitude dimension = %v, want %v", mag.Dim(), v.Dim())
	}
}

func TestMagnitudeSquared(t *testing.T) {
	v := NewVelocity(
		units.MeterPerSecond(3),
		units.MeterPerSecond(4),
		units.MeterPerSecond(0),
	)

	magSq := v.MagnitudeSquared()

	// 3² + 4² = 25 (m/s)²
	if !almostEqual(magSq.Val(), 25.0, 1e-10) {
		t.Errorf("MagnitudeSquared() = %v, want 25", magSq.Val())
	}

	// Dimension should be squared
	expectedDim := units.Dimension{L: 2, T: -2}
	if magSq.Dim() != expectedDim {
		t.Errorf("MagnitudeSquared dimension = %v, want %v", magSq.Dim(), expectedDim)
	}
}

// -----------------------------------------------------------------------------
// Angle Tests
// -----------------------------------------------------------------------------

func TestAngleBetween(t *testing.T) {
	tests := []struct {
		name          string
		v1            Vector3
		v2            Vector3
		expectedAngle float64 // in radians
	}{
		{
			name:          "perpendicular",
			v1:            NewPosition(units.Meter(1), units.Meter(0), units.Meter(0)),
			v2:            NewPosition(units.Meter(0), units.Meter(1), units.Meter(0)),
			expectedAngle: math.Pi / 2, // 90 degrees
		},
		{
			name:          "parallel",
			v1:            NewPosition(units.Meter(1), units.Meter(0), units.Meter(0)),
			v2:            NewPosition(units.Meter(2), units.Meter(0), units.Meter(0)),
			expectedAngle: 0, // 0 degrees
		},
		{
			name:          "antiparallel",
			v1:            NewPosition(units.Meter(1), units.Meter(0), units.Meter(0)),
			v2:            NewPosition(units.Meter(-1), units.Meter(0), units.Meter(0)),
			expectedAngle: math.Pi, // 180 degrees
		},
		{
			name:          "45 degrees",
			v1:            NewPosition(units.Meter(1), units.Meter(0), units.Meter(0)),
			v2:            NewPosition(units.Meter(1), units.Meter(1), units.Meter(0)),
			expectedAngle: math.Pi / 4, // 45 degrees
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			angle, err := tt.v1.AngleBetween(tt.v2)
			if err != nil {
				t.Errorf("AngleBetween() failed: %v", err)
			}

			if !almostEqual(angle, tt.expectedAngle, 1e-10) {
				t.Errorf("AngleBetween() = %v rad, want %v rad", angle, tt.expectedAngle)
			}
		})
	}
}

// -----------------------------------------------------------------------------
// Projection Tests
// -----------------------------------------------------------------------------

func TestProjectOnto(t *testing.T) {
	// Project (3, 4, 0) onto (1, 0, 0)
	// Should give (3, 0, 0)
	v := NewPosition(units.Meter(3), units.Meter(4), units.Meter(0))
	axis := NewPosition(units.Meter(1), units.Meter(0), units.Meter(0))

	proj, err := v.ProjectOnto(axis)
	if err != nil {
		t.Errorf("ProjectOnto() failed: %v", err)
	}

	if !almostEqual(proj.X.Val(), 3, 1e-10) ||
		!almostEqual(proj.Y.Val(), 0, 1e-10) ||
		!almostEqual(proj.Z.Val(), 0, 1e-10) {
		t.Errorf("ProjectOnto() = (%v, %v, %v), want (3, 0, 0)",
			proj.X.Val(), proj.Y.Val(), proj.Z.Val())
	}
}

// -----------------------------------------------------------------------------
// Geometric Relations Tests
// -----------------------------------------------------------------------------

func TestIsParallel(t *testing.T) {
	v1 := NewPosition(units.Meter(1), units.Meter(2), units.Meter(3))
	v2 := NewPosition(units.Meter(2), units.Meter(4), units.Meter(6)) // Parallel: 2*v1
	v3 := NewPosition(units.Meter(1), units.Meter(0), units.Meter(0)) // Not parallel

	if !v1.IsParallel(v2, 1e-10) {
		t.Error("v1 and v2 should be parallel")
	}

	if v1.IsParallel(v3, 1e-10) {
		t.Error("v1 and v3 should not be parallel")
	}
}

func TestIsPerpendicular(t *testing.T) {
	v1 := NewPosition(units.Meter(1), units.Meter(0), units.Meter(0))
	v2 := NewPosition(units.Meter(0), units.Meter(1), units.Meter(0)) // Perpendicular
	v3 := NewPosition(units.Meter(1), units.Meter(1), units.Meter(0)) // Not perpendicular

	if !v1.IsPerpendicular(v2, 1e-10) {
		t.Error("v1 and v2 should be perpendicular")
	}

	if v1.IsPerpendicular(v3, 1e-10) {
		t.Error("v1 and v3 should not be perpendicular")
	}
}

// -----------------------------------------------------------------------------
// Physical Application Tests
// -----------------------------------------------------------------------------

func TestPhysics_Projectile(t *testing.T) {
	// Projectile motion: r(t) = r0 + v0*t + ½*a*t²
	r0 := NewPosition(units.Meter(0), units.Meter(0), units.Meter(0))
	v0 := NewVelocity(units.MeterPerSecond(10), units.MeterPerSecond(10), units.MeterPerSecond(0))
	a := NewAcceleration(units.MeterPerSecond2(0), units.MeterPerSecond2(-9.8), units.MeterPerSecond2(0))
	time := units.Second(1.0)

	// v0 * t
	v0t := Vector3{
		X: v0.X.Multiply(time.Value),
		Y: v0.Y.Multiply(time.Value),
		Z: v0.Z.Multiply(time.Value),
	}

	// ½ * a * t²
	tSquared := time.Value.Power(2)
	halfAt2 := Vector3{
		X: a.X.Multiply(tSquared).Scale(0.5),
		Y: a.Y.Multiply(tSquared).Scale(0.5),
		Z: a.Z.Multiply(tSquared).Scale(0.5),
	}

	// r = r0 + v0*t + ½*a*t²
	r1, _ := r0.Add(v0t)
	r, _ := r1.Add(halfAt2)

	// After 1 second: x = 10 m, y = 10 - 4.9 = 5.1 m
	if !almostEqual(r.X.Val(), 10.0, 1e-10) {
		t.Errorf("x position = %v m, want 10 m", r.X.Val())
	}
	if !almostEqual(r.Y.Val(), 5.1, 1e-10) {
		t.Errorf("y position = %v m, want 5.1 m", r.Y.Val())
	}
}

func TestPhysics_AngularMomentum(t *testing.T) {
	// L = r × p (angular momentum)
	r := NewPosition(units.Meter(1), units.Meter(0), units.Meter(0))

	// Linear momentum: p = m*v
	m := units.Kilogram(2.0)
	v := NewVelocity(units.MeterPerSecond(0), units.MeterPerSecond(5), units.MeterPerSecond(0))
	p := Vector3{
		X: m.Value.Multiply(v.X),
		Y: m.Value.Multiply(v.Y),
		Z: m.Value.Multiply(v.Z),
	}

	// Angular momentum
	L := r.Cross(p)

	// L_z = r_x * p_y - r_y * p_x = 1 * 10 - 0 * 0 = 10 kg⋅m²/s
	if !almostEqual(L.Z.Val(), 10.0, 1e-10) {
		t.Errorf("Angular momentum L_z = %v, want 10", L.Z.Val())
	}

	// Dimension should be [L²MT⁻¹] (angular momentum)
	expectedDim := units.Dimension{L: 2, M: 1, T: -1}
	if L.Dim() != expectedDim {
		t.Errorf("Angular momentum dimension = %v, want %v", L.Dim(), expectedDim)
	}
}
