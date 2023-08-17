package mathutils

func MinFloat32(a, b float32) float32 {
	if a < b {
		return a
	}

	return b
}

func MaxFloat32(a, b float32) float32 {
	if a > b {
		return a
	}

	return b
}

func ClampFloat32(v, lo, hi float32) float32 {
	return MinFloat32(MaxFloat32(v, lo), hi)
}
