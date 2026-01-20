package units

import (
	"testing"
)

// -----------------------------------------------------------------------------
// Dimension Tests
// -----------------------------------------------------------------------------

func TestDimensionString(t *testing.T) {
	tests := []struct {
		name string
		dim  Dimension
		want string
	}{
		{
			name: "dimensionless",
			dim:  Dimension{},
			want: "[1]",
		},
		{
			name: "length",
			dim:  Dimension{L: 1},
			want: "[L^1]",
		},
		{
			name: "velocity",
			dim:  Dimension{L: 1, T: -1},
			want: "[L^1 T^-1]",
		},
		{
			name: "energy",
			dim:  Dimension{L: 2, M: 1, T: -2},
			want: "[L^2 M^1 T^-2]",
		},
		{
			name: "all dimensions",
			dim:  Dimension{L: 1, M: 2, T: -3, I: 4, Θ: -5, N: 6, J: -7},
			want: "[L^1 M^2 T^-3 I^4 Θ^-5 N^6 J^-7]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.dim.String()
			if got != tt.want {
				t.Errorf("Dimension.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

// -----------------------------------------------------------------------------
// Value Basic Operations Tests
// -----------------------------------------------------------------------------

func TestValueAdd(t *testing.T) {
	tests := []struct {
		name    string
		v1      Value
		v2      Value
		want    Value
		wantErr bool
	}{
		{
			name:    "same dimension",
			v1:      Meter(5.0).Value,
			v2:      Meter(3.0).Value,
			want:    Meter(8.0).Value,
			wantErr: false,
		},
		{
			name:    "different dimensions",
			v1:      Meter(5.0).Value,
			v2:      Kilogram(3.0).Value,
			want:    Value{},
			wantErr: true,
		},
		{
			name:    "negative values",
			v1:      Meter(-5.0).Value,
			v2:      Meter(3.0).Value,
			want:    Meter(-2.0).Value,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v1.Add(tt.v2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Value.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !got.Equal(tt.want) {
				t.Errorf("Value.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValueSubtract(t *testing.T) {
	tests := []struct {
		name    string
		v1      Value
		v2      Value
		want    Value
		wantErr bool
	}{
		{
			name:    "same dimension",
			v1:      Meter(5.0).Value,
			v2:      Meter(3.0).Value,
			want:    Meter(2.0).Value,
			wantErr: false,
		},
		{
			name:    "different dimensions",
			v1:      Meter(5.0).Value,
			v2:      Second(3.0).Value,
			want:    Value{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v1.Subtract(tt.v2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Value.Subtract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !got.Equal(tt.want) {
				t.Errorf("Value.Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValueMultiply(t *testing.T) {
	tests := []struct {
		name    string
		v1      Value
		v2      Value
		wantDim Dimension
		wantVal float64
	}{
		{
			name:    "length × length = area",
			v1:      Meter(5.0).Value,
			v2:      Meter(3.0).Value,
			wantDim: Dimension{L: 2},
			wantVal: 15.0,
		},
		{
			name:    "length × mass = [L M]",
			v1:      Meter(2.0).Value,
			v2:      Kilogram(3.0).Value,
			wantDim: Dimension{L: 1, M: 1},
			wantVal: 6.0,
		},
		{
			name:    "velocity × time = length",
			v1:      NewValue(10.0, Dimension{L: 1, T: -1}), // 10 m/s
			v2:      Second(2.0).Value,
			wantDim: Dimension{L: 1},
			wantVal: 20.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.v1.Multiply(tt.v2)
			if got.dim != tt.wantDim {
				t.Errorf("Value.Multiply() dimension = %v, want %v", got.dim, tt.wantDim)
			}
			if !almostEqual(got.value, tt.wantVal, 1e-14) {
				t.Errorf("Value.Multiply() value = %v, want %v", got.value, tt.wantVal)
			}
		})
	}
}

func TestValueDivide(t *testing.T) {
	tests := []struct {
		name    string
		v1      Value
		v2      Value
		wantDim Dimension
		wantVal float64
	}{
		{
			name:    "length / time = velocity",
			v1:      Meter(10.0).Value,
			v2:      Second(2.0).Value,
			wantDim: Dimension{L: 1, T: -1},
			wantVal: 5.0,
		},
		{
			name:    "area / length = length",
			v1:      NewValue(20.0, Dimension{L: 2}),
			v2:      Meter(4.0).Value,
			wantDim: Dimension{L: 1},
			wantVal: 5.0,
		},
		{
			name:    "same dimension = dimensionless",
			v1:      Meter(10.0).Value,
			v2:      Meter(2.0).Value,
			wantDim: Dimension{},
			wantVal: 5.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.v1.Divide(tt.v2)
			if got.dim != tt.wantDim {
				t.Errorf("Value.Divide() dimension = %v, want %v", got.dim, tt.wantDim)
			}
			if !almostEqual(got.value, tt.wantVal, 1e-14) {
				t.Errorf("Value.Divide() value = %v, want %v", got.value, tt.wantVal)
			}
		})
	}
}

func TestValueScale(t *testing.T) {
	tests := []struct {
		name   string
		value  Value
		scalar float64
		want   float64
	}{
		{
			name:   "scale positive",
			value:  Meter(5.0).Value,
			scalar: 2.0,
			want:   10.0,
		},
		{
			name:   "scale negative",
			value:  Meter(5.0).Value,
			scalar: -2.0,
			want:   -10.0,
		},
		{
			name:   "scale fractional",
			value:  Kilogram(10.0).Value,
			scalar: 0.5,
			want:   5.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.value.Scale(tt.scalar)
			if !almostEqual(got.value, tt.want, 1e-14) {
				t.Errorf("Value.Scale() = %v, want %v", got.value, tt.want)
			}
			if got.dim != tt.value.dim {
				t.Errorf("Value.Scale() changed dimension: %v -> %v", tt.value.dim, got.dim)
			}
		})
	}
}

func TestValuePower(t *testing.T) {
	tests := []struct {
		name    string
		value   Value
		n       int
		wantDim Dimension
		wantVal float64
	}{
		{
			name:    "length^2 = area",
			value:   Meter(5.0).Value,
			n:       2,
			wantDim: Dimension{L: 2},
			wantVal: 25.0,
		},
		{
			name:    "length^3 = volume",
			value:   Meter(2.0).Value,
			n:       3,
			wantDim: Dimension{L: 3},
			wantVal: 8.0,
		},
		{
			name:    "power 0 = dimensionless 1",
			value:   Meter(5.0).Value,
			n:       0,
			wantDim: Dimension{},
			wantVal: 1.0,
		},
		{
			name:    "power -1 = inverse",
			value:   Meter(2.0).Value,
			n:       -1,
			wantDim: Dimension{L: -1},
			wantVal: 0.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.value.Power(tt.n)
			if got.dim != tt.wantDim {
				t.Errorf("Value.Power() dimension = %v, want %v", got.dim, tt.wantDim)
			}
			if !almostEqual(got.value, tt.wantVal, 1e-14) {
				t.Errorf("Value.Power() value = %v, want %v", got.value, tt.wantVal)
			}
		})
	}
}

func TestValueSqrt(t *testing.T) {
	tests := []struct {
		name    string
		value   Value
		wantDim Dimension
		wantVal float64
		wantErr bool
	}{
		{
			name:    "sqrt of area = length",
			value:   NewValue(25.0, Dimension{L: 2}),
			wantDim: Dimension{L: 1},
			wantVal: 5.0,
			wantErr: false,
		},
		{
			name:    "sqrt of volume^2 = volume",
			value:   NewValue(64.0, Dimension{L: 6}),
			wantDim: Dimension{L: 3},
			wantVal: 8.0,
			wantErr: false,
		},
		{
			name:    "sqrt of odd dimension fails",
			value:   Meter(5.0).Value,
			wantDim: Dimension{},
			wantVal: 0.0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.value.Sqrt()
			if (err != nil) != tt.wantErr {
				t.Errorf("Value.Sqrt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got.dim != tt.wantDim {
					t.Errorf("Value.Sqrt() dimension = %v, want %v", got.dim, tt.wantDim)
				}
				if !almostEqual(got.value, tt.wantVal, 1e-14) {
					t.Errorf("Value.Sqrt() value = %v, want %v", got.value, tt.wantVal)
				}
			}
		})
	}
}

func TestValueAbs(t *testing.T) {
	tests := []struct {
		name  string
		value Value
		want  float64
	}{
		{
			name:  "positive value",
			value: Meter(5.0).Value,
			want:  5.0,
		},
		{
			name:  "negative value",
			value: Meter(-5.0).Value,
			want:  5.0,
		},
		{
			name:  "zero",
			value: Meter(0.0).Value,
			want:  0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.value.Abs()
			if !almostEqual(got.value, tt.want, 1e-14) {
				t.Errorf("Value.Abs() = %v, want %v", got.value, tt.want)
			}
		})
	}
}

func TestValueNegate(t *testing.T) {
	tests := []struct {
		name  string
		value Value
		want  float64
	}{
		{
			name:  "positive value",
			value: Meter(5.0).Value,
			want:  -5.0,
		},
		{
			name:  "negative value",
			value: Meter(-5.0).Value,
			want:  5.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.value.Negate()
			if !almostEqual(got.value, tt.want, 1e-14) {
				t.Errorf("Value.Negate() = %v, want %v", got.value, tt.want)
			}
		})
	}
}

func TestValueIsDimensionless(t *testing.T) {
	tests := []struct {
		name  string
		value Value
		want  bool
	}{
		{
			name:  "dimensionless",
			value: Dimensionless(5.0),
			want:  true,
		},
		{
			name:  "length",
			value: Meter(5.0).Value,
			want:  false,
		},
		{
			name:  "ratio (dimensionless from division)",
			value: Meter(10.0).Value.Divide(Meter(2.0).Value),
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.value.IsDimensionless()
			if got != tt.want {
				t.Errorf("Value.IsDimensionless() = %v, want %v", got, tt.want)
			}
		})
	}
}

// -----------------------------------------------------------------------------
// SI Base Unit Tests
// -----------------------------------------------------------------------------

func TestLength(t *testing.T) {
	tests := []struct {
		name     string
		length   Length
		wantVal  float64
		wantUnit string
	}{
		{"meter", Meter(1.0), 1.0, "m"},
		{"millimeter", Millimeter(1000.0), 1.0, "m"},
		{"centimeter", Centimeter(100.0), 1.0, "m"},
		{"kilometer", Kilometer(0.001), 1.0, "m"},
		{"micrometer", Micrometer(1e6), 1.0, "m"},
		{"nanometer", Nanometer(1e9), 1.0, "m"},
		{"angstrom", Angstrom(1e10), 1.0, "m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !almostEqual(tt.length.Val(), tt.wantVal, 1e-10) {
				t.Errorf("%s = %v m, want %v m", tt.name, tt.length.Val(), tt.wantVal)
			}
			if tt.length.Dim() != (Dimension{L: 1}) {
				t.Errorf("%s has incorrect dimension: %v", tt.name, tt.length.Dim())
			}
		})
	}
}

func TestMass(t *testing.T) {
	tests := []struct {
		name    string
		mass    Mass
		wantVal float64
	}{
		{"kilogram", Kilogram(1.0), 1.0},
		{"gram", Gram(1000.0), 1.0},
		{"milligram", Milligram(1e6), 1.0},
		{"microgram", Microgram(1e9), 1.0},
		{"tonne", Tonne(0.001), 1.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !almostEqual(tt.mass.Val(), tt.wantVal, 1e-10) {
				t.Errorf("%s = %v kg, want %v kg", tt.name, tt.mass.Val(), tt.wantVal)
			}
			if tt.mass.Dim() != (Dimension{M: 1}) {
				t.Errorf("%s has incorrect dimension: %v", tt.name, tt.mass.Dim())
			}
		})
	}
}

func TestTime(t *testing.T) {
	tests := []struct {
		name    string
		time    Time
		wantVal float64
	}{
		{"second", Second(1.0), 1.0},
		{"millisecond", Millisecond(1000.0), 1.0},
		{"microsecond", Microsecond(1e6), 1.0},
		{"nanosecond", Nanosecond(1e9), 1.0},
		{"minute", Minute(1.0 / 60.0), 1.0},
		{"hour", Hour(1.0 / 3600.0), 1.0},
		{"day", Day(1.0 / 86400.0), 1.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !almostEqual(tt.time.Val(), tt.wantVal, 1e-10) {
				t.Errorf("%s = %v s, want %v s", tt.name, tt.time.Val(), tt.wantVal)
			}
			if tt.time.Dim() != (Dimension{T: 1}) {
				t.Errorf("%s has incorrect dimension: %v", tt.name, tt.time.Dim())
			}
		})
	}
}

func TestCurrent(t *testing.T) {
	ampere := Ampere(1.0)
	if !almostEqual(ampere.Val(), 1.0, 1e-14) {
		t.Errorf("Ampere(1.0) = %v A, want 1.0 A", ampere.Val())
	}
	if ampere.Dim() != (Dimension{I: 1}) {
		t.Errorf("Ampere has incorrect dimension: %v", ampere.Dim())
	}

	milliampere := Milliampere(1000.0)
	if !almostEqual(milliampere.Val(), 1.0, 1e-10) {
		t.Errorf("Milliampere(1000.0) = %v A, want 1.0 A", milliampere.Val())
	}
}

func TestTemperature(t *testing.T) {
	tests := []struct {
		name    string
		temp    Temperature
		wantVal float64
	}{
		{"kelvin", Kelvin(273.15), 273.15},
		{"celsius zero", Celsius(0.0), 273.15},
		{"celsius 100", Celsius(100.0), 373.15},
		{"fahrenheit 32", Fahrenheit(32.0), 273.15},
		{"fahrenheit 212", Fahrenheit(212.0), 373.15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !almostEqual(tt.temp.Val(), tt.wantVal, 1e-10) {
				t.Errorf("%s = %v K, want %v K", tt.name, tt.temp.Val(), tt.wantVal)
			}
			if tt.temp.Dim() != (Dimension{Θ: 1}) {
				t.Errorf("%s has incorrect dimension: %v", tt.name, tt.temp.Dim())
			}
		})
	}
}

func TestAmount(t *testing.T) {
	mole := Mole(1.0)
	if !almostEqual(mole.Val(), 1.0, 1e-14) {
		t.Errorf("Mole(1.0) = %v mol, want 1.0 mol", mole.Val())
	}
	if mole.Dim() != (Dimension{N: 1}) {
		t.Errorf("Mole has incorrect dimension: %v", mole.Dim())
	}

	millimole := Millimole(1000.0)
	if !almostEqual(millimole.Val(), 1.0, 1e-10) {
		t.Errorf("Millimole(1000.0) = %v mol, want 1.0 mol", millimole.Val())
	}
}

func TestLuminousIntensity(t *testing.T) {
	candela := Candela(1.0)
	if !almostEqual(candela.Val(), 1.0, 1e-14) {
		t.Errorf("Candela(1.0) = %v cd, want 1.0 cd", candela.Val())
	}
	if candela.Dim() != (Dimension{J: 1}) {
		t.Errorf("Candela has incorrect dimension: %v", candela.Dim())
	}
}

// -----------------------------------------------------------------------------
// Astronomical Unit Tests
// -----------------------------------------------------------------------------

func TestAstronomicalUnits(t *testing.T) {
	au := AstronomicalUnit(1.0)
	if !almostEqual(au.Val(), 1.495978707e11, 1e-5) {
		t.Errorf("AstronomicalUnit(1.0) = %v m, want 1.495978707e11 m", au.Val())
	}

	ly := LightYear(1.0)
	if !almostEqual(ly.Val(), 9.4607304725808e15, 1e-5) {
		t.Errorf("LightYear(1.0) = %v m, want 9.4607304725808e15 m", ly.Val())
	}

	pc := Parsec(1.0)
	if !almostEqual(pc.Val(), 3.0856775814913673e16, 1e-5) {
		t.Errorf("Parsec(1.0) = %v m, want 3.0856775814913673e16 m", pc.Val())
	}
}

func TestSolarMass(t *testing.T) {
	sm := SolarMass(1.0)
	if !almostEqual(sm.Val(), 1.98892e30, 1e-5) {
		t.Errorf("SolarMass(1.0) = %v kg, want 1.98892e30 kg", sm.Val())
	}
}

// -----------------------------------------------------------------------------
// Type Safety Tests
// -----------------------------------------------------------------------------

func TestTypeSafety_AddDifferentDimensions(t *testing.T) {
	length := Meter(5.0)
	mass := Kilogram(3.0)

	_, err := length.Value.Add(mass.Value)
	if err == nil {
		t.Error("Expected error when adding length and mass, got nil")
	}
}

func TestTypeSafety_SubtractDifferentDimensions(t *testing.T) {
	time := Second(10.0)
	current := Ampere(2.0)

	_, err := time.Value.Subtract(current.Value)
	if err == nil {
		t.Error("Expected error when subtracting time and current, got nil")
	}
}

func TestTypeSafety_MultiplyPreservesDimensions(t *testing.T) {
	length := Meter(5.0)
	time := Second(2.0)

	// length × time should give [L¹T¹]
	result := length.Value.Multiply(time.Value)
	expected := Dimension{L: 1, T: 1}

	if result.Dim() != expected {
		t.Errorf("Length × Time = %v, want %v", result.Dim(), expected)
	}
}

func TestTypeSafety_DivideCreatesNewDimension(t *testing.T) {
	length := Meter(10.0)
	time := Second(2.0)

	// length / time should give velocity [L¹T⁻¹]
	velocity := length.Value.Divide(time.Value)
	expected := Dimension{L: 1, T: -1}

	if velocity.Dim() != expected {
		t.Errorf("Length / Time = %v, want %v (velocity)", velocity.Dim(), expected)
	}

	if !almostEqual(velocity.Val(), 5.0, 1e-14) {
		t.Errorf("Velocity value = %v m/s, want 5.0 m/s", velocity.Val())
	}
}

// -----------------------------------------------------------------------------
// Complex Calculation Tests
// -----------------------------------------------------------------------------

func TestComplexCalculation_KineticEnergy(t *testing.T) {
	// KE = ½mv²
	// m = 2 kg, v = 3 m/s
	// KE = 0.5 × 2 × 9 = 9 J = 9 kg⋅m²/s²

	mass := Kilogram(2.0)
	velocity := Meter(3.0).Value.Divide(Second(1.0).Value) // 3 m/s

	// v²
	vSquared := velocity.Multiply(velocity)

	// m × v²
	result := mass.Value.Multiply(vSquared).Scale(0.5)

	// Should have dimensions of energy [L²M¹T⁻²]
	expectedDim := Dimension{L: 2, M: 1, T: -2}
	if result.Dim() != expectedDim {
		t.Errorf("Kinetic energy dimension = %v, want %v", result.Dim(), expectedDim)
	}

	if !almostEqual(result.Val(), 9.0, 1e-10) {
		t.Errorf("Kinetic energy = %v J, want 9.0 J", result.Val())
	}
}

func TestComplexCalculation_UniversalGravitation(t *testing.T) {
	// F = G × m1 × m2 / r²
	// G = 6.674e-11 N⋅m²/kg² = [L³M⁻¹T⁻²]
	// m1 = m2 = 1 kg
	// r = 1 m
	// F should have dimension [LMT⁻²] (force)

	G := NewValue(6.674e-11, Dimension{L: 3, M: -1, T: -2})
	m1 := Kilogram(1.0)
	m2 := Kilogram(1.0)
	r := Meter(1.0)

	// G × m1 × m2
	numerator := G.Multiply(m1.Value).Multiply(m2.Value)

	// r²
	rSquared := r.Value.Power(2)

	// F = numerator / r²
	force := numerator.Divide(rSquared)

	// Should have dimensions of force [LMT⁻²]
	expectedDim := Dimension{L: 1, M: 1, T: -2}
	if force.Dim() != expectedDim {
		t.Errorf("Force dimension = %v, want %v", force.Dim(), expectedDim)
	}

	if !almostEqual(force.Val(), 6.674e-11, 1e-20) {
		t.Errorf("Force = %v N, want 6.674e-11 N", force.Val())
	}
}

// -----------------------------------------------------------------------------
// Helper Functions Tests
// -----------------------------------------------------------------------------

func TestAlmostEqual(t *testing.T) {
	tests := []struct {
		name      string
		a, b, tol float64
		want      bool
	}{
		{"exact equal", 1.0, 1.0, 1e-10, true},
		{"within tolerance", 1.0, 1.0000000001, 1e-8, true},
		{"outside tolerance", 1.0, 1.001, 1e-10, false},
		{"zero comparison", 0.0, 1e-15, 1e-10, true},
		{"negative values", -1.0, -1.0000000001, 1e-8, true},
		{"large numbers", 1e10, 1e10 + 1e-5, 1e-10, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := almostEqual(tt.a, tt.b, tt.tol)
			if got != tt.want {
				t.Errorf("almostEqual(%v, %v, %v) = %v, want %v",
					tt.a, tt.b, tt.tol, got, tt.want)
			}
		})
	}
}

// -----------------------------------------------------------------------------
// Benchmark Tests
// -----------------------------------------------------------------------------

func BenchmarkValueAdd(b *testing.B) {
	v1 := Meter(5.0).Value
	v2 := Meter(3.0).Value
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v1.Add(v2)
	}
}

func BenchmarkValueMultiply(b *testing.B) {
	v1 := Meter(5.0).Value
	v2 := Meter(3.0).Value
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v1.Multiply(v2)
	}
}

func BenchmarkValueDivide(b *testing.B) {
	v1 := Meter(10.0).Value
	v2 := Second(2.0).Value
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v1.Divide(v2)
	}
}

func BenchmarkValuePower(b *testing.B) {
	v := Meter(5.0).Value
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.Power(2)
	}
}

func BenchmarkMeterCreation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Meter(float64(i))
	}
}
