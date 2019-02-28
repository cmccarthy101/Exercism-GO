package connect

import (
	"fmt"
	"unicode"
)

type Tile struct {
	row int
	col int
	checked bool
	value rune
	edgeCount int
	headP *Edge
	leftP *Tile
	rightP *Tile
	queueNextP *Tile
}


func NewTile(x int, y int, value rune) *Tile {

	t:= new(Tile)
	t.row = x
	t.col = y
	t.value = value
	t.checked = false
	t.edgeCount = 0
	t.headP = nil
	t.leftP = nil
	t.rightP = nil
	t.queueNextP = nil

	return t
}

func MakeTile(x int, y int, value rune) Tile {

	 t:= Tile {x, y, false, value, 0,
					nil, nil, nil, nil }

	 return t
}


func (t Tile) printTile() {
	fmt.Printf("Tile id: [%d,%d], Tile Value: %c\n", t.row,t.col, t.value)
}


type Edge struct{
	origin *Tile
	destination *Tile
	nextP *Edge
}


type Board struct {
	tiles []Tile
	height int
	width int
}

func NewBoard (sboard []string) *Board{
	b := new(Board)
	b.tiles = make([]Tile, len(sboard)*len(sboard[0]) )
	b.height = len(sboard)
	b.width = len(sboard[0])


	// want to iterate over the sboard []string and create tile for each character

	for i := range sboard {
		for col, char := range sboard[i] {
			b.tiles[i*b.width + col] = MakeTile(i, col, char)
		}
	}
	return b
}

func (b Board) printBoard() {

	for _, t := range b.tiles{
		t.printTile()
	}
}



func ResultOf (sboard []string) (bool,error) {

	for i, s := range sboard {
		niceboard := ""
		for _, c := range s {
			if unicode.IsSpace(c){

				continue
			}
			niceboard = niceboard + string(c)

		}
		sboard[i] = niceboard
	}

	tile := NewTile(1,1, 'X')
	board := NewBoard(sboard)

	tile.printTile()
	board.printBoard()




	return true, nil
}
