package utils

func Clamp(val, min, max float64) float64 {
	if val < min {
		return min
	}

	if val > max {
		return max
	}

	return val
}
