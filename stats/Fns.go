package stats

import (
	"math"
)

// Average computes an average (mean) value of a set of floating-point numbers.
func Average(vals *Floats) (res float64) {
	len := len(*vals)
	if len > 0 {
		for _, v := range *vals {
			res += v
		}

		res /= float64(len)
	}

	return
}

// Dispersion computes a dispersion of a set of floating-point numbers.
func Dispersion(vals *Floats) (res float64) {
	avg := Average(vals)
	len := len(*vals)

	if len > 0 {
		for _, v := range *vals {
			d := v - avg
			res += d * d
		}

		if len >= 2 {
			len = len - 1
		}

		res /= float64(len)
	}

	return
}

// StdDev computes a standard deviation of a set of floating-point numbers.
func StdDev(vals *Floats) float64 {
	return math.Sqrt(Dispersion(vals))
}
