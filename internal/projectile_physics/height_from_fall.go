package projectilephysics

func CalculateHeightFromFallDuration(fallDuration float64) float64 {
	const g = 9.81
	return 0.5 * g * (fallDuration * fallDuration)
}

func PredictFallTimeFromHeight(height float64) float64 {
	const g = 9.81
	fallTime := (2 * height / g)
	return fallTime * fallTime
}
