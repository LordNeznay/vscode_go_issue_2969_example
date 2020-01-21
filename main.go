package hexmap

import (
	"formatproblem/coord"
)

type TileMap struct {
	Width  int
	Height int
	Data   []int
}

func NewTileMap(width int, height int, defaultTile int) TileMap {
	res := TileMap{
		Width:  width,
		Height: height,
		Data:   make([]int, width*height),
	}

	for i := 0; i < len(res.Data); i += 1 {
		res.Data[i] = defaultTile
	}

	return res
}

func InitTileMap(width int, height int, data []int) TileMap {
	return TileMap{
		Width:  width,
		Height: height,
		Data:   data,
	}
}

func (m TileMap) calcIndex(pos coord.OddR) int {
	return pos.Q + pos.R * m.Width
}

func (m * TileMap) SetTile(tile int, pos coord.OddR) {
	m.Data[m.calcIndex(pos)] = tile
}

func (m TileMap) GetTile(pos coord.OddR) int {
	return m.Data[m.calcIndex(pos)]
}