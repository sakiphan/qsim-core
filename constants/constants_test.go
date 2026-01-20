package constants

import (
	"math"
	"testing"

	"github.com/sakiphan/qsim-core/units"
)

// Helper function for approximate equality
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

// -----------------------------------------------------------------------------
// Universal Constants Tests
// -----------------------------------------------------------------------------

func TestSpeedOfLight(t *testing.T) {
	// Speed of light is exact by definition
	expected := 299792458.0
	if SpeedOfLight.Val() != expected {
		t.Errorf("SpeedOfLight = %v m/s, want %v m/s", SpeedOfLight.Val(), expected)
	}
}

func TestPlanckConstants(t *testing.T) {
	// Test Planck constant
	if !almostEqual(PlanckConstant.Val(), 6.62607015e-34, 1e-42) {
		t.Errorf("PlanckConstant = %e, want 6.62607015e-34", PlanckConstant.Val())
	}

	// Test reduced Planck constant: ℏ = h/(2π)
	expectedHbar := PlanckConstant.Val() / (2.0 * math.Pi)
	if !almostEqual(PlanckReduced.Val(), expectedHbar, 1e-42) {
		t.Errorf("PlanckReduced = %e, want %e", PlanckReduced.Val(), expectedHbar)
	}
}

func TestGravitationalConstant(t *testing.T) {
	expected := 6.67430e-11
	if !almostEqual(GravitationalConstant.Val(), expected, 1e-16) {
		t.Errorf("GravitationalConstant = %e, want %e", GravitationalConstant.Val(), expected)
	}
}

func TestBoltzmannConstant(t *testing.T) {
	// Boltzmann constant is exact by definition
	expected := 1.380649e-23
	if !almostEqual(BoltzmannConstant.Val(), expected, 1e-31) {
		t.Errorf("BoltzmannConstant = %e, want %e", BoltzmannConstant.Val(), expected)
	}
}

func TestAvogadroConstant(t *testing.T) {
	// Avogadro constant is exact by definition
	expected := 6.02214076e23
	if !almostEqual(AvogadroConstant.Val(), expected, 1e15) {
		t.Errorf("AvogadroConstant = %e, want %e", AvogadroConstant.Val(), expected)
	}
}

func TestUniversalGasConstant(t *testing.T) {
	// R = N_A × k_B
	expectedR := AvogadroConstant.Val() * BoltzmannConstant.Val()
	if !almostEqual(UniversalGasConstant.Val(), expectedR, 1e-8) {
		t.Errorf("UniversalGasConstant = %v, expected N_A × k_B = %v", UniversalGasConstant.Val(), expectedR)
	}
}

func TestVacuumConstants(t *testing.T) {
	// Test vacuum permittivity
	expectedEpsilon0 := 8.8541878128e-12
	if !almostEqual(VacuumPermittivity.Val(), expectedEpsilon0, 1e-22) {
		t.Errorf("VacuumPermittivity = %e, want %e", VacuumPermittivity.Val(), expectedEpsilon0)
	}

	// Test vacuum permeability
	expectedMu0 := 1.25663706212e-6
	if !almostEqual(VacuumPermeability.Val(), expectedMu0, 1e-17) {
		t.Errorf("VacuumPermeability = %e, want %e", VacuumPermeability.Val(), expectedMu0)
	}

	// Test relationship: c² = 1/(μ₀ε₀)
	c := SpeedOfLight.Val()
	mu0 := VacuumPermeability.Val()
	eps0 := VacuumPermittivity.Val()
	cSquaredFromConstants := 1.0 / (mu0 * eps0)

	if !almostEqual(c*c, cSquaredFromConstants, 1e-2) {
		t.Errorf("c² = %e, 1/(μ₀ε₀) = %e, should be equal", c*c, cSquaredFromConstants)
	}
}

func TestElementaryCharge(t *testing.T) {
	// Elementary charge is exact by definition
	expected := 1.602176634e-19
	if !almostEqual(ElementaryCharge.Val(), expected, 1e-27) {
		t.Errorf("ElementaryCharge = %e, want %e", ElementaryCharge.Val(), expected)
	}
}

func TestCoulombConstant(t *testing.T) {
	// k_e = 1/(4πε₀)
	expectedKe := 1.0 / (4.0 * math.Pi * VacuumPermittivity.Val())
	if !almostEqual(CoulombConstant.Val(), expectedKe, 1e-2) {
		t.Errorf("CoulombConstant = %e, 1/(4πε₀) = %e", CoulombConstant.Val(), expectedKe)
	}
}

func TestFineStructureConstant(t *testing.T) {
	// α ≈ 1/137
	alpha := FineStructureConstant.Val()
	inverseAlpha := 1.0 / alpha

	if !almostEqual(inverseAlpha, 137.036, 0.01) {
		t.Errorf("1/α = %v, want ≈ 137.036", inverseAlpha)
	}

	// Test α = e²/(4πε₀ℏc)
	e := ElementaryCharge.Val()
	eps0 := VacuumPermittivity.Val()
	hbar := PlanckReduced.Val()
	c := SpeedOfLight.Val()

	alphaCalculated := (e * e) / (4.0 * math.Pi * eps0 * hbar * c)

	if !almostEqual(alpha, alphaCalculated, 1e-10) {
		t.Errorf("α = %e, calculated from e²/(4πε₀ℏc) = %e", alpha, alphaCalculated)
	}
}

func TestBohrRadius(t *testing.T) {
	expected := 5.29177210903e-11
	if !almostEqual(BohrRadius.Val(), expected, 1e-20) {
		t.Errorf("BohrRadius = %e, want %e", BohrRadius.Val(), expected)
	}
}

func TestStandardGravity(t *testing.T) {
	// Standard gravity is exact by definition
	expected := 9.80665
	if StandardGravity.Value.Val() != expected {
		t.Errorf("StandardGravity = %v m/s², want %v m/s²", StandardGravity.Value.Val(), expected)
	}
}

// -----------------------------------------------------------------------------
// Astronomical Constants Tests
// -----------------------------------------------------------------------------

func TestAstronomicalUnit(t *testing.T) {
	// AU is exact by definition
	expected := 1.495978707e11
	if !almostEqual(AstronomicalUnit.Val(), expected, 1e3) {
		t.Errorf("AstronomicalUnit = %e m, want %e m", AstronomicalUnit.Val(), expected)
	}
}

func TestParsec(t *testing.T) {
	// 1 parsec ≈ 3.26 light-years
	pcInLightYears := Parsec.Val() / LightYear.Val()
	if !almostEqual(pcInLightYears, 3.26, 0.01) {
		t.Errorf("1 parsec = %v ly, want ≈ 3.26 ly", pcInLightYears)
	}
}

func TestSolarMass(t *testing.T) {
	expected := 1.98892e30
	if !almostEqual(SolarMass.Val(), expected, 1e24) {
		t.Errorf("SolarMass = %e kg, want %e kg", SolarMass.Val(), expected)
	}
}

func TestHubbleConstant(t *testing.T) {
	// H₀ ≈ 67.4 km/s/Mpc = 2.18e-18 s⁻¹
	h0 := HubbleConstant.Val()

	if !almostEqual(h0, 2.18e-18, 1e-20) {
		t.Errorf("HubbleConstant = %e s⁻¹, want ≈ 2.18e-18 s⁻¹", h0)
	}

	// Test Hubble time: t_H = 1/H₀
	hubbleTimeCalculated := 1.0 / h0 // in seconds
	hubbleTimeInYears := hubbleTimeCalculated / (365.25 * 24 * 3600)

	if !almostEqual(hubbleTimeInYears, 14.5e9, 1e9) {
		t.Errorf("Hubble time = %e years, want ≈ 14.5e9 years", hubbleTimeInYears)
	}
}

func TestCMBTemperature(t *testing.T) {
	expected := 2.7255
	if !almostEqual(CMBTemperature.Val(), expected, 0.001) {
		t.Errorf("CMBTemperature = %v K, want %v K", CMBTemperature.Val(), expected)
	}
}

// -----------------------------------------------------------------------------
// Planck Units Tests
// -----------------------------------------------------------------------------

func TestPlanckLength(t *testing.T) {
	// l_P = √(ℏG/c³)
	hbar := PlanckReduced.Val()
	G := GravitationalConstant.Val()
	c := SpeedOfLight.Val()

	lPlanck := math.Sqrt((hbar * G) / (c * c * c))

	if !almostEqual(PlanckLength.Val(), lPlanck, 1e-37) {
		t.Errorf("PlanckLength = %e, calculated √(ℏG/c³) = %e", PlanckLength.Val(), lPlanck)
	}
}

func TestPlanckMass(t *testing.T) {
	// m_P = √(ℏc/G)
	hbar := PlanckReduced.Val()
	c := SpeedOfLight.Val()
	G := GravitationalConstant.Val()

	mPlanck := math.Sqrt((hbar * c) / G)

	if !almostEqual(PlanckMass.Val(), mPlanck, 1e-10) {
		t.Errorf("PlanckMass = %e, calculated √(ℏc/G) = %e", PlanckMass.Val(), mPlanck)
	}
}

func TestPlanckTime(t *testing.T) {
	// t_P = √(ℏG/c⁵)
	hbar := PlanckReduced.Val()
	G := GravitationalConstant.Val()
	c := SpeedOfLight.Val()

	tPlanck := math.Sqrt((hbar * G) / math.Pow(c, 5))

	if !almostEqual(PlanckTime.Val(), tPlanck, 1e-46) {
		t.Errorf("PlanckTime = %e, calculated √(ℏG/c⁵) = %e", PlanckTime.Val(), tPlanck)
	}
}

// -----------------------------------------------------------------------------
// Conversion Factor Tests
// -----------------------------------------------------------------------------

func TestElectronVoltConversion(t *testing.T) {
	// Test that eV ↔ J conversions are inverses
	product := ElectronVoltToJoule * JouleToElectronVolt
	if !almostEqual(product, 1.0, 1e-10) {
		t.Errorf("eV↔J conversion: product = %v, want 1.0", product)
	}

	// Test against elementary charge
	if !almostEqual(ElectronVoltToJoule, ElementaryCharge.Val(), 1e-27) {
		t.Errorf("1 eV = %e J, elementary charge = %e C", ElectronVoltToJoule, ElementaryCharge.Val())
	}
}

// -----------------------------------------------------------------------------
// Dimensional Consistency Tests
// -----------------------------------------------------------------------------

func TestDimensionalConsistency(t *testing.T) {
	tests := []struct {
		name     string
		constant units.Value
		expected units.Dimension
	}{
		{
			name:     "SpeedOfLight [LT⁻¹]",
			constant: SpeedOfLight.Value,
			expected: units.Dimension{L: 1, T: -1},
		},
		{
			name:     "GravitationalConstant [L³M⁻¹T⁻²]",
			constant: GravitationalConstant,
			expected: units.Dimension{L: 3, M: -1, T: -2},
		},
		{
			name:     "BoltzmannConstant [L²MT⁻²Θ⁻¹]",
			constant: BoltzmannConstant,
			expected: units.Dimension{L: 2, M: 1, T: -2, Θ: -1},
		},
		{
			name:     "VacuumPermittivity [L⁻³M⁻¹T⁴I²]",
			constant: VacuumPermittivity,
			expected: units.Dimension{L: -3, M: -1, T: 4, I: 2},
		},
		{
			name:     "ElementaryCharge [IT]",
			constant: ElementaryCharge.Value,
			expected: units.Dimension{I: 1, T: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant.Dim() != tt.expected {
				t.Errorf("%s dimension = %s, want %s", tt.name, tt.constant.Dim(), tt.expected)
			}
		})
	}
}

// -----------------------------------------------------------------------------
// Physical Relationship Tests
// -----------------------------------------------------------------------------

func TestEnergyMassEquivalence(t *testing.T) {
	// E = mc² for electron
	m := ElectronMass
	c := SpeedOfLight
	cSquared := c.Value.Multiply(c.Value)
	energy := m.Value.Multiply(cSquared)

	// Compare with electron rest energy
	if !almostEqual(energy.Val(), ElectronRestEnergy.Val(), 1e-20) {
		t.Errorf("m_e c² = %e J, ElectronRestEnergy = %e J", energy.Val(), ElectronRestEnergy.Val())
	}

	// Convert to MeV
	energyMeV := energy.Val() / (ElectronVoltToJoule * 1e6)
	if !almostEqual(energyMeV, ElectronRestEnergyMeV, 1e-6) {
		t.Errorf("m_e c² = %v MeV, want %v MeV", energyMeV, ElectronRestEnergyMeV)
	}
}

func TestComptonWavelength(t *testing.T) {
	// λ_C = h/(m_e c)
	h := PlanckConstant.Val()
	me := ElectronMass.Val()
	c := SpeedOfLight.Val()

	lambdaC := h / (me * c)

	if !almostEqual(lambdaC, ElectronComptonWavelength.Val(), 1e-20) {
		t.Errorf("Electron Compton wavelength: calculated %e, constant %e",
			lambdaC, ElectronComptonWavelength.Val())
	}
}

func TestBohrRadiusFormula(t *testing.T) {
	// a₀ = 4πε₀ℏ²/(m_e e²)
	eps0 := VacuumPermittivity.Val()
	hbar := PlanckReduced.Val()
	me := ElectronMass.Val()
	e := ElementaryCharge.Val()

	a0 := (4.0 * math.Pi * eps0 * hbar * hbar) / (me * e * e)

	if !almostEqual(a0, BohrRadius.Val(), 1e-18) {
		t.Errorf("Bohr radius: calculated %e, constant %e", a0, BohrRadius.Val())
	}
}

func TestMassRatios(t *testing.T) {
	// Test proton/electron mass ratio
	ratio := ProtonMass.Val() / ElectronMass.Val()
	if !almostEqual(ratio, ProtonElectronMassRatio, 1e-6) {
		t.Errorf("m_p/m_e = %v, want %v", ratio, ProtonElectronMassRatio)
	}

	// Test neutron/proton mass ratio
	ratio = NeutronMass.Val() / ProtonMass.Val()
	if !almostEqual(ratio, NeutronProtonMassRatio, 1e-9) {
		t.Errorf("m_n/m_p = %v, want %v", ratio, NeutronProtonMassRatio)
	}
}

// -----------------------------------------------------------------------------
// Example Usage Tests
// -----------------------------------------------------------------------------

func TestExampleCalculations(t *testing.T) {
	// Test photon energy: E = hν for 5 GHz
	frequency := units.Gigahertz(5.0)
	energy := PlanckConstant.Multiply(frequency.Value)
	expectedEnergy := 6.62607015e-34 * 5e9 // J

	if !almostEqual(energy.Val(), expectedEnergy, 1e-32) {
		t.Errorf("Photon energy E=hν: got %e J, want %e J", energy.Val(), expectedEnergy)
	}

	// Test gravitational force: F = Gm₁m₂/r²
	m1 := units.Kilogram(1.0)
	m2 := units.Kilogram(1.0)
	r := units.Meter(1.0)
	rSquared := r.Value.Power(2)

	F := GravitationalConstant.Multiply(m1.Value).Multiply(m2.Value).Divide(rSquared)
	expectedF := GravitationalConstant.Val()

	if !almostEqual(F.Val(), expectedF, 1e-20) {
		t.Errorf("Gravitational force: got %e N, want %e N", F.Val(), expectedF)
	}
}
