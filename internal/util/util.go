package util

import "math"

func ByteToGigabyte(value uint64) float64 {
	return math.Ceil(float64(value) / 1073741824.0)
}
