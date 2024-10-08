package gameplay

type TileTemplate struct {
	Type       string
	GetTexture TileGetTexture
}

type (
	TileGetTexture func(tile *Tile) int
)

func NewTileTemplate(tileType string) *TileTemplate {
	return &TileTemplate{
		Type: tileType,
		GetTexture: func(tile *Tile) int {
			return 0
		},
	}
}
