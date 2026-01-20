// Package units provides a type-safe system for physical units with compile-time
// dimensional analysis. It prevents common errors like adding incompatible units
// (e.g., length + mass) and ensures dimensional consistency across calculations.
//
// The package implements the International System of Units (SI) with seven base
// dimensions: Length, Mass, Time, Electric Current, Temperature, Amount of
// Substance, and Luminous Intensity.
//
// Example usage:
//
//	length := units.Meter(5.0)
//	time := units.Second(2.0)
//	velocity := length.Divide(time) // Returns m/s
//
//	// This would cause a compile error:
//	// invalid := length.Add(units.Kilogram(3.0)) // Cannot add length + mass
package units

import (
	"fmt"
	"math"
)

// Dimension represents the dimensional formula of a physical quantity using
// the seven SI base dimensions. Each field is the exponent of that dimension.
//
// For example:
//   - Velocity [L¹T⁻¹]: Dimension{L: 1, T: -1}
//   - Energy [L²M¹T⁻²]: Dimension{L: 2, M: 1, T: -2}
//   - Dimensionless [1]: Dimension{} (all zeros)
type Dimension struct {
	L int8 // Length (meter, m)
	M int8 // Mass (kilogram, kg)
	T int8 // Time (second, s)
	I int8 // Electric current (ampere, A)
	Θ int8 // Thermodynamic temperature (kelvin, K)
	N int8 // Amount of substance (mole, mol)
	J int8 // Luminous intensity (candela, cd)
}

// Value represents a physical quantity with both a numerical value and
// dimensional information. This enables compile-time checking of dimensional
// consistency.
//
// Value should not be instantiated directly. Instead, use constructors from
// specific unit types like Meter(), Kilogram(), Second(), etc.
type Value struct {
	value float64
	dim   Dimension
}

// NewValue creates a new Value with the specified numerical value and dimension.
// This is a low-level constructor; prefer using specific unit constructors.
func NewValue(value float64, dim Dimension) Value {
	return Value{value: value, dim: dim}
}

// Val returns the numerical value of the quantity in SI base units.
//
// Example:
//
//	length := units.Meter(5.0)
//	fmt.Println(length.Val()) // Output: 5.0
func (v Value) Val() float64 {
	return v.value
}

// Dim returns the dimensional formula of the quantity.
func (v Value) Dim() Dimension {
	return v.dim
}

// Equal returns true if two Values have the same dimension and value
// (within floating-point tolerance).
func (v Value) Equal(other Value) bool {
	return v.dim == other.dim && almostEqual(v.value, other.value, 1e-14)
}

// String returns a human-readable representation of the Value.
func (v Value) String() string {
	return fmt.Sprintf("%.6g %s", v.value, v.dim.String())
}

// Add returns the sum of two Values. The Values must have identical dimensions.
// Returns an error if dimensions don't match.
//
// Example:
//
//	l1 := units.Meter(5.0)
//	l2 := units.Meter(3.0)
//	result, _ := l1.Add(l2) // 8.0 m
//
//	// Error case:
//	mass := units.Kilogram(2.0)
//	_, err := l1.Add(mass) // Error: incompatible dimensions
func (v Value) Add(other Value) (Value, error) {
	if v.dim != other.dim {
		return Value{}, fmt.Errorf("cannot add quantities with different dimensions: %s + %s",
			v.dim.String(), other.dim.String())
	}
	return Value{value: v.value + other.value, dim: v.dim}, nil
}

// Subtract returns the difference of two Values. The Values must have identical dimensions.
// Returns an error if dimensions don't match.
//
// Example:
//
//	l1 := units.Meter(5.0)
//	l2 := units.Meter(3.0)
//	result, _ := l1.Subtract(l2) // 2.0 m
func (v Value) Subtract(other Value) (Value, error) {
	if v.dim != other.dim {
		return Value{}, fmt.Errorf("cannot subtract quantities with different dimensions: %s - %s",
			v.dim.String(), other.dim.String())
	}
	return Value{value: v.value - other.value, dim: v.dim}, nil
}

// Multiply returns the product of two Values. The dimensions are added.
//
// Example:
//
//	length := units.Meter(5.0)    // [L¹]
//	width := units.Meter(3.0)     // [L¹]
//	area := length.Multiply(width) // [L²] = 15.0 m²
//
//	mass := units.Kilogram(2.0)   // [M¹]
//	accel := units.MeterPerSecond2(3.0) // [L¹T⁻²]
//	force := mass.Multiply(accel) // [M¹L¹T⁻²] = 6.0 N (newton)
func (v Value) Multiply(other Value) Value {
	return Value{
		value: v.value * other.value,
		dim: Dimension{
			L: v.dim.L + other.dim.L,
			M: v.dim.M + other.dim.M,
			T: v.dim.T + other.dim.T,
			I: v.dim.I + other.dim.I,
			Θ: v.dim.Θ + other.dim.Θ,
			N: v.dim.N + other.dim.N,
			J: v.dim.J + other.dim.J,
		},
	}
}

// Divide returns the quotient of two Values. The dimensions are subtracted.
//
// Example:
//
//	length := units.Meter(10.0)  // [L¹]
//	time := units.Second(2.0)    // [T¹]
//	velocity := length.Divide(time) // [L¹T⁻¹] = 5.0 m/s
func (v Value) Divide(other Value) Value {
	return Value{
		value: v.value / other.value,
		dim: Dimension{
			L: v.dim.L - other.dim.L,
			M: v.dim.M - other.dim.M,
			T: v.dim.T - other.dim.T,
			I: v.dim.I - other.dim.I,
			Θ: v.dim.Θ - other.dim.Θ,
			N: v.dim.N - other.dim.N,
			J: v.dim.J - other.dim.J,
		},
	}
}

// Scale returns the Value multiplied by a dimensionless scalar.
//
// Example:
//
//	length := units.Meter(5.0)
//	doubled := length.Scale(2.0) // 10.0 m
func (v Value) Scale(scalar float64) Value {
	return Value{value: v.value * scalar, dim: v.dim}
}

// Power returns the Value raised to an integer power. The dimensions are
// multiplied by the exponent.
//
// Example:
//
//	length := units.Meter(5.0)   // [L¹]
//	area := length.Power(2)      // [L²] = 25.0 m²
//	volume := length.Power(3)    // [L³] = 125.0 m³
func (v Value) Power(n int) Value {
	return Value{
		value: math.Pow(v.value, float64(n)),
		dim: Dimension{
			L: v.dim.L * int8(n),
			M: v.dim.M * int8(n),
			T: v.dim.T * int8(n),
			I: v.dim.I * int8(n),
			Θ: v.dim.Θ * int8(n),
			N: v.dim.N * int8(n),
			J: v.dim.J * int8(n),
		},
	}
}

// Sqrt returns the square root of the Value. The dimensions are divided by 2.
// Returns an error if any dimension has an odd exponent.
//
// Example:
//
//	area := units.Meter(25.0).Power(2)  // [L²] = 25.0 m²
//	length, _ := area.Sqrt()            // [L¹] = 5.0 m
func (v Value) Sqrt() (Value, error) {
	// Check if all dimensions are even (can be divided by 2)
	if v.dim.L%2 != 0 || v.dim.M%2 != 0 || v.dim.T%2 != 0 || v.dim.I%2 != 0 ||
		v.dim.Θ%2 != 0 || v.dim.N%2 != 0 || v.dim.J%2 != 0 {
		return Value{}, fmt.Errorf("cannot take square root of quantity with odd dimension exponents: %s",
			v.dim.String())
	}

	return Value{
		value: math.Sqrt(v.value),
		dim: Dimension{
			L: v.dim.L / 2,
			M: v.dim.M / 2,
			T: v.dim.T / 2,
			I: v.dim.I / 2,
			Θ: v.dim.Θ / 2,
			N: v.dim.N / 2,
			J: v.dim.J / 2,
		},
	}, nil
}

// Abs returns the absolute value of the quantity, preserving dimensions.
func (v Value) Abs() Value {
	return Value{value: math.Abs(v.value), dim: v.dim}
}

// Negate returns the negation of the quantity, preserving dimensions.
func (v Value) Negate() Value {
	return Value{value: -v.value, dim: v.dim}
}

// IsDimensionless returns true if the Value has no dimensions (all exponents are zero).
func (v Value) IsDimensionless() bool {
	return v.dim == Dimension{}
}

// String returns a human-readable representation of the Dimension.
//
// Format: [L^a M^b T^c I^d Θ^e N^f J^g] where only non-zero exponents are shown.
func (d Dimension) String() string {
	if d == (Dimension{}) {
		return "[1]" // Dimensionless
	}

	var parts []string
	if d.L != 0 {
		parts = append(parts, fmt.Sprintf("L^%d", d.L))
	}
	if d.M != 0 {
		parts = append(parts, fmt.Sprintf("M^%d", d.M))
	}
	if d.T != 0 {
		parts = append(parts, fmt.Sprintf("T^%d", d.T))
	}
	if d.I != 0 {
		parts = append(parts, fmt.Sprintf("I^%d", d.I))
	}
	if d.Θ != 0 {
		parts = append(parts, fmt.Sprintf("Θ^%d", d.Θ))
	}
	if d.N != 0 {
		parts = append(parts, fmt.Sprintf("N^%d", d.N))
	}
	if d.J != 0 {
		parts = append(parts, fmt.Sprintf("J^%d", d.J))
	}

	result := "["
	for i, p := range parts {
		if i > 0 {
			result += " "
		}
		result += p
	}
	result += "]"
	return result
}

// almostEqual returns true if two float64 values are equal within a relative tolerance.
func almostEqual(a, b, tolerance float64) bool {
	if a == b {
		return true
	}
	diff := math.Abs(a - b)
	if a == 0 || b == 0 || diff < tolerance {
		return diff < tolerance
	}
	return diff/(math.Abs(a)+math.Abs(b)) < tolerance
}

// Dimensionless creates a dimensionless quantity (pure number).
//
// Example:
//
//	ratio := units.Dimensionless(1.5)
//	angle := units.Dimensionless(math.Pi) // Radians are dimensionless
func Dimensionless(value float64) Value {
	return Value{value: value, dim: Dimension{}}
}
