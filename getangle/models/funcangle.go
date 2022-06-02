package models

func Operation(hours, minutes float64) float64 {
	calculoM := (minutes / 60)
	total := float64(calculoM) + hours
	totalHours := (total * 30)
	min_grados := (minutes * 6)
	totalResult := (min_grados - totalHours)

	if totalResult < 0 {
		convert := -totalResult

		return convert
	}

	return totalResult
}
