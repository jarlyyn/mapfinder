package localmap

import (
	"modules/components/textmap"
	"strings"
)

type LocalMap struct {
	ID       string
	Name     string
	Map      *textmap.TextMap
	Raw      string
	TileList []*Tile
}

type Tile struct {
	Data   string
	Map    *LocalMap
	Left   int
	Top    int
	Width  int
	Height int
}

func (t *Tile) Clone() *Tile {
	return &Tile{
		Data:   t.Data,
		Map:    t.Map,
		Left:   t.Left,
		Top:    t.Top,
		Width:  t.Width,
		Height: t.Height,
	}
}
func Create(id string, name string, data string) *LocalMap {
	m := textmap.Import(Replacerfilters.Replace(data))
	lm := &LocalMap{
		Raw:      data,
		ID:       id,
		Name:     name,
		Map:      m,
		TileList: []*Tile{},
	}
	for x := 0; x <= m.Width-TileWidth; x = x + 2 {
		for y := 0; y <= m.Width-TileHeight; y = y + 1 {
			tile := &Tile{
				Data:   strings.Join(m.Crop(x, y, TileWidth, TileHeight), "\n"),
				Map:    lm,
				Left:   x,
				Top:    y,
				Width:  TileWidth,
				Height: TileHeight,
			}
			lm.TileList = append(lm.TileList, tile)
		}
	}
	return lm
}
