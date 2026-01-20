package units_test

import (
	"fmt"

	"github.com/sakiphan/qsim-core/units"
)

// Example demonstrating basic length operations.
func ExampleLength() {
	// Create length values
	length1 := units.Meter(5.0)
	length2 := units.Centimeter(200.0) // 200 cm = 2 m

	// Add lengths (type-safe)
	total, _ := length1.Value.Add(length2.Value)
	fmt.Printf("Total length: %.1f m\n", total.Val())

	// Convert to different units
	fmt.Printf("Length in km: %.3f km\n", length1.ToKilometers())
	fmt.Printf("Length in mm: %.0f mm\n", length1.ToMillimeters())

	// Output:
	// Total length: 7.0 m
	// Length in km: 0.005 km
	// Length in mm: 5000 mm
}

// Example demonstrating velocity calculation.
func ExampleVelocity() {
	distance := units.Kilometer(100.0)
	time := units.Hour(2.0)

	// Calculate velocity: v = d / t
	velocity := distance.Value.Divide(time.Value)

	fmt.Printf("Velocity: %.1f m/s\n", velocity.Val())
	fmt.Printf("Velocity: %.1f km/h\n", velocity.Val()*3.6)

	// Output:
	// Velocity: 13.9 m/s
	// Velocity: 50.0 km/h
}

// Example demonstrating kinetic energy calculation.
func ExampleEnergy_kineticEnergy() {
	// KE = ½mv²
	mass := units.Kilogram(2.0)
	velocity := units.MeterPerSecond(10.0)

	// Calculate v²
	vSquared := velocity.Value.Multiply(velocity.Value)

	// Calculate KE = ½ × m × v²
	kineticEnergy := mass.Value.Multiply(vSquared).Scale(0.5)

	fmt.Printf("Kinetic Energy: %.0f J\n", kineticEnergy.Val())
	fmt.Printf("In electron volts: %.2e eV\n", kineticEnergy.Val()/1.602176634e-19)

	// Output:
	// Kinetic Energy: 100 J
	// In electron volts: 6.24e+20 eV
}

// Example demonstrating force calculation using Newton's second law.
func ExampleForce_newtonSecondLaw() {
	// F = ma
	mass := units.Kilogram(10.0)
	acceleration := units.MeterPerSecond2(9.8)

	force := mass.MultiplyAcceleration(acceleration)

	fmt.Printf("Force: %.0f N\n", force.ToNewtons())

	// Output:
	// Force: 98 N
}

// Example demonstrating work-energy calculation.
func ExampleEnergy_work() {
	// Work = Force × Distance
	force := units.Newton(50.0)
	distance := units.Meter(10.0)

	work := force.Multiply(distance)

	fmt.Printf("Work done: %.0f J\n", work.ToJoules())
	fmt.Printf("In calories: %.2f cal\n", work.ToCalories())

	// Output:
	// Work done: 500 J
	// In calories: 119.50 cal
}

// Example demonstrating power calculation.
func ExamplePower() {
	// Power = Energy / Time
	energy := units.Kilojoule(10.0) // 10 kJ
	time := units.Second(5.0)

	power := energy.Divide(time)

	fmt.Printf("Power: %.0f W\n", power.ToWatts())
	fmt.Printf("In horsepower: %.3f hp\n", power.ToHorsepower())

	// Output:
	// Power: 2000 W
	// In horsepower: 2.682 hp
}

// Example demonstrating pressure calculation.
func ExamplePressure() {
	// Pressure = Force / Area
	force := units.Newton(1000.0)
	area := units.SquareMeter(2.0)

	pressure := force.Divide(area)

	fmt.Printf("Pressure: %.0f Pa\n", pressure.ToPascals())
	fmt.Printf("In atmospheres: %.5f atm\n", pressure.ToAtmospheres())
	fmt.Printf("In PSI: %.2f psi\n", pressure.ToPSI())

	// Output:
	// Pressure: 500 Pa
	// In atmospheres: 0.00493 atm
	// In PSI: 0.07 psi
}

// Example demonstrating Ohm's law.
func ExampleResistance_ohmsLaw() {
	// R = V / I
	voltage := units.Volt(12.0)
	current := units.Ampere(0.5)

	resistance := voltage.Divide(current)

	fmt.Printf("Resistance: %.0f Ω\n", resistance.ToOhms())

	// Output:
	// Resistance: 24 Ω
}

// Example demonstrating electric charge calculation.
func ExampleCharge() {
	// Q = I × t
	current := units.Ampere(2.0)
	time := units.Second(10.0)

	charge := current.Multiply(time)

	fmt.Printf("Charge: %.0f C\n", charge.ToCoulombs())
	fmt.Printf("Elementary charges: %.2e e\n", charge.ToElementaryCharges())

	// Output:
	// Charge: 20 C
	// Elementary charges: 1.25e+20 e
}

// Example demonstrating electrical energy calculation.
func ExampleEnergy_electrical() {
	// E = P × t
	power := units.Kilowatt(2.0) // 2 kW appliance
	time := units.Hour(3.0)      // Running for 3 hours

	energy := power.Multiply(time)

	fmt.Printf("Energy consumed: %.1f kWh\n", energy.ToJoules()/3.6e6)
	fmt.Printf("In joules: %.2e J\n", energy.ToJoules())

	// Output:
	// Energy consumed: 6.0 kWh
	// In joules: 2.16e+07 J
}

// Example demonstrating frequency and period relationship.
func ExampleFrequency() {
	// f = 1 / T
	period := units.Millisecond(10.0) // 10 ms period

	// Frequency = 1 / Period
	frequency := units.Dimensionless(1.0).Divide(period.Value)

	fmt.Printf("Frequency: %.0f Hz\n", frequency.Val())

	// Output:
	// Frequency: 100 Hz
}

// Example demonstrating temperature conversions.
func ExampleTemperature() {
	tempC := units.Celsius(100.0)   // Boiling point of water
	tempF := units.Fahrenheit(32.0) // Freezing point of water

	fmt.Printf("Boiling point: %.2f K\n", tempC.ToKelvin())
	fmt.Printf("Boiling point: %.1f °C\n", tempC.ToCelsius())

	fmt.Printf("Freezing point: %.2f K\n", tempF.ToKelvin())
	fmt.Printf("Freezing point: %.1f °C\n", tempF.ToCelsius())

	// Output:
	// Boiling point: 373.15 K
	// Boiling point: 100.0 °C
	// Freezing point: 273.15 K
	// Freezing point: 0.0 °C
}

// Example demonstrating astronomical distances.
func ExampleLength_astronomical() {
	earthSunDistance := units.AstronomicalUnit(1.0)
	proximaDistance := units.LightYear(4.24)

	fmt.Printf("Earth-Sun distance: %.2e m\n", earthSunDistance.ToMeters())
	fmt.Printf("In light-years: %.8f ly\n", earthSunDistance.ToLightYears())

	fmt.Printf("\nProxima Centauri distance: %.2e m\n", proximaDistance.ToMeters())
	fmt.Printf("In parsecs: %.2f pc\n", proximaDistance.ToParsecs())

	// Output:
	// Earth-Sun distance: 1.50e+11 m
	// In light-years: 0.00001581 ly
	//
	// Proxima Centauri distance: 4.01e+16 m
	// In parsecs: 1.30 pc
}

// Example demonstrating mass of celestial objects.
func ExampleMass_astronomical() {
	sunMass := units.SolarMass(1.0)
	earthMass := units.EarthMass(1.0)

	fmt.Printf("Solar mass: %.2e kg\n", sunMass.ToKilograms())
	fmt.Printf("Earth mass: %.2e kg\n", earthMass.ToKilograms())

	// Sun / Earth mass ratio
	ratio := sunMass.Value.Divide(earthMass.Value)
	fmt.Printf("Sun is %.0f times more massive than Earth\n", ratio.Val())

	// Output:
	// Solar mass: 1.99e+30 kg
	// Earth mass: 5.97e+24 kg
	// Sun is 333030 times more massive than Earth
}

// Example demonstrating type safety - compilation errors.
func Example_typeSafety() {
	length := units.Meter(10.0)
	mass := units.Kilogram(5.0)

	// This would cause a compile-time error:
	// _, err := length.Value.Add(mass.Value)
	// Error: cannot add quantities with different dimensions

	// Instead, we can only add compatible units:
	length2 := units.Meter(5.0)
	totalLength, _ := length.Value.Add(length2.Value)

	fmt.Printf("Total length: %.0f m\n", totalLength.Val())

	// Multiplication works across dimensions:
	// This creates a new dimension [L M]
	combined := length.Value.Multiply(mass.Value)
	fmt.Printf("Combined dimension: %s\n", combined.Dim().String())

	// Output:
	// Total length: 15 m
	// Combined dimension: [L^1 M^1]
}

// Example demonstrating area calculation.
func ExampleArea() {
	length := units.Meter(5.0)
	width := units.Meter(3.0)

	// Area = length × width
	area := length.Value.Multiply(width.Value)

	fmt.Printf("Area: %.0f m²\n", area.Val())
	fmt.Printf("In square centimeters: %.0f cm²\n", area.Val()*1e4)

	// Output:
	// Area: 15 m²
	// In square centimeters: 150000 cm²
}

// Example demonstrating volume calculation.
func ExampleVolume() {
	length := units.Meter(2.0)
	width := units.Meter(3.0)
	height := units.Meter(4.0)

	// Volume = length × width × height
	area := length.Value.Multiply(width.Value)
	volume := area.Multiply(height.Value)

	fmt.Printf("Volume: %.0f m³\n", volume.Val())
	fmt.Printf("In liters: %.0f L\n", volume.Val()*1e3)

	// Output:
	// Volume: 24 m³
	// In liters: 24000 L
}

// Example demonstrating density calculation.
func Example_density() {
	mass := units.Kilogram(8.0)
	volume := units.Liter(10.0) // 10 L = 0.01 m³

	// Density = mass / volume [kg/m³]
	density := mass.Value.Divide(volume.Value)

	fmt.Printf("Density: %.0f kg/m³\n", density.Val())

	// Output:
	// Density: 800 kg/m³
}

// Example demonstrating acceleration due to gravity.
func ExampleAcceleration() {
	// Standard gravity
	g := units.StandardGravity(1.0)

	fmt.Printf("Standard gravity: %.2f m/s²\n", g.Value.Val())

	// Free fall: v = gt
	time := units.Second(5.0)
	velocity := g.Value.Multiply(time.Value)

	fmt.Printf("Velocity after 5s of free fall: %.1f m/s\n", velocity.Val())

	// Output:
	// Standard gravity: 9.81 m/s²
	// Velocity after 5s of free fall: 49.0 m/s
}

// Example demonstrating RPM to angular velocity conversion.
func ExampleAngularVelocity() {
	rpm := units.RPM(3000.0) // 3000 RPM (typical car engine)

	fmt.Printf("Angular velocity: %.2f rad/s\n", rpm.Value.Val())

	// Period = 2π / ω
	period := units.Dimensionless(6.283185307).Divide(rpm.Value) // 2π / ω

	fmt.Printf("Period: %.4f s\n", period.Val())
	fmt.Printf("Frequency: %.2f Hz\n", 1.0/period.Val())

	// Output:
	// Angular velocity: 314.16 rad/s
	// Period: 0.0200 s
	// Frequency: 50.00 Hz
}

// Example demonstrating energy units in particle physics.
func ExampleEnergy_particlePhysics() {
	// Rest mass energy of electron: E = mc²
	electronMass := units.AtomicMassUnit(0.000548579909) // Electron mass in u
	c := units.SpeedOfLight(1.0)
	cSquared := c.Value.Multiply(c.Value)

	energy := electronMass.Value.Multiply(cSquared)

	fmt.Printf("Electron rest mass energy: %.3e J\n", energy.Val())
	fmt.Printf("In MeV: %.3f MeV\n", energy.Val()/(1.602176634e-19*1e6))

	// Alternative: direct conversion
	electronEnergy := units.MegaelectronVolt(0.511)
	fmt.Printf("Electron rest mass: %.3f MeV (standard value)\n", electronEnergy.ToMeV())

	// Output:
	// Electron rest mass energy: 8.187e-14 J
	// In MeV: 0.511 MeV
	// Electron rest mass: 0.511 MeV (standard value)
}
