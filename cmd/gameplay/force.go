package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

type ForceResult uint8

var (
	ForceResultOutOfBounds = ForceResult(0)
	ForceResultCollided    = ForceResult(1)
	ForceResultReached     = ForceResult(2)
	ForceResultNone        = ForceResult(3)
)

type Force struct {
	Position     sdlutils.FVector3
	Target       sdl.FPoint
	Delta        sdl.FPoint
	IgnoredMobId uint32
	Data         interface{}
}

func NewForce(position sdlutils.FVector3, target sdl.FPoint, ignoredMobId uint32, data interface{}) *Force {
	// this shit is trash
	force := &Force{
		Position:     position,
		Target:       target,
		IgnoredMobId: ignoredMobId,
		Data:         data,
	}
	force.Delta = sdlutils.SubtractFPoints(force.Target, position.Base)
	distance := sdlutils.GetDistanceF(position.Base, force.Target)
	if distance != 0 {
		force.Delta.X /= distance * 3
		force.Delta.Y /= distance * 3
	}
	if IsMapPositionOutOfBounds(GameInstance.Map, sdlutils.Vector3{Base: sdlutils.FPointToPointFloored(target), Z: position.Z}) {
		force.Target.X += force.Delta.X * 100
		force.Target.Y += force.Delta.Y * 100
	}

	return force
}

func AdvanceForce(force *Force) (ForceResult, interface{}) {
	force.Position.Base = sdlutils.AddFPoints(force.Position.Base, force.Delta)
	testVector := sdlutils.Vector3{Base: sdlutils.FPointToPointFloored(force.Position.Base), Z: force.Position.Z}
	if IsMapPositionOutOfBounds(GameInstance.Map, testVector) {
		return ForceResultOutOfBounds, nil
	}
	tile := GetTileAt(GameInstance.Map, testVector)
	if tile != nil {
		if CanEnterMapPosition(GameInstance.Map, testVector) != EnteranceStatusYes {
			return ForceResultCollided, GetObjectAtTile(tile)
		}
		mobs := GetMobsOnTileWithout(GameInstance.Map, tile, force.IgnoredMobId)
		if len(mobs) > 0 {
			return ForceResultCollided, mobs[0]
		}
	}
	if sdlutils.GetDistanceF(force.Position.Base, force.Target) < 0.01 {
		return ForceResultReached, nil
	}

	return ForceResultNone, nil
}

func RushForce(force *Force) (ForceResult, interface{}) {
	for {
		result, collider := AdvanceForce(force)
		if result != ForceResultNone {
			return result, collider
		}
	}
}
