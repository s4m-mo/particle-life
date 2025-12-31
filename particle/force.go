package particle

import (
	"life/settings"
	"math"
)

var (
	// Polynomial Force Parameters
	polyP float64 = 10                                  // Distance under which the force decreases then flips
	polyM float64 = settings.MaxInfluenceRadius - polyP // Width of polynomial portion
	polyA float64 = 4                                   // Power of the polynomial portion (how quickly it decays)
	polyQ float64 = 3                                   // Magnitude of repulsion at close distance

	polyMA float64 = math.Pow(polyM, polyA) // Precomputed value for later
)

// Recomputes force params that depend on other values
func RecomputeForceValues() {
	polyM = settings.MaxInfluenceRadius - polyP
}

/*
Input should alwyas be greater or equal to 0 or else negative values will appear
*/
func PolyForce(x float64, attractionFactor float64) float64 {
	switch {
	case x < polyP:
		return ((1+polyQ)*x)/polyP - polyQ
	case x < settings.MaxInfluenceRadius:
		return math.Pow(settings.MaxInfluenceRadius-x, polyA) / polyMA * attractionFactor
	default:
		return 0
	}
}
