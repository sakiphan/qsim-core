package constants

import "github.com/sakiphan/qsim-core/units"

// Universal Constants
// All values from CODATA 2018 recommended values unless otherwise noted.

// SpeedOfLight is the speed of light in vacuum (c).
// Value: 299,792,458 m/s (exact by definition)
//
// References:
//   - CODATA 2018
var SpeedOfLight = units.MeterPerSecond(299792458.0)

// PlanckConstant is Planck's constant (h).
// Value: 6.62607015 × 10⁻³⁴ J⋅s (exact by definition)
//
// The quantum of action, relating energy and frequency: E = hν
//
// References:
//   - CODATA 2018
var PlanckConstant = units.NewValue(6.62607015e-34, units.Dimension{L: 2, M: 1, T: -1})

// PlanckReduced is the reduced Planck constant (ℏ = h/2π).
// Value: 1.054571817... × 10⁻³⁴ J⋅s
//
// Commonly used in quantum mechanics: [x, p] = iℏ
//
// References:
//   - CODATA 2018
var PlanckReduced = units.NewValue(1.054571817e-34, units.Dimension{L: 2, M: 1, T: -1})

// GravitationalConstant is Newton's gravitational constant (G).
// Value: 6.67430(15) × 10⁻¹¹ m³/(kg⋅s²)
// Relative standard uncertainty: 2.2 × 10⁻⁵
//
// Used in universal gravitation: F = G(m₁m₂)/r²
//
// References:
//   - CODATA 2018
var GravitationalConstant = units.NewValue(6.67430e-11, units.Dimension{L: 3, M: -1, T: -2})

// BoltzmannConstant is the Boltzmann constant (k_B).
// Value: 1.380649 × 10⁻²³ J/K (exact by definition)
//
// Relates temperature to energy: E = k_B T
//
// References:
//   - CODATA 2018
var BoltzmannConstant = units.NewValue(1.380649e-23, units.Dimension{L: 2, M: 1, T: -2, Θ: -1})

// AvogadroConstant is Avogadro's number (N_A).
// Value: 6.02214076 × 10²³ mol⁻¹ (exact by definition)
//
// Number of constituent particles per mole.
//
// References:
//   - CODATA 2018
var AvogadroConstant = units.NewValue(6.02214076e23, units.Dimension{N: -1})

// UniversalGasConstant is the molar gas constant (R = N_A k_B).
// Value: 8.314462618... J/(mol⋅K)
//
// Used in ideal gas law: PV = nRT
//
// References:
//   - CODATA 2018
var UniversalGasConstant = units.NewValue(8.314462618, units.Dimension{L: 2, M: 1, T: -2, Θ: -1, N: -1})

// VacuumPermittivity is the electric constant (ε₀).
// Value: 8.8541878128(13) × 10⁻¹² F/m
//
// Used in Coulomb's law: F = (1/4πε₀)(q₁q₂/r²)
//
// References:
//   - CODATA 2018
var VacuumPermittivity = units.NewValue(8.8541878128e-12, units.Dimension{L: -3, M: -1, T: 4, I: 2})

// VacuumPermeability is the magnetic constant (μ₀).
// Value: 1.25663706212(19) × 10⁻⁶ H/m
//
// Related to ε₀ and c: μ₀ε₀c² = 1
//
// References:
//   - CODATA 2018
var VacuumPermeability = units.NewValue(1.25663706212e-6, units.Dimension{L: 1, M: 1, T: -2, I: -2})

// ElementaryCharge is the elementary charge (e).
// Value: 1.602176634 × 10⁻¹⁹ C (exact by definition)
//
// Charge of a proton, magnitude of electron charge.
//
// References:
//   - CODATA 2018
var ElementaryCharge = units.Coulomb(1.602176634e-19)

// CoulombConstant is Coulomb's constant (k_e = 1/4πε₀).
// Value: 8.9875517923(14) × 10⁹ N⋅m²/C²
//
// Used in Coulomb's law: F = k_e(q₁q₂/r²)
//
// References:
//   - CODATA 2018
var CoulombConstant = units.NewValue(8.9875517923e9, units.Dimension{L: 3, M: 1, T: -4, I: -2})

// StefanBoltzmannConstant is the Stefan-Boltzmann constant (σ).
// Value: 5.670374419... × 10⁻⁸ W/(m²⋅K⁴)
//
// Used in blackbody radiation: P = σAT⁴
//
// References:
//   - CODATA 2018
var StefanBoltzmannConstant = units.NewValue(5.670374419e-8, units.Dimension{M: 1, T: -3, Θ: -4})

// WienDisplacementConstant is Wien's displacement law constant (b).
// Value: 2.897771955... × 10⁻³ m⋅K
//
// Relates peak wavelength to temperature: λ_max = b/T
//
// References:
//   - CODATA 2018
var WienDisplacementConstant = units.NewValue(2.897771955e-3, units.Dimension{L: 1, Θ: 1})

// RydbergConstant is the Rydberg constant (R_∞).
// Value: 10,973,731.568160(21) m⁻¹
// Relative standard uncertainty: 1.9 × 10⁻¹²
//
// Used in hydrogen spectral lines: 1/λ = R_∞(1/n₁² - 1/n₂²)
//
// References:
//   - CODATA 2018
var RydbergConstant = units.NewValue(10973731.568160, units.Dimension{L: -1})

// FineStructureConstant is the fine-structure constant (α ≈ 1/137).
// Value: 7.2973525693(11) × 10⁻³ (dimensionless)
// Relative standard uncertainty: 1.5 × 10⁻¹⁰
//
// Characterizes strength of electromagnetic interaction.
// α = e²/(4πε₀ℏc)
//
// References:
//   - CODATA 2018
var FineStructureConstant = units.Dimensionless(7.2973525693e-3)

// BohrRadius is the Bohr radius (a₀).
// Value: 5.29177210903(80) × 10⁻¹¹ m
// Relative standard uncertainty: 1.5 × 10⁻¹⁰
//
// Radius of hydrogen atom ground state: a₀ = 4πε₀ℏ²/(m_e e²)
//
// References:
//   - CODATA 2018
var BohrRadius = units.Meter(5.29177210903e-11)

// BohrMagneton is the Bohr magneton (μ_B).
// Value: 9.2740100783(28) × 10⁻²⁴ J/T
// Relative standard uncertainty: 3.0 × 10⁻¹⁰
//
// Magnetic moment of electron due to orbital or spin angular momentum.
//
// References:
//   - CODATA 2018
var BohrMagneton = units.NewValue(9.2740100783e-24, units.Dimension{L: 2, I: 1})

// StandardGravity is standard acceleration due to gravity on Earth (g).
// Value: 9.80665 m/s² (exact by definition)
//
// Standard value for gravitational acceleration at Earth's surface.
//
// References:
//   - ISO 80000-3:2006
var StandardGravity = units.MeterPerSecond2(9.80665)

// AtomicMassUnit is the unified atomic mass unit (u or Da).
// Value: 1.66053906660(50) × 10⁻²⁷ kg
// Relative standard uncertainty: 3.0 × 10⁻¹⁰
//
// Defined as 1/12 of the mass of carbon-12 atom.
//
// References:
//   - CODATA 2018
var AtomicMassUnit = units.Kilogram(1.66053906660e-27)

// -----------------------------------------------------------------------------
// Astronomical and Cosmological Constants
// -----------------------------------------------------------------------------

// AstronomicalUnit is the astronomical unit (AU).
// Value: 149,597,870,700 m (exact by definition)
//
// Approximate mean distance from Earth to Sun.
//
// References:
//   - IAU 2012 Resolution B2
var AstronomicalUnit = units.Meter(1.495978707e11)

// Parsec is one parsec (pc).
// Value: 3.0856775814913673 × 10¹⁶ m
//
// Distance at which 1 AU subtends 1 arcsecond.
// 1 pc ≈ 3.26 light-years
//
// References:
//   - IAU definition
var Parsec = units.Meter(3.0856775814913673e16)

// LightYear is one light-year (ly).
// Value: 9.4607304725808 × 10¹⁵ m
//
// Distance light travels in one Julian year (365.25 days).
//
// References:
//   - IAU definition
var LightYear = units.Meter(9.4607304725808e15)

// SolarMass is the mass of the Sun (M☉).
// Value: 1.98892 × 10³⁰ kg
//
// Nominal value defined by IAU 2015 Resolution B3.
//
// References:
//   - IAU 2015 Resolution B3
var SolarMass = units.Kilogram(1.98892e30)

// EarthMass is the mass of Earth (M⊕).
// Value: 5.9722 × 10²⁴ kg
//
// References:
//   - NASA JPL planetary fact sheet
var EarthMass = units.Kilogram(5.9722e24)

// SolarLuminosity is the luminosity of the Sun (L☉).
// Value: 3.828 × 10²⁶ W
//
// Nominal value defined by IAU 2015 Resolution B3.
//
// References:
//   - IAU 2015 Resolution B3
var SolarLuminosity = units.Watt(3.828e26)

// SolarRadius is the radius of the Sun (R☉).
// Value: 6.957 × 10⁸ m
//
// Nominal value defined by IAU 2015 Resolution B3.
//
// References:
//   - IAU 2015 Resolution B3
var SolarRadius = units.Meter(6.957e8)

// EarthRadius is the mean radius of Earth (R⊕).
// Value: 6.371 × 10⁶ m
//
// References:
//   - NASA Earth Fact Sheet
var EarthRadius = units.Meter(6.371e6)

// HubbleConstant is the Hubble constant (H₀).
// Value: 67.4 km/(s⋅Mpc) (Planck 2018)
// Uncertainty: ±0.5 km/(s⋅Mpc)
//
// Expansion rate of the universe. Different measurements give slightly
// different values (Hubble tension).
//
// Converted to SI: 67.4 km/s/Mpc = 2.18e-18 s⁻¹
//
// References:
//   - Planck Collaboration 2018 (Planck 2018 results. VI. Cosmological parameters)
var HubbleConstant = units.Hertz(2.18e-18)

// HubbleTime is the Hubble time (1/H₀).
// Value: ≈ 14.5 × 10⁹ years
//
// Approximate age scale of the universe.
//
// References:
//   - Derived from Planck 2018 H₀
var HubbleTime = units.Year(14.5e9)

// CriticalDensity is the critical density of the universe (ρ_c).
// Value: 3H₀²/(8πG) ≈ 9.47 × 10⁻²⁷ kg/m³
//
// Density required for a flat universe.
//
// References:
//   - Derived from Planck 2018 H₀ and CODATA G
var CriticalDensity = units.NewValue(9.47e-27, units.Dimension{L: -3, M: 1})

// CMBTemperature is the cosmic microwave background temperature (T_CMB).
// Value: 2.7255(6) K
//
// Temperature of the CMB radiation today.
//
// References:
//   - Fixsen 2009, ApJ 707, 916
var CMBTemperature = units.Kelvin(2.7255)

// -----------------------------------------------------------------------------
// Conversion Factors
// -----------------------------------------------------------------------------

// ElectronVoltToJoule is the conversion factor from electron volts to joules.
// Value: 1 eV = 1.602176634 × 10⁻¹⁹ J
var ElectronVoltToJoule = 1.602176634e-19

// JouleToElectronVolt is the conversion factor from joules to electron volts.
// Value: 1 J = 6.241509074... × 10¹⁸ eV
var JouleToElectronVolt = 6.241509074e18

// PlanckLength is the Planck length (l_P = √(ℏG/c³)).
// Value: 1.616255(18) × 10⁻³⁵ m
//
// Natural unit of length in quantum gravity.
//
// References:
//   - CODATA 2018
var PlanckLength = units.Meter(1.616255e-35)

// PlanckMass is the Planck mass (m_P = √(ℏc/G)).
// Value: 2.176434(24) × 10⁻⁸ kg
//
// Natural unit of mass in quantum gravity.
//
// References:
//   - CODATA 2018
var PlanckMass = units.Kilogram(2.176434e-8)

// PlanckTime is the Planck time (t_P = √(ℏG/c⁵)).
// Value: 5.391247(60) × 10⁻⁴⁴ s
//
// Natural unit of time in quantum gravity.
//
// References:
//   - CODATA 2018
var PlanckTime = units.Second(5.391247e-44)

// PlanckTemperature is the Planck temperature (T_P = √(ℏc⁵/(Gk_B²))).
// Value: 1.416784(16) × 10³² K
//
// Natural unit of temperature in quantum gravity.
//
// References:
//   - CODATA 2018
var PlanckTemperature = units.Kelvin(1.416784e32)
