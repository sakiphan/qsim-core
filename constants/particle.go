package constants

import "github.com/sakiphan/qsim-core/units"

// Particle Properties
// All values from CODATA 2018 recommended values.

// -----------------------------------------------------------------------------
// Electron Properties
// -----------------------------------------------------------------------------

// ElectronMass is the electron rest mass (m_e).
// Value: 9.1093837015(28) × 10⁻³¹ kg
// Relative standard uncertainty: 3.0 × 10⁻¹⁰
//
// References:
//   - CODATA 2018
var ElectronMass = units.Kilogram(9.1093837015e-31)

// ElectronCharge is the electron charge (-e).
// Value: -1.602176634 × 10⁻¹⁹ C (exact magnitude)
//
// Note: This is the negative of the elementary charge.
//
// References:
//   - CODATA 2018
var ElectronCharge = units.Coulomb(-1.602176634e-19)

// ElectronRestEnergy is the electron rest energy (m_e c²).
// Value: 8.1871057769(25) × 10⁻¹⁴ J = 0.51099895000(15) MeV
// Relative standard uncertainty: 3.0 × 10⁻¹⁰
//
// References:
//   - CODATA 2018
var ElectronRestEnergy = units.Joule(8.1871057769e-14)

// ElectronRestEnergyMeV is the electron rest energy in MeV.
// Value: 0.51099895000(15) MeV
//
// References:
//   - CODATA 2018
var ElectronRestEnergyMeV = 0.51099895000

// ElectronMagneticMoment is the electron magnetic moment (μ_e).
// Value: -9.2847647043(28) × 10⁻²⁴ J/T
// Relative standard uncertainty: 3.0 × 10⁻¹⁰
//
// References:
//   - CODATA 2018
var ElectronMagneticMoment = units.NewValue(-9.2847647043e-24, units.Dimension{L: 2, I: 1})

// ElectronGFactor is the electron g-factor.
// Value: -2.00231930436256(35)
// Relative standard uncertainty: 1.7 × 10⁻¹³
//
// Relates magnetic moment to spin: μ_e = g_e μ_B s
//
// References:
//   - CODATA 2018
var ElectronGFactor = -2.00231930436256

// ElectronComptonWavelength is the Compton wavelength of the electron (λ_C).
// Value: 2.42631023867(73) × 10⁻¹² m
// Relative standard uncertainty: 3.0 × 10⁻¹⁰
//
// λ_C = h/(m_e c)
//
// References:
//   - CODATA 2018
var ElectronComptonWavelength = units.Meter(2.42631023867e-12)

// -----------------------------------------------------------------------------
// Proton Properties
// -----------------------------------------------------------------------------

// ProtonMass is the proton rest mass (m_p).
// Value: 1.67262192369(51) × 10⁻²⁷ kg
// Relative standard uncertainty: 3.1 × 10⁻¹⁰
//
// References:
//   - CODATA 2018
var ProtonMass = units.Kilogram(1.67262192369e-27)

// ProtonCharge is the proton charge (+e).
// Value: +1.602176634 × 10⁻¹⁹ C (exact)
//
// Equal in magnitude to the elementary charge.
//
// References:
//   - CODATA 2018
var ProtonCharge = units.Coulomb(1.602176634e-19)

// ProtonRestEnergy is the proton rest energy (m_p c²).
// Value: 1.50327761598(46) × 10⁻¹⁰ J = 938.27208816(29) MeV
// Relative standard uncertainty: 3.1 × 10⁻¹⁰
//
// References:
//   - CODATA 2018
var ProtonRestEnergy = units.Joule(1.50327761598e-10)

// ProtonRestEnergyMeV is the proton rest energy in MeV.
// Value: 938.27208816(29) MeV
//
// References:
//   - CODATA 2018
var ProtonRestEnergyMeV = 938.27208816

// ProtonMagneticMoment is the proton magnetic moment (μ_p).
// Value: 1.41060679736(60) × 10⁻²⁶ J/T
// Relative standard uncertainty: 4.3 × 10⁻¹⁰
//
// References:
//   - CODATA 2018
var ProtonMagneticMoment = units.NewValue(1.41060679736e-26, units.Dimension{L: 2, I: 1})

// ProtonGFactor is the proton g-factor.
// Value: 5.5856946893(16)
// Relative standard uncertainty: 2.9 × 10⁻¹⁰
//
// References:
//   - CODATA 2018
var ProtonGFactor = 5.5856946893

// ProtonComptonWavelength is the Compton wavelength of the proton (λ_C,p).
// Value: 1.32140985539(40) × 10⁻¹⁵ m
// Relative standard uncertainty: 3.0 × 10⁻¹⁰
//
// λ_C,p = h/(m_p c)
//
// References:
//   - CODATA 2018
var ProtonComptonWavelength = units.Meter(1.32140985539e-15)

// -----------------------------------------------------------------------------
// Neutron Properties
// -----------------------------------------------------------------------------

// NeutronMass is the neutron rest mass (m_n).
// Value: 1.67492749804(95) × 10⁻²⁷ kg
// Relative standard uncertainty: 5.7 × 10⁻¹⁰
//
// References:
//   - CODATA 2018
var NeutronMass = units.Kilogram(1.67492749804e-27)

// NeutronCharge is the neutron charge.
// Value: 0 C (neutral)
//
// The neutron has no net electric charge.
//
// References:
//   - Experimental limit: |q_n| < 10⁻²¹ e
var NeutronCharge = units.Coulomb(0.0)

// NeutronRestEnergy is the neutron rest energy (m_n c²).
// Value: 1.50534976287(86) × 10⁻¹⁰ J = 939.56542052(54) MeV
// Relative standard uncertainty: 5.7 × 10⁻¹⁰
//
// References:
//   - CODATA 2018
var NeutronRestEnergy = units.Joule(1.50534976287e-10)

// NeutronRestEnergyMeV is the neutron rest energy in MeV.
// Value: 939.56542052(54) MeV
//
// References:
//   - CODATA 2018
var NeutronRestEnergyMeV = 939.56542052

// NeutronMagneticMoment is the neutron magnetic moment (μ_n).
// Value: -9.6623651(23) × 10⁻²⁷ J/T
// Relative standard uncertainty: 2.4 × 10⁻⁷
//
// The neutron has a magnetic moment despite being electrically neutral.
//
// References:
//   - CODATA 2018
var NeutronMagneticMoment = units.NewValue(-9.6623651e-27, units.Dimension{L: 2, I: 1})

// NeutronGFactor is the neutron g-factor.
// Value: -3.82608545(90)
// Relative standard uncertainty: 2.4 × 10⁻⁷
//
// References:
//   - CODATA 2018
var NeutronGFactor = -3.82608545

// NeutronComptonWavelength is the Compton wavelength of the neutron (λ_C,n).
// Value: 1.31959090581(75) × 10⁻¹⁵ m
// Relative standard uncertainty: 5.7 × 10⁻¹⁰
//
// λ_C,n = h/(m_n c)
//
// References:
//   - CODATA 2018
var NeutronComptonWavelength = units.Meter(1.31959090581e-15)

// NeutronMeanLifetime is the neutron mean lifetime (τ_n).
// Value: 879.4(6) s (PDG 2020)
// Uncertainty: ±0.6 s
//
// Free neutrons decay via beta decay: n → p + e⁻ + ν̄_e
//
// References:
//   - Particle Data Group 2020
var NeutronMeanLifetime = units.Second(879.4)

// -----------------------------------------------------------------------------
// Muon Properties
// -----------------------------------------------------------------------------

// MuonMass is the muon rest mass (m_μ).
// Value: 1.883531627(42) × 10⁻²⁸ kg
// Relative standard uncertainty: 2.2 × 10⁻⁸
//
// The muon is a heavier version of the electron.
//
// References:
//   - CODATA 2018
var MuonMass = units.Kilogram(1.883531627e-28)

// MuonCharge is the muon charge (-e).
// Value: -1.602176634 × 10⁻¹⁹ C (exact magnitude)
//
// References:
//   - CODATA 2018
var MuonCharge = units.Coulomb(-1.602176634e-19)

// MuonRestEnergyMeV is the muon rest energy in MeV.
// Value: 105.6583755(23) MeV
// Relative standard uncertainty: 2.2 × 10⁻⁸
//
// References:
//   - CODATA 2018
var MuonRestEnergyMeV = 105.6583755

// MuonMeanLifetime is the muon mean lifetime (τ_μ).
// Value: 2.1969811(22) × 10⁻⁶ s
// Relative standard uncertainty: 1.0 × 10⁻⁵
//
// Muons decay via weak interaction: μ⁻ → e⁻ + ν̄_e + ν_μ
//
// References:
//   - Particle Data Group 2020
var MuonMeanLifetime = units.Microsecond(2.1969811)

// -----------------------------------------------------------------------------
// Tau Properties
// -----------------------------------------------------------------------------

// TauMass is the tau lepton rest mass (m_τ).
// Value: 3.16754(21) × 10⁻²⁷ kg
// Relative standard uncertainty: 6.6 × 10⁻⁵
//
// The tau is the heaviest charged lepton.
//
// References:
//   - Particle Data Group 2020
var TauMass = units.Kilogram(3.16754e-27)

// TauCharge is the tau charge (-e).
// Value: -1.602176634 × 10⁻¹⁹ C (exact magnitude)
//
// References:
//   - CODATA 2018
var TauCharge = units.Coulomb(-1.602176634e-19)

// TauRestEnergyMeV is the tau rest energy in MeV.
// Value: 1776.86(12) MeV
// Relative standard uncertainty: 6.8 × 10⁻⁵
//
// References:
//   - Particle Data Group 2020
var TauRestEnergyMeV = 1776.86

// TauMeanLifetime is the tau mean lifetime (τ_τ).
// Value: 2.903(5) × 10⁻¹³ s
// Relative standard uncertainty: 1.7 × 10⁻³
//
// Tau decays via weak interaction to lighter leptons.
//
// References:
//   - Particle Data Group 2020
var TauMeanLifetime = units.Second(2.903e-13)

// -----------------------------------------------------------------------------
// Mass Ratios
// -----------------------------------------------------------------------------

// ProtonElectronMassRatio is the ratio m_p/m_e.
// Value: 1836.15267343(11)
// Relative standard uncertainty: 6.0 × 10⁻¹¹
//
// The proton is about 1836 times heavier than the electron.
//
// References:
//   - CODATA 2018
var ProtonElectronMassRatio = 1836.15267343

// NeutronElectronMassRatio is the ratio m_n/m_e.
// Value: 1838.68366173(89)
// Relative standard uncertainty: 4.8 × 10⁻¹⁰
//
// References:
//   - CODATA 2018
var NeutronElectronMassRatio = 1838.68366173

// NeutronProtonMassRatio is the ratio m_n/m_p.
// Value: 1.00137841931(49)
// Relative standard uncertainty: 4.9 × 10⁻¹⁰
//
// The neutron is slightly heavier than the proton.
//
// References:
//   - CODATA 2018
var NeutronProtonMassRatio = 1.00137841931

// MuonElectronMassRatio is the ratio m_μ/m_e.
// Value: 206.768283(46)
// Relative standard uncertainty: 2.2 × 10⁻⁸
//
// The muon is about 207 times heavier than the electron.
//
// References:
//   - CODATA 2018
var MuonElectronMassRatio = 206.768283

// -----------------------------------------------------------------------------
// Nuclear Constants
// -----------------------------------------------------------------------------

// DeuteronMass is the deuteron rest mass (m_d).
// Value: 3.3435837724(10) × 10⁻²⁷ kg
// Relative standard uncertainty: 3.0 × 10⁻¹⁰
//
// The deuteron is the nucleus of deuterium (heavy hydrogen).
//
// References:
//   - CODATA 2018
var DeuteronMass = units.Kilogram(3.3435837724e-27)

// HelionMass is the helion rest mass (m_h).
// Value: 5.0064127796(15) × 10⁻²⁷ kg
// Relative standard uncertainty: 3.0 × 10⁻¹⁰
//
// The helion is the nucleus of helium-3.
//
// References:
//   - CODATA 2018
var HelionMass = units.Kilogram(5.0064127796e-27)

// AlphaParticleMass is the alpha particle rest mass (m_α).
// Value: 6.6446573357(20) × 10⁻²⁷ kg
// Relative standard uncertainty: 3.0 × 10⁻¹⁰
//
// The alpha particle is the nucleus of helium-4.
//
// References:
//   - CODATA 2018
var AlphaParticleMass = units.Kilogram(6.6446573357e-27)

// -----------------------------------------------------------------------------
// Weak Interaction
// -----------------------------------------------------------------------------

// FermiCouplingConstant is the Fermi coupling constant (G_F/(ℏc)³).
// Value: 1.1663787(6) × 10⁻⁵ GeV⁻²
//
// Characterizes the strength of the weak interaction.
//
// References:
//   - Particle Data Group 2020
var FermiCouplingConstant = 1.1663787e-5 // GeV⁻²

// WeakMixingAngle is the Weinberg angle (sin²θ_W).
// Value: 0.23122(4) (MS-bar scheme at m_Z)
//
// Mixing angle in electroweak theory.
//
// References:
//   - Particle Data Group 2020
var WeakMixingAngle = 0.23122

// -----------------------------------------------------------------------------
// Strong Interaction
// -----------------------------------------------------------------------------

// StrongCouplingConstant is the strong coupling constant (α_s).
// Value: 0.1179(10) at m_Z (MS-bar scheme)
//
// Characterizes the strength of the strong interaction.
//
// References:
//   - Particle Data Group 2020
var StrongCouplingConstant = 0.1179

// -----------------------------------------------------------------------------
// Gauge Boson Masses
// -----------------------------------------------------------------------------

// WBosonMassMeV is the W boson mass in MeV.
// Value: 80379(12) MeV/c²
// Uncertainty: ±12 MeV
//
// The W boson mediates weak interactions.
//
// References:
//   - Particle Data Group 2020
var WBosonMassMeV = 80379.0

// ZBosonMassMeV is the Z boson mass in MeV.
// Value: 91187.6(2.1) MeV/c²
// Uncertainty: ±2.1 MeV
//
// The Z boson mediates weak interactions.
//
// References:
//   - Particle Data Group 2020
var ZBosonMassMeV = 91187.6

// HiggsMassMeV is the Higgs boson mass in MeV.
// Value: 125090(24) MeV/c²
// Uncertainty: ±24 MeV
//
// The Higgs boson is responsible for electroweak symmetry breaking.
//
// References:
//   - ATLAS and CMS Collaborations, combined result
var HiggsMassMeV = 125090.0
