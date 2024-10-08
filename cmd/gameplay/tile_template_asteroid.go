package gameplay

func NewTileTemplateAsteroid() *TileTemplate {
	template := NewTileTemplate("asteroid")
	template.GetTexture = func(tile *Tile) int {
		if tile.Seed%100 < 3 {
			return 1
		}

		return 0
	}

	return template
}
