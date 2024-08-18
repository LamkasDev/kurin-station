package gameplay

import "github.com/veandco/go-sdl2/sdl"

type KurinForce struct {
	Item   *KurinItem
	Target sdl.FPoint
	Delta  sdl.FPoint
}
