package mathutils

func Lerp(a float32, b float32, f float32) float32 {
	return a*(1.0-f) + (b * f)
}
