package gameplay

type TileTemplate struct {
	Type       string
	GetTexture TileGetTexture
	OnCreate   TileOnCreate
}

type (
	TileGetTexture func(tile *Tile) int
	TileOnCreate   func(tile *Tile)
)

func NewTileTemplate(tileType string) *TileTemplate {
	return &TileTemplate{
		Type: tileType,
		GetTexture: func(tile *Tile) int {
			return 0
		},
		OnCreate: func(tile *Tile) {
		},
	}
}
