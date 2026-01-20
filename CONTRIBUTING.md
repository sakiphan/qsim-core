# Contributing to qsim

Thank you for your interest in contributing to qsim! This document provides guidelines and standards for contributing to the project.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Workflow](#development-workflow)
- [Code Standards](#code-standards)
- [Adding New Physics Formulas](#adding-new-physics-formulas)
- [Testing Requirements](#testing-requirements)
- [Documentation Standards](#documentation-standards)
- [Pull Request Process](#pull-request-process)

## Code of Conduct

- Be respectful and professional
- Focus on constructive feedback
- Prioritize scientific accuracy and academic rigor
- Cite sources for all formulas and algorithms

## Getting Started

1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/YOUR_USERNAME/qsim-core.git
   cd qsim-core
   ```
3. Create a feature branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```

## Development Workflow

1. Make your changes following code standards
2. Write tests (unit, validation, and benchmarks where applicable)
3. Run tests and linting:
   ```bash
   go test ./...
   go fmt ./...
   go vet ./...
   ```
4. Commit with descriptive messages
5. Push to your fork and open a pull request

## Code Standards

### Go Style

- Follow [Effective Go](https://golang.org/doc/effective_go)
- Use `gofmt` for formatting
- Pass `go vet` and `staticcheck` linting
- Keep functions focused and small
- Use meaningful variable names (prefer physics notation where appropriate)

### Naming Conventions

- Use physics notation in documentation: `ψ`, `ℏ`, `Δx`
- Use descriptive Go names in code: `waveFunction`, `planckReduced`, `deltaX`
- Types use PascalCase: `Energy`, `WaveFunction`
- Functions use PascalCase for public: `KineticEnergy`
- Variables use camelCase: `totalEnergy`, `initialVelocity`

### Unit Safety

- **Always** use unit-safe types for physical quantities
- Document units in comments: `// velocity: m/s`, `// energy: J`
- Ensure dimensional consistency in all operations

Example:
```go
// Good
func KineticEnergy(m units.Mass, v units.Velocity) units.Energy {
    return units.Joule(0.5 * m.Value() * v.Value() * v.Value())
}

// Bad - no unit safety
func KineticEnergy(m, v float64) float64 {
    return 0.5 * m * v * v
}
```

## Adding New Physics Formulas

When adding a new formula, follow these steps:

### 1. Research and Validation

- Find authoritative sources (textbooks, peer-reviewed papers)
- Verify the formula against multiple sources
- Note any assumptions or limitations

### 2. Implementation

```go
// Example structure
package classical

import "github.com/sakiphan/qsim-core/units"

// KineticEnergy calculates the kinetic energy of an object.
//
// The kinetic energy is given by KE = ½mv², where m is the mass
// and v is the velocity of the object.
//
// Parameters:
//   - m: Mass of the object (kg)
//   - v: Velocity of the object (m/s)
//
// Returns:
//   - Energy in joules (kg⋅m²/s²)
//
// Formula:
//   KE = ½mv²
//
// Example:
//
//  m := units.Kilogram(2.0)
//  v := units.MeterPerSecond(3.0)
//  ke := classical.KineticEnergy(m, v)
//  fmt.Printf("Kinetic energy: %v\n", ke)  // Output: 9.0 J
//
// References:
//   - Halliday, Resnick, Walker. "Fundamentals of Physics", 10th ed., Ch. 7
//   - Goldstein, "Classical Mechanics", 3rd ed., Ch. 1
func KineticEnergy(m units.Mass, v units.Velocity) units.Energy {
    val := 0.5 * m.Value() * v.Value() * v.Value()
    return units.Joule(val)
}
```

### 3. Key Elements

- **Clear description**: What does the function calculate?
- **Parameters**: Physical meaning and units
- **Returns**: Physical meaning and units
- **Formula**: Mathematical representation
- **Example**: Working code example
- **References**: Academic sources (textbooks, papers)

## Testing Requirements

### 1. Unit Tests

Every public function must have unit tests:

```go
func TestKineticEnergy(t *testing.T) {
    tests := []struct {
        name     string
        mass     units.Mass
        velocity units.Velocity
        want     units.Energy
    }{
        {
            name:     "basic calculation",
            mass:     units.Kilogram(2.0),
            velocity: units.MeterPerSecond(3.0),
            want:     units.Joule(9.0), // ½ * 2 * 3² = 9
        },
        {
            name:     "zero velocity",
            mass:     units.Kilogram(5.0),
            velocity: units.MeterPerSecond(0.0),
            want:     units.Joule(0.0),
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := KineticEnergy(tt.mass, tt.velocity)
            if !almostEqual(got.Value(), tt.want.Value(), 1e-10) {
                t.Errorf("KineticEnergy() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### 2. Validation Tests

Compare with published results:

```go
func TestQuantumHarmonicOscillator_Validation(t *testing.T) {
    // Reference: Griffiths, "Introduction to Quantum Mechanics", Ch. 2.3
    // Ground state energy for quantum harmonic oscillator: E₀ = ½ℏω

    solver := schrodinger.TimeIndependentSolver{
        Mass:      constants.ElectronMass,
        Potential: schrodinger.HarmonicPotential(1.0),
        // ... other parameters
    }

    energies, _, err := solver.Solve()
    if err != nil {
        t.Fatal(err)
    }

    omega := 1.0
    expectedE0 := 0.5 * constants.PlanckReduced.Value() * omega

    // Allow 1% error due to numerical approximation
    if !almostEqual(energies[0].Value(), expectedE0, 0.01) {
        t.Errorf("Ground state energy = %v, want %v (Griffiths Ch. 2.3)",
            energies[0], expectedE0)
    }
}
```

### 3. Benchmarks

For performance-critical code:

```go
func BenchmarkKineticEnergy(b *testing.B) {
    m := units.Kilogram(2.0)
    v := units.MeterPerSecond(3.0)

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        KineticEnergy(m, v)
    }
}
```

## Documentation Standards

### GoDoc Comments

- First sentence is a summary (appears in package overview)
- Use present tense: "KineticEnergy calculates..." not "KineticEnergy will calculate..."
- Include mathematical formulas using Unicode: `E = mc²`
- Provide working examples in Example tests
- Cite academic references

### Academic References

Use this format:
```
// References:
//   - Author. "Book Title", Edition, Chapter/Section
//   - Author et al. "Paper Title", Journal, Year, DOI
```

Example:
```
// References:
//   - Halliday, Resnick, Walker. "Fundamentals of Physics", 10th ed., Ch. 7
//   - Einstein, A. "Does the Inertia of a Body Depend Upon Its Energy Content?",
//     Annalen der Physik, 1905, doi:10.1002/andp.19053231314
```

## Pull Request Process

1. **Before submitting**:
   - All tests pass: `go test ./...`
   - Code is formatted: `go fmt ./...`
   - No lint errors: `go vet ./...`
   - Documentation is complete
   - Examples are provided

2. **PR Description**:
   - Describe what the PR adds/fixes
   - Reference any related issues
   - List academic sources used
   - Include example usage if applicable

3. **Review Process**:
   - Maintainers will review for:
     - Scientific accuracy
     - Code quality
     - Test coverage
     - Documentation completeness
   - Address review comments
   - Maintain professional discourse

4. **After Approval**:
   - PR will be merged by maintainer
   - Changes will be included in next release

## Questions?

- Open an issue for questions about contributing
- Check existing issues and PRs for similar discussions
- Be patient - we're all volunteers!

## Thank You!

Your contributions help make physics computations more accessible and reliable for researchers and students worldwide.
