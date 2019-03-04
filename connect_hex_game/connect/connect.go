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

	fmt.Printf("Tile id: [%d,%d], Tile Value: %c, Edge Count: %d\n", t.row,t.col, t.value, t.edgeCount)
}

func (t Tile) printConnected() {

	if t.headP == nil {
		fmt.Println("Tile has no connected tiles.")
		return
	} else {
		curr := t.headP

		for curr != nil {
			fmt.Print("		")
			curr.destination.printTile()
			curr = curr.nextP
		}
	}
}



type Edge struct{
	origin *Tile
	destination *Tile
	nextP *Edge
}


type Board struct {
	tiles []*Tile
	height int
	width int
}

func NewEdge (oTile *Tile, dTile *Tile) *Edge {

	e := new(Edge)
	e.origin = oTile
	e.destination = dTile
	e.nextP = nil

	return e
}

func NewBoard (sboard []string) *Board{

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

	b := new(Board)
	b.tiles = make([]*Tile, len(sboard)*len(sboard[0]) )
	b.height = len(sboard)
	b.width = len(sboard[0])


	// want to iterate over the sboard []string and create tile for each character

	for i := range sboard {
		for col, char := range sboard[i] {
			b.tiles[i*b.width + col] = NewTile(i, col, char)
		}
	}
	return b
}

func (b Board) printBoard() {
	fmt.Println("Printing Board....")
	for _, t := range b.tiles{
		t.printTile()
		t.printConnected()
	}
	fmt.Println("------------------")
}

func (b Board) returnTile(x int, y int) *Tile{

	if x > b.height || x < 0 {
		return nil
	}

	if y > b.width || y < 0 {
		return nil
	}

	for _, w := range b.tiles {
		if w.row == x && w.col == y {
			return w
		}
	}

	return nil
}

func initGraph (tboard *Board) {

	// ok so I want this function to take in a board,
	// look at it, and add the relevant edges between the nodes.
	// don't need a directed linked list, should be easier
	// or maybe not, doing it a tile at a time could mean its easier.


	for _, t := range tboard.tiles {
		i := t.row
		j := t.col

		fmt.Printf("Tile [%d, %d], Value %s\n", i, j, string(t.value))

		connected := [6]*Tile {
			tboard.returnTile(i,j - 1),
			tboard.returnTile(i+1,j - 1),
			tboard.returnTile(i-1,j),
			tboard.returnTile(i+1,j),
			tboard.returnTile(i-1,j+1),
			tboard.returnTile(i,j+1),
		}

		for _, c := range connected {
			if c != nil {
				fmt.Printf("	Connected Tile [%d, %d], Value %s\n", c.row, c.col, string(c.value))

				if t.value == c.value {
					fmt.Print("		Same Value, Adding Edge\n")
					t.edgeCount++
					newedge:= NewEdge(t, c)
					if t.headP == nil {
						t.headP = newedge
					} else {
						curr := t.headP

						for curr.nextP != nil {
							curr = curr.nextP
						}

						curr.nextP = newedge
					}
				}
			}
		}
	}

}





func ResultOf (sboard []string) (bool,error) {

	//tile := NewTile(1,1, 'X')
	board := NewBoard(sboard)

	//tile.printTile()
	//board.printBoard()

	//board.returnTile(2,2).printTile()

	initGraph(board)
	board.printBoard()



	return true, nil
}
