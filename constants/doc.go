// Package constants provides fundamental physical constants with high precision.
//
// All values are based on the CODATA 2018 internationally recommended values
// and Planck 2018 cosmological parameters where applicable.
//
// Constants are provided as unit-safe types from the units package, ensuring
// dimensional consistency in calculations.
//
// Example usage:
//
//	import (
//	    "github.com/sakiphan/qsim-core/constants"
//	    "github.com/sakiphan/qsim-core/units"
//	)
//
//	// Calculate photon energy: E = hν
//	frequency := units.Gigahertz(5.0) // 5 GHz
//	energy := constants.PlanckConstant.Value.Multiply(frequency.Value)
//
//	// Calculate Schwarzschild radius: Rs = 2GM/c²
//	mass := units.SolarMass(10.0)
//	c2 := constants.SpeedOfLight.Value.Multiply(constants.SpeedOfLight.Value)
//	rs := constants.GravitationalConstant.Value.Multiply(mass.Value).Scale(2.0).Divide(c2)
//
// References:
//   - CODATA 2018: Tiesinga et al., "CODATA recommended values of the fundamental
//     physical constants: 2018", Rev. Mod. Phys. 93, 025010 (2021)
//   - Planck 2018: Planck Collaboration, "Planck 2018 results. VI. Cosmological parameters",
//     A&A 641, A6 (2020)
package constants
