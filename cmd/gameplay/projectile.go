package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"golang.org/x/exp/slices"
)

type Projectile struct {
	Type     string
	Position sdlutils.FVector3
	Source   interface{}
}

func AddProjectileToMap(kmap *Map, projectile *Projectile) *Projectile {
	kmap.Projectiles = append(kmap.Projectiles, projectile)

	return projectile
}

func RemoveProjectileFromMap(kmap *Map, projectile *Projectile) bool {
	i := slices.Index(kmap.Projectiles, projectile)
	if i == -1 {
		return false
	}

	kmap.Projectiles = slices.Delete(kmap.Projectiles, i, i+1)
	return true
}
