package gameplay

import "github.com/veandco/go-sdl2/sdl"

type Force struct {
	Item   *Item
	Target sdl.FPoint
	Delta  sdl.FPoint
}
