package units

// This file defines derived SI units that are combinations of base units.
// Each derived unit has a specific dimensional formula and physical meaning.
//
// Common derived units include:
//   - Mechanical: Force, Energy, Power, Pressure
//   - Electromagnetic: Charge, Voltage, Resistance, Capacitance, Magnetic Field
//   - Frequency and other special units
//
// References:
//   - BIPM, "The International System of Units (SI)", 9th edition, 2019

// -----------------------------------------------------------------------------
// Geometric Units
// -----------------------------------------------------------------------------

// Area represents a physical area with dimension [L²].
type Area struct{ Value }

// SquareMeter creates an Area value in square meters.
func SquareMeter(value float64) Area {
	return Area{NewValue(value, Dimension{L: 2})}
}

// SquareCentimeter creates an Area value in square centimeters (10⁻⁴ m²).
func SquareCentimeter(value float64) Area {
	return SquareMeter(value * 1e-4)
}

// SquareKilometer creates an Area value in square kilometers (10⁶ m²).
func SquareKilometer(value float64) Area {
	return SquareMeter(value * 1e6)
}

// Hectare creates an Area value in hectares (10⁴ m²).
func Hectare(value float64) Area {
	return SquareMeter(value * 1e4)
}

// Volume represents a physical volume with dimension [L³].
type Volume struct{ Value }

// CubicMeter creates a Volume value in cubic meters.
func CubicMeter(value float64) Volume {
	return Volume{NewValue(value, Dimension{L: 3})}
}

// Liter creates a Volume value in liters (10⁻³ m³).
func Liter(value float64) Volume {
	return CubicMeter(value * 1e-3)
}

// Milliliter creates a Volume value in milliliters (10⁻⁶ m³).
func Milliliter(value float64) Volume {
	return CubicMeter(value * 1e-6)
}

// CubicCentimeter creates a Volume value in cubic centimeters (10⁻⁶ m³).
func CubicCentimeter(value float64) Volume {
	return CubicMeter(value * 1e-6)
}

// -----------------------------------------------------------------------------
// Kinematic Units
// -----------------------------------------------------------------------------

// Velocity represents a velocity with dimension [LT⁻¹].
type Velocity struct{ Value }

// MeterPerSecond creates a Velocity value in meters per second.
func MeterPerSecond(value float64) Velocity {
	return Velocity{NewValue(value, Dimension{L: 1, T: -1})}
}

// KilometerPerHour creates a Velocity value in kilometers per hour.
func KilometerPerHour(value float64) Velocity {
	return MeterPerSecond(value / 3.6)
}

// MilePerHour creates a Velocity value in miles per hour.
func MilePerHour(value float64) Velocity {
	return MeterPerSecond(value * 0.44704)
}

// SpeedOfLight creates a Velocity value as a multiple of the speed of light.
// c = 299,792,458 m/s
func SpeedOfLight(value float64) Velocity {
	return MeterPerSecond(value * 299792458.0)
}

// Acceleration represents an acceleration with dimension [LT⁻²].
type Acceleration struct{ Value }

// MeterPerSecond2 creates an Acceleration value in meters per second squared.
func MeterPerSecond2(value float64) Acceleration {
	return Acceleration{NewValue(value, Dimension{L: 1, T: -2})}
}

// StandardGravity creates an Acceleration value in standard gravity units (9.80665 m/s²).
func StandardGravity(value float64) Acceleration {
	return MeterPerSecond2(value * 9.80665)
}

// -----------------------------------------------------------------------------
// Mechanical Units
// -----------------------------------------------------------------------------

// Force represents a force with dimension [LMT⁻²].
type Force struct{ Value }

// Newton creates a Force value in newtons (kg⋅m/s²).
func Newton(value float64) Force {
	return Force{NewValue(value, Dimension{L: 1, M: 1, T: -2})}
}

// Kilonewton creates a Force value in kilonewtons (10³ N).
func Kilonewton(value float64) Force {
	return Newton(value * 1e3)
}

// Dyne creates a Force value in dynes (10⁻⁵ N).
// The dyne is the CGS unit of force.
func Dyne(value float64) Force {
	return Newton(value * 1e-5)
}

// PoundForce creates a Force value in pounds-force (4.4482216152605 N).
func PoundForce(value float64) Force {
	return Newton(value * 4.4482216152605)
}

// Energy represents an energy with dimension [L²MT⁻²].
type Energy struct{ Value }

// Joule creates an Energy value in joules (kg⋅m²/s²).
func Joule(value float64) Energy {
	return Energy{NewValue(value, Dimension{L: 2, M: 1, T: -2})}
}

// Kilojoule creates an Energy value in kilojoules (10³ J).
func Kilojoule(value float64) Energy {
	return Joule(value * 1e3)
}

// Megajoule creates an Energy value in megajoules (10⁶ J).
func Megajoule(value float64) Energy {
	return Joule(value * 1e6)
}

// Calorie creates an Energy value in thermochemical calories (4.184 J).
func Calorie(value float64) Energy {
	return Joule(value * 4.184)
}

// Kilocalorie creates an Energy value in kilocalories (4184 J).
func Kilocalorie(value float64) Energy {
	return Joule(value * 4184.0)
}

// ElectronVolt creates an Energy value in electron volts (1.602176634e-19 J).
// Commonly used in atomic and particle physics.
func ElectronVolt(value float64) Energy {
	return Joule(value * 1.602176634e-19)
}

// KiloelectronVolt creates an Energy value in kiloelectron volts (keV).
func KiloelectronVolt(value float64) Energy {
	return ElectronVolt(value * 1e3)
}

// MegaelectronVolt creates an Energy value in megaelectron volts (MeV).
func MegaelectronVolt(value float64) Energy {
	return ElectronVolt(value * 1e6)
}

// GigaelectronVolt creates an Energy value in gigaelectron volts (GeV).
func GigaelectronVolt(value float64) Energy {
	return ElectronVolt(value * 1e9)
}

// Power represents a power (energy per time) with dimension [L²MT⁻³].
type Power struct{ Value }

// Watt creates a Power value in watts (J/s = kg⋅m²/s³).
func Watt(value float64) Power {
	return Power{NewValue(value, Dimension{L: 2, M: 1, T: -3})}
}

// Kilowatt creates a Power value in kilowatts (10³ W).
func Kilowatt(value float64) Power {
	return Watt(value * 1e3)
}

// Megawatt creates a Power value in megawatts (10⁶ W).
func Megawatt(value float64) Power {
	return Watt(value * 1e6)
}

// Gigawatt creates a Power value in gigawatts (10⁹ W).
func Gigawatt(value float64) Power {
	return Watt(value * 1e9)
}

// Horsepower creates a Power value in mechanical horsepower (745.69987158227022 W).
func Horsepower(value float64) Power {
	return Watt(value * 745.69987158227022)
}

// Pressure represents a pressure (force per area) with dimension [L⁻¹MT⁻²].
type Pressure struct{ Value }

// Pascal creates a Pressure value in pascals (N/m² = kg/(m⋅s²)).
func Pascal(value float64) Pressure {
	return Pressure{NewValue(value, Dimension{L: -1, M: 1, T: -2})}
}

// Kilopascal creates a Pressure value in kilopascals (10³ Pa).
func Kilopascal(value float64) Pressure {
	return Pascal(value * 1e3)
}

// Megapascal creates a Pressure value in megapascals (10⁶ Pa).
func Megapascal(value float64) Pressure {
	return Pascal(value * 1e6)
}

// Bar creates a Pressure value in bars (10⁵ Pa).
func Bar(value float64) Pressure {
	return Pascal(value * 1e5)
}

// Atmosphere creates a Pressure value in standard atmospheres (101325 Pa).
func Atmosphere(value float64) Pressure {
	return Pascal(value * 101325.0)
}

// Torr creates a Pressure value in torr (133.322368421 Pa).
// 1 torr = 1/760 atm
func Torr(value float64) Pressure {
	return Pascal(value * 133.322368421)
}

// PSI creates a Pressure value in pounds per square inch (6894.757293168 Pa).
func PSI(value float64) Pressure {
	return Pascal(value * 6894.757293168)
}

// -----------------------------------------------------------------------------
// Frequency and Angular Units
// -----------------------------------------------------------------------------

// Frequency represents a frequency with dimension [T⁻¹].
type Frequency struct{ Value }

// Hertz creates a Frequency value in hertz (s⁻¹).
func Hertz(value float64) Frequency {
	return Frequency{NewValue(value, Dimension{T: -1})}
}

// Kilohertz creates a Frequency value in kilohertz (10³ Hz).
func Kilohertz(value float64) Frequency {
	return Hertz(value * 1e3)
}

// Megahertz creates a Frequency value in megahertz (10⁶ Hz).
func Megahertz(value float64) Frequency {
	return Hertz(value * 1e6)
}

// Gigahertz creates a Frequency value in gigahertz (10⁹ Hz).
func Gigahertz(value float64) Frequency {
	return Hertz(value * 1e9)
}

// AngularVelocity represents an angular velocity with dimension [T⁻¹].
// Note: Radians are dimensionless, so angular velocity has the same dimension as frequency.
type AngularVelocity struct{ Value }

// RadianPerSecond creates an AngularVelocity value in radians per second.
func RadianPerSecond(value float64) AngularVelocity {
	return AngularVelocity{NewValue(value, Dimension{T: -1})}
}

// RPM creates an AngularVelocity value from revolutions per minute.
// 1 RPM = 2π/60 rad/s
func RPM(value float64) AngularVelocity {
	return RadianPerSecond(value * 0.10471975511965977) // 2π/60
}

// -----------------------------------------------------------------------------
// Electromagnetic Units
// -----------------------------------------------------------------------------

// Charge represents an electric charge with dimension [IT].
type Charge struct{ Value }

// Coulomb creates a Charge value in coulombs (A⋅s).
func Coulomb(value float64) Charge {
	return Charge{NewValue(value, Dimension{I: 1, T: 1})}
}

// Millicoulomb creates a Charge value in millicoulombs (10⁻³ C).
func Millicoulomb(value float64) Charge {
	return Coulomb(value * 1e-3)
}

// Microcoulomb creates a Charge value in microcoulombs (10⁻⁶ C).
func Microcoulomb(value float64) Charge {
	return Coulomb(value * 1e-6)
}

// ElementaryCharge creates a Charge value in multiples of the elementary charge.
// e = 1.602176634e-19 C
func ElementaryCharge(value float64) Charge {
	return Coulomb(value * 1.602176634e-19)
}

// Voltage represents an electric potential difference with dimension [L²MT⁻³I⁻¹].
type Voltage struct{ Value }

// Volt creates a Voltage value in volts (W/A = kg⋅m²/(s³⋅A)).
func Volt(value float64) Voltage {
	return Voltage{NewValue(value, Dimension{L: 2, M: 1, T: -3, I: -1})}
}

// Millivolt creates a Voltage value in millivolts (10⁻³ V).
func Millivolt(value float64) Voltage {
	return Volt(value * 1e-3)
}

// Microvolt creates a Voltage value in microvolts (10⁻⁶ V).
func Microvolt(value float64) Voltage {
	return Volt(value * 1e-6)
}

// Kilovolt creates a Voltage value in kilovolts (10³ V).
func Kilovolt(value float64) Voltage {
	return Volt(value * 1e3)
}

// Resistance represents an electrical resistance with dimension [L²MT⁻³I⁻²].
type Resistance struct{ Value }

// Ohm creates a Resistance value in ohms (V/A = kg⋅m²/(s³⋅A²)).
func Ohm(value float64) Resistance {
	return Resistance{NewValue(value, Dimension{L: 2, M: 1, T: -3, I: -2})}
}

// Milliohm creates a Resistance value in milliohms (10⁻³ Ω).
func Milliohm(value float64) Resistance {
	return Ohm(value * 1e-3)
}

// Kiloohm creates a Resistance value in kiloohms (10³ Ω).
func Kiloohm(value float64) Resistance {
	return Ohm(value * 1e3)
}

// Megaohm creates a Resistance value in megaohms (10⁶ Ω).
func Megaohm(value float64) Resistance {
	return Ohm(value * 1e6)
}

// Capacitance represents an electrical capacitance with dimension [L⁻²M⁻¹T⁴I²].
type Capacitance struct{ Value }

// Farad creates a Capacitance value in farads (C/V = s⁴⋅A²/(kg⋅m²)).
func Farad(value float64) Capacitance {
	return Capacitance{NewValue(value, Dimension{L: -2, M: -1, T: 4, I: 2})}
}

// Microfarad creates a Capacitance value in microfarads (10⁻⁶ F).
func Microfarad(value float64) Capacitance {
	return Farad(value * 1e-6)
}

// Nanofarad creates a Capacitance value in nanofarads (10⁻⁹ F).
func Nanofarad(value float64) Capacitance {
	return Farad(value * 1e-9)
}

// Picofarad creates a Capacitance value in picofarads (10⁻¹² F).
func Picofarad(value float64) Capacitance {
	return Farad(value * 1e-12)
}

// Inductance represents an electrical inductance with dimension [L²MT⁻²I⁻²].
type Inductance struct{ Value }

// Henry creates an Inductance value in henries (Wb/A = kg⋅m²/(s²⋅A²)).
func Henry(value float64) Inductance {
	return Inductance{NewValue(value, Dimension{L: 2, M: 1, T: -2, I: -2})}
}

// Millihenry creates an Inductance value in millihenries (10⁻³ H).
func Millihenry(value float64) Inductance {
	return Henry(value * 1e-3)
}

// Microhenry creates an Inductance value in microhenries (10⁻⁶ H).
func Microhenry(value float64) Inductance {
	return Henry(value * 1e-6)
}

// MagneticField represents a magnetic flux density with dimension [MT⁻²I⁻¹].
type MagneticField struct{ Value }

// Tesla creates a MagneticField value in teslas (Wb/m² = kg/(s²⋅A)).
func Tesla(value float64) MagneticField {
	return MagneticField{NewValue(value, Dimension{M: 1, T: -2, I: -1})}
}

// Millitesla creates a MagneticField value in milliteslas (10⁻³ T).
func Millitesla(value float64) MagneticField {
	return Tesla(value * 1e-3)
}

// Microtesla creates a MagneticField value in microteslas (10⁻⁶ T).
func Microtesla(value float64) MagneticField {
	return Tesla(value * 1e-6)
}

// Gauss creates a MagneticField value in gauss (10⁻⁴ T).
// The gauss is the CGS unit of magnetic flux density.
func Gauss(value float64) MagneticField {
	return Tesla(value * 1e-4)
}

// MagneticFlux represents a magnetic flux with dimension [L²MT⁻²I⁻¹].
type MagneticFlux struct{ Value }

// Weber creates a MagneticFlux value in webers (V⋅s = kg⋅m²/(s²⋅A)).
func Weber(value float64) MagneticFlux {
	return MagneticFlux{NewValue(value, Dimension{L: 2, M: 1, T: -2, I: -1})}
}

// Milliweber creates a MagneticFlux value in milliwebers (10⁻³ Wb).
func Milliweber(value float64) MagneticFlux {
	return Weber(value * 1e-3)
}

// Maxwell creates a MagneticFlux value in maxwells (10⁻⁸ Wb).
// The maxwell is the CGS unit of magnetic flux.
func Maxwell(value float64) MagneticFlux {
	return Weber(value * 1e-8)
}

// -----------------------------------------------------------------------------
// Type-Safe Operations for Derived Units
// -----------------------------------------------------------------------------

// Multiply operations that produce new unit types

// AreaMultiply returns Volume when multiplying Area by Length.
func (a Area) Multiply(l Length) Volume {
	return Volume{a.Value.Multiply(l.Value)}
}

// VelocityDivide returns Acceleration when dividing Velocity by Time.
func (v Velocity) Divide(t Time) Acceleration {
	return Acceleration{v.Value.Divide(t.Value)}
}

// LengthDivide returns Velocity when dividing Length by Time.
func (l Length) Divide(t Time) Velocity {
	return Velocity{l.Value.Divide(t.Value)}
}

// MassMultiplyAcceleration returns Force (F = ma).
func (m Mass) MultiplyAcceleration(a Acceleration) Force {
	return Force{m.Value.Multiply(a.Value)}
}

// ForceMultiply returns Energy when multiplying Force by Length (work = F⋅d).
func (f Force) Multiply(l Length) Energy {
	return Energy{f.Value.Multiply(l.Value)}
}

// EnergyDivide returns Power when dividing Energy by Time (P = E/t).
func (e Energy) Divide(t Time) Power {
	return Power{e.Value.Divide(t.Value)}
}

// PowerMultiply returns Energy when multiplying Power by Time (E = P⋅t).
func (p Power) Multiply(t Time) Energy {
	return Energy{p.Value.Multiply(t.Value)}
}

// ForceDivide returns Pressure when dividing Force by Area (P = F/A).
func (f Force) Divide(a Area) Pressure {
	return Pressure{f.Value.Divide(a.Value)}
}

// VoltageDivide returns Resistance when dividing Voltage by Current (R = V/I).
func (v Voltage) Divide(i Current) Resistance {
	return Resistance{v.Value.Divide(i.Value)}
}

// CurrentMultiply returns Charge when multiplying Current by Time (Q = I⋅t).
func (i Current) Multiply(t Time) Charge {
	return Charge{i.Value.Multiply(t.Value)}
}

// ChargeDivide returns Current when dividing Charge by Time (I = Q/t).
func (q Charge) Divide(t Time) Current {
	return Current{q.Value.Divide(t.Value)}
}
