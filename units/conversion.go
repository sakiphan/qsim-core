package units

// This file provides utility functions for unit conversions and value extraction.

// -----------------------------------------------------------------------------
// Value Extraction Helpers
// -----------------------------------------------------------------------------

// ToMeters returns the length value in meters.
func (l Length) ToMeters() float64 {
	return l.Val()
}

// ToKilometers returns the length value in kilometers.
func (l Length) ToKilometers() float64 {
	return l.Val() / 1e3
}

// ToCentimeters returns the length value in centimeters.
func (l Length) ToCentimeters() float64 {
	return l.Val() * 1e2
}

// ToMillimeters returns the length value in millimeters.
func (l Length) ToMillimeters() float64 {
	return l.Val() * 1e3
}

// ToNanometers returns the length value in nanometers.
func (l Length) ToNanometers() float64 {
	return l.Val() * 1e9
}

// ToAngstroms returns the length value in angstroms.
func (l Length) ToAngstroms() float64 {
	return l.Val() * 1e10
}

// ToAU returns the length value in astronomical units.
func (l Length) ToAU() float64 {
	return l.Val() / 1.495978707e11
}

// ToLightYears returns the length value in light-years.
func (l Length) ToLightYears() float64 {
	return l.Val() / 9.4607304725808e15
}

// ToParsecs returns the length value in parsecs.
func (l Length) ToParsecs() float64 {
	return l.Val() / 3.0856775814913673e16
}

// ToKilograms returns the mass value in kilograms.
func (m Mass) ToKilograms() float64 {
	return m.Val()
}

// ToGrams returns the mass value in grams.
func (m Mass) ToGrams() float64 {
	return m.Val() * 1e3
}

// ToMilligrams returns the mass value in milligrams.
func (m Mass) ToMilligrams() float64 {
	return m.Val() * 1e6
}

// ToTonnes returns the mass value in metric tonnes.
func (m Mass) ToTonnes() float64 {
	return m.Val() / 1e3
}

// ToPounds returns the mass value in pounds.
func (m Mass) ToPounds() float64 {
	return m.Val() / 0.45359237
}

// ToSolarMasses returns the mass value in solar masses.
func (m Mass) ToSolarMasses() float64 {
	return m.Val() / 1.98892e30
}

// ToEarthMasses returns the mass value in Earth masses.
func (m Mass) ToEarthMasses() float64 {
	return m.Val() / 5.9722e24
}

// ToSeconds returns the time value in seconds.
func (t Time) ToSeconds() float64 {
	return t.Val()
}

// ToMilliseconds returns the time value in milliseconds.
func (t Time) ToMilliseconds() float64 {
	return t.Val() * 1e3
}

// ToMicroseconds returns the time value in microseconds.
func (t Time) ToMicroseconds() float64 {
	return t.Val() * 1e6
}

// ToNanoseconds returns the time value in nanoseconds.
func (t Time) ToNanoseconds() float64 {
	return t.Val() * 1e9
}

// ToMinutes returns the time value in minutes.
func (t Time) ToMinutes() float64 {
	return t.Val() / 60.0
}

// ToHours returns the time value in hours.
func (t Time) ToHours() float64 {
	return t.Val() / 3600.0
}

// ToDays returns the time value in days.
func (t Time) ToDays() float64 {
	return t.Val() / 86400.0
}

// ToYears returns the time value in Julian years.
func (t Time) ToYears() float64 {
	return t.Val() / 31557600.0
}

// ToAmperes returns the current value in amperes.
func (i Current) ToAmperes() float64 {
	return i.Val()
}

// ToMilliamperes returns the current value in milliamperes.
func (i Current) ToMilliamperes() float64 {
	return i.Val() * 1e3
}

// ToKelvin returns the temperature value in kelvins.
func (t Temperature) ToKelvin() float64 {
	return t.Val()
}

// ToCelsius returns the temperature value in degrees Celsius.
func (t Temperature) ToCelsius() float64 {
	return t.Val() - 273.15
}

// ToFahrenheit returns the temperature value in degrees Fahrenheit.
func (t Temperature) ToFahrenheit() float64 {
	return (t.Val() * 9.0 / 5.0) - 459.67
}

// ToJoules returns the energy value in joules.
func (e Energy) ToJoules() float64 {
	return e.Val()
}

// ToKilojoules returns the energy value in kilojoules.
func (e Energy) ToKilojoules() float64 {
	return e.Val() / 1e3
}

// ToCalories returns the energy value in thermochemical calories.
func (e Energy) ToCalories() float64 {
	return e.Val() / 4.184
}

// ToKilocalories returns the energy value in kilocalories.
func (e Energy) ToKilocalories() float64 {
	return e.Val() / 4184.0
}

// ToElectronVolts returns the energy value in electron volts.
func (e Energy) ToElectronVolts() float64 {
	return e.Val() / 1.602176634e-19
}

// ToKeV returns the energy value in kiloelectron volts.
func (e Energy) ToKeV() float64 {
	return e.ToElectronVolts() / 1e3
}

// ToMeV returns the energy value in megaelectron volts.
func (e Energy) ToMeV() float64 {
	return e.ToElectronVolts() / 1e6
}

// ToGeV returns the energy value in gigaelectron volts.
func (e Energy) ToGeV() float64 {
	return e.ToElectronVolts() / 1e9
}

// ToNewtons returns the force value in newtons.
func (f Force) ToNewtons() float64 {
	return f.Val()
}

// ToKilonewtons returns the force value in kilonewtons.
func (f Force) ToKilonewtons() float64 {
	return f.Val() / 1e3
}

// ToPoundsForce returns the force value in pounds-force.
func (f Force) ToPoundsForce() float64 {
	return f.Val() / 4.4482216152605
}

// ToWatts returns the power value in watts.
func (p Power) ToWatts() float64 {
	return p.Val()
}

// ToKilowatts returns the power value in kilowatts.
func (p Power) ToKilowatts() float64 {
	return p.Val() / 1e3
}

// ToMegawatts returns the power value in megawatts.
func (p Power) ToMegawatts() float64 {
	return p.Val() / 1e6
}

// ToHorsepower returns the power value in mechanical horsepower.
func (p Power) ToHorsepower() float64 {
	return p.Val() / 745.69987158227022
}

// ToPascals returns the pressure value in pascals.
func (p Pressure) ToPascals() float64 {
	return p.Val()
}

// ToKilopascals returns the pressure value in kilopascals.
func (p Pressure) ToKilopascals() float64 {
	return p.Val() / 1e3
}

// ToBars returns the pressure value in bars.
func (p Pressure) ToBars() float64 {
	return p.Val() / 1e5
}

// ToAtmospheres returns the pressure value in standard atmospheres.
func (p Pressure) ToAtmospheres() float64 {
	return p.Val() / 101325.0
}

// ToPSI returns the pressure value in pounds per square inch.
func (p Pressure) ToPSI() float64 {
	return p.Val() / 6894.757293168
}

// ToHertz returns the frequency value in hertz.
func (f Frequency) ToHertz() float64 {
	return f.Val()
}

// ToKilohertz returns the frequency value in kilohertz.
func (f Frequency) ToKilohertz() float64 {
	return f.Val() / 1e3
}

// ToMegahertz returns the frequency value in megahertz.
func (f Frequency) ToMegahertz() float64 {
	return f.Val() / 1e6
}

// ToGigahertz returns the frequency value in gigahertz.
func (f Frequency) ToGigahertz() float64 {
	return f.Val() / 1e9
}

// ToMeterPerSecond returns the velocity value in meters per second.
func (v Velocity) ToMeterPerSecond() float64 {
	return v.Val()
}

// ToKilometerPerHour returns the velocity value in kilometers per hour.
func (v Velocity) ToKilometerPerHour() float64 {
	return v.Val() * 3.6
}

// ToMilePerHour returns the velocity value in miles per hour.
func (v Velocity) ToMilePerHour() float64 {
	return v.Val() / 0.44704
}

// ToSpeedOfLight returns the velocity value as a fraction of the speed of light.
func (v Velocity) ToSpeedOfLight() float64 {
	return v.Val() / 299792458.0
}

// ToVolts returns the voltage value in volts.
func (v Voltage) ToVolts() float64 {
	return v.Val()
}

// ToMillivolts returns the voltage value in millivolts.
func (v Voltage) ToMillivolts() float64 {
	return v.Val() * 1e3
}

// ToKilovolts returns the voltage value in kilovolts.
func (v Voltage) ToKilovolts() float64 {
	return v.Val() / 1e3
}

// ToOhms returns the resistance value in ohms.
func (r Resistance) ToOhms() float64 {
	return r.Val()
}

// ToKiloohms returns the resistance value in kiloohms.
func (r Resistance) ToKiloohms() float64 {
	return r.Val() / 1e3
}

// ToMegaohms returns the resistance value in megaohms.
func (r Resistance) ToMegaohms() float64 {
	return r.Val() / 1e6
}

// ToCoulombs returns the charge value in coulombs.
func (q Charge) ToCoulombs() float64 {
	return q.Val()
}

// ToMillicoulombs returns the charge value in millicoulombs.
func (q Charge) ToMillicoulombs() float64 {
	return q.Val() * 1e3
}

// ToElementaryCharges returns the charge value in elementary charges.
func (q Charge) ToElementaryCharges() float64 {
	return q.Val() / 1.602176634e-19
}

// ToTeslas returns the magnetic field value in teslas.
func (b MagneticField) ToTeslas() float64 {
	return b.Val()
}

// ToGauss returns the magnetic field value in gauss.
func (b MagneticField) ToGauss() float64 {
	return b.Val() * 1e4
}

// ToSquareMeters returns the area value in square meters.
func (a Area) ToSquareMeters() float64 {
	return a.Val()
}

// ToSquareKilometers returns the area value in square kilometers.
func (a Area) ToSquareKilometers() float64 {
	return a.Val() / 1e6
}

// ToHectares returns the area value in hectares.
func (a Area) ToHectares() float64 {
	return a.Val() / 1e4
}

// ToCubicMeters returns the volume value in cubic meters.
func (v Volume) ToCubicMeters() float64 {
	return v.Val()
}

// ToLiters returns the volume value in liters.
func (v Volume) ToLiters() float64 {
	return v.Val() * 1e3
}

// ToMilliliters returns the volume value in milliliters.
func (v Volume) ToMilliliters() float64 {
	return v.Val() * 1e6
}
