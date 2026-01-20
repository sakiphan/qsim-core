package units

// This file defines the seven SI base units and their common multiples.
//
// SI Base Units:
//   - Length: meter (m)
//   - Mass: kilogram (kg)
//   - Time: second (s)
//   - Electric Current: ampere (A)
//   - Temperature: kelvin (K)
//   - Amount of Substance: mole (mol)
//   - Luminous Intensity: candela (cd)
//
// References:
//   - BIPM, "The International System of Units (SI)", 9th edition, 2019

// -----------------------------------------------------------------------------
// Length [L]
// -----------------------------------------------------------------------------

// Length represents a physical length with dimension [L¹].
type Length struct{ Value }

// Meter creates a Length value in meters (SI base unit).
//
// Example:
//
//	length := units.Meter(5.0) // 5 meters
func Meter(value float64) Length {
	return Length{NewValue(value, Dimension{L: 1})}
}

// Millimeter creates a Length value in millimeters (10⁻³ m).
func Millimeter(value float64) Length {
	return Meter(value * 1e-3)
}

// Centimeter creates a Length value in centimeters (10⁻² m).
func Centimeter(value float64) Length {
	return Meter(value * 1e-2)
}

// Kilometer creates a Length value in kilometers (10³ m).
func Kilometer(value float64) Length {
	return Meter(value * 1e3)
}

// Micrometer creates a Length value in micrometers (10⁻⁶ m).
func Micrometer(value float64) Length {
	return Meter(value * 1e-6)
}

// Nanometer creates a Length value in nanometers (10⁻⁹ m).
func Nanometer(value float64) Length {
	return Meter(value * 1e-9)
}

// Angstrom creates a Length value in angstroms (10⁻¹⁰ m).
// Commonly used in atomic physics and chemistry.
func Angstrom(value float64) Length {
	return Meter(value * 1e-10)
}

// Inch creates a Length value in inches (1 in = 0.0254 m).
func Inch(value float64) Length {
	return Meter(value * 0.0254)
}

// Foot creates a Length value in feet (1 ft = 0.3048 m).
func Foot(value float64) Length {
	return Meter(value * 0.3048)
}

// Mile creates a Length value in miles (1 mi = 1609.344 m).
func Mile(value float64) Length {
	return Meter(value * 1609.344)
}

// AstronomicalUnit creates a Length value in astronomical units (1 AU = 1.495978707e11 m).
// One AU is approximately the mean distance from Earth to the Sun.
func AstronomicalUnit(value float64) Length {
	return Meter(value * 1.495978707e11)
}

// LightYear creates a Length value in light-years (1 ly = 9.4607304725808e15 m).
// One light-year is the distance light travels in one Julian year.
func LightYear(value float64) Length {
	return Meter(value * 9.4607304725808e15)
}

// Parsec creates a Length value in parsecs (1 pc = 3.0856775814913673e16 m).
// One parsec is approximately 3.26 light-years.
func Parsec(value float64) Length {
	return Meter(value * 3.0856775814913673e16)
}

// -----------------------------------------------------------------------------
// Mass [M]
// -----------------------------------------------------------------------------

// Mass represents a physical mass with dimension [M¹].
type Mass struct{ Value }

// Kilogram creates a Mass value in kilograms (SI base unit).
//
// Example:
//
//	mass := units.Kilogram(2.0) // 2 kilograms
func Kilogram(value float64) Mass {
	return Mass{NewValue(value, Dimension{M: 1})}
}

// Gram creates a Mass value in grams (10⁻³ kg).
func Gram(value float64) Mass {
	return Kilogram(value * 1e-3)
}

// Milligram creates a Mass value in milligrams (10⁻⁶ kg).
func Milligram(value float64) Mass {
	return Kilogram(value * 1e-6)
}

// Microgram creates a Mass value in micrograms (10⁻⁹ kg).
func Microgram(value float64) Mass {
	return Kilogram(value * 1e-9)
}

// Tonne creates a Mass value in metric tonnes (10³ kg).
func Tonne(value float64) Mass {
	return Kilogram(value * 1e3)
}

// Pound creates a Mass value in pounds (1 lb = 0.45359237 kg).
func Pound(value float64) Mass {
	return Kilogram(value * 0.45359237)
}

// Ounce creates a Mass value in ounces (1 oz = 0.028349523125 kg).
func Ounce(value float64) Mass {
	return Kilogram(value * 0.028349523125)
}

// AtomicMassUnit creates a Mass value in atomic mass units (1 u = 1.66053906660e-27 kg).
// Also known as dalton (Da).
func AtomicMassUnit(value float64) Mass {
	return Kilogram(value * 1.66053906660e-27)
}

// SolarMass creates a Mass value in solar masses (1 M☉ = 1.98892e30 kg).
// One solar mass is the mass of the Sun.
func SolarMass(value float64) Mass {
	return Kilogram(value * 1.98892e30)
}

// EarthMass creates a Mass value in Earth masses (1 M⊕ = 5.9722e24 kg).
func EarthMass(value float64) Mass {
	return Kilogram(value * 5.9722e24)
}

// -----------------------------------------------------------------------------
// Time [T]
// -----------------------------------------------------------------------------

// Time represents a physical time interval with dimension [T¹].
type Time struct{ Value }

// Second creates a Time value in seconds (SI base unit).
//
// Example:
//
//	duration := units.Second(10.0) // 10 seconds
func Second(value float64) Time {
	return Time{NewValue(value, Dimension{T: 1})}
}

// Millisecond creates a Time value in milliseconds (10⁻³ s).
func Millisecond(value float64) Time {
	return Second(value * 1e-3)
}

// Microsecond creates a Time value in microseconds (10⁻⁶ s).
func Microsecond(value float64) Time {
	return Second(value * 1e-6)
}

// Nanosecond creates a Time value in nanoseconds (10⁻⁹ s).
func Nanosecond(value float64) Time {
	return Second(value * 1e-9)
}

// Minute creates a Time value in minutes (1 min = 60 s).
func Minute(value float64) Time {
	return Second(value * 60)
}

// Hour creates a Time value in hours (1 h = 3600 s).
func Hour(value float64) Time {
	return Second(value * 3600)
}

// Day creates a Time value in days (1 d = 86400 s).
func Day(value float64) Time {
	return Second(value * 86400)
}

// Week creates a Time value in weeks (1 week = 604800 s).
func Week(value float64) Time {
	return Second(value * 604800)
}

// Year creates a Time value in Julian years (1 yr = 31557600 s).
// One Julian year is exactly 365.25 days.
func Year(value float64) Time {
	return Second(value * 31557600)
}

// -----------------------------------------------------------------------------
// Electric Current [I]
// -----------------------------------------------------------------------------

// Current represents an electric current with dimension [I¹].
type Current struct{ Value }

// Ampere creates a Current value in amperes (SI base unit).
//
// Example:
//
//	current := units.Ampere(2.5) // 2.5 amperes
func Ampere(value float64) Current {
	return Current{NewValue(value, Dimension{I: 1})}
}

// Milliampere creates a Current value in milliamperes (10⁻³ A).
func Milliampere(value float64) Current {
	return Ampere(value * 1e-3)
}

// Microampere creates a Current value in microamperes (10⁻⁶ A).
func Microampere(value float64) Current {
	return Ampere(value * 1e-6)
}

// Kiloampere creates a Current value in kiloamperes (10³ A).
func Kiloampere(value float64) Current {
	return Ampere(value * 1e3)
}

// -----------------------------------------------------------------------------
// Thermodynamic Temperature [Θ]
// -----------------------------------------------------------------------------

// Temperature represents a thermodynamic temperature with dimension [Θ¹].
type Temperature struct{ Value }

// Kelvin creates a Temperature value in kelvins (SI base unit).
//
// Example:
//
//	temp := units.Kelvin(273.15) // 273.15 K (0°C)
func Kelvin(value float64) Temperature {
	return Temperature{NewValue(value, Dimension{Θ: 1})}
}

// Celsius creates a Temperature value from degrees Celsius.
// Converts to kelvin: K = °C + 273.15
func Celsius(value float64) Temperature {
	return Kelvin(value + 273.15)
}

// Fahrenheit creates a Temperature value from degrees Fahrenheit.
// Converts to kelvin: K = (°F + 459.67) × 5/9
func Fahrenheit(value float64) Temperature {
	return Kelvin((value + 459.67) * 5.0 / 9.0)
}

// -----------------------------------------------------------------------------
// Amount of Substance [N]
// -----------------------------------------------------------------------------

// Amount represents an amount of substance with dimension [N¹].
type Amount struct{ Value }

// Mole creates an Amount value in moles (SI base unit).
//
// Example:
//
//	amount := units.Mole(2.0) // 2 moles
func Mole(value float64) Amount {
	return Amount{NewValue(value, Dimension{N: 1})}
}

// Millimole creates an Amount value in millimoles (10⁻³ mol).
func Millimole(value float64) Amount {
	return Mole(value * 1e-3)
}

// Micromole creates an Amount value in micromoles (10⁻⁶ mol).
func Micromole(value float64) Amount {
	return Mole(value * 1e-6)
}

// Kilomole creates an Amount value in kilomoles (10³ mol).
func Kilomole(value float64) Amount {
	return Mole(value * 1e3)
}

// -----------------------------------------------------------------------------
// Luminous Intensity [J]
// -----------------------------------------------------------------------------

// LuminousIntensity represents a luminous intensity with dimension [J¹].
type LuminousIntensity struct{ Value }

// Candela creates a LuminousIntensity value in candelas (SI base unit).
//
// Example:
//
//	intensity := units.Candela(100.0) // 100 candelas
func Candela(value float64) LuminousIntensity {
	return LuminousIntensity{NewValue(value, Dimension{J: 1})}
}

// Millicandela creates a LuminousIntensity value in millicandelas (10⁻³ cd).
func Millicandela(value float64) LuminousIntensity {
	return Candela(value * 1e-3)
}

// Kilocandela creates a LuminousIntensity value in kilocandelas (10³ cd).
func Kilocandela(value float64) LuminousIntensity {
	return Candela(value * 1e3)
}
