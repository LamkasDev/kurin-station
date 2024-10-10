package gameplay

func NewTileTemplateAsteroid() *TileTemplate {
	template := NewTileTemplate("asteroid")
	template.GetTexture = func(tile *Tile) int {
		return int(tile.Seed)
	}
	template.OnCreate = func(tile *Tile) {
		texture := uint8(0)
		if tile.Seed < 1 {
			texture = 5
		} else if tile.Seed < 2 {
			texture = 12
		} else if tile.Seed < 3 {
			texture = 8
		} else if tile.Seed < 4 {
			texture = 9
		} else if tile.Seed < 5 {
			texture = 2
		} else if tile.Seed < 6 {
			texture = 6
		}

		tile.Seed = texture
	}

	return template
}
