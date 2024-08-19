package sdlutils

import "github.com/veandco/go-sdl2/sdl"

type Vector3 struct {
	Base sdl.Point
	Z    uint8
}

type FVector3 struct {
	Base sdl.FPoint
	Z    uint8
}

func Vector3ToFVector3(vector3 Vector3) FVector3 {
	return FVector3{
		Base: PointToFPoint(vector3.Base),
		Z:    vector3.Z,
	}
}

func Vector3ToFVector3Center(vector3 Vector3) FVector3 {
	return FVector3{
		Base: PointToFPointCenter(vector3.Base),
		Z:    vector3.Z,
	}
}

func FVector3ToVector3(fvector3 FVector3) Vector3 {
	return Vector3{
		Base: FPointToPoint(fvector3.Base),
		Z:    fvector3.Z,
	}
}

func CompareVector3(a Vector3, b Vector3) bool {
	return a.Base.X == b.Base.X && a.Base.Y == b.Base.Y && a.Z == b.Z
}
