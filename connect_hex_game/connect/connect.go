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
	b.width = len(sboard[0])
	b.height = len(sboard)


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

	if x > b.width || x < 0 {
		return nil
	}

	if y > b.height || y < 0 {
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

		//fmt.Printf("Tile [%d, %d], Value %s\n", i, j, string(t.value))

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
				//fmt.Printf("	Connected Tile [%d, %d], Value %s\n", c.row, c.col, string(c.value))

				if t.value == c.value {
					//fmt.Print("		Same Value, Adding Edge\n")
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

func (b Board) checkWinner () string {

	// first start with checking X
	// X goes left to right.
	// make a queue of

	Xwin := checkX(b)

	Owin := checkY(b)

	if Xwin {
		return "X"
	} else if Owin {
		return "O"
	} else {
		return ""
	}
}

func checkX (b Board) (winner bool) {

	// left to right

	var queue []*Tile

	for j := 0; j < b.height; j ++ {
		if b.returnTile(j, 0).value == 'X' {

			queue= append(queue, b.returnTile(j, 0))
		}
	}

	for _, t := range queue {

		if isWinner(b, t) {
			return true
		}
		winner = depthFirstSearch(b, t)
	}


	return
}

func checkY (b Board) (winner bool) {

	// bottom to top

	var queue []*Tile

	for j := 0; j < b.width; j ++ {
		if b.returnTile(0, j).value == 'O' {

			queue= append(queue, b.returnTile(0, j))
		}
	}

	for _, t := range queue {
		if isWinner(b, t) {
			return true
		}
		winner = depthFirstSearch(b, t)
	}


	return

}

func depthFirstSearch (b Board, t *Tile) bool {

	curr := t.headP

	t.checked = true

	if curr == nil {
		return false
	} else {
		for curr != nil {
			if isWinner(b, curr.destination) {
				return true
			}
			if !curr.destination.checked {
				return depthFirstSearch(b, curr.destination)
			}
			curr = curr.nextP
		}
	}

	return false
}

func isWinner(b Board, t *Tile) bool {

	if t.value == 'X' {
		if t.col + 1 == b.width {
			return true
		}
	} else if t.value == 'O' {
		if t.row + 1 == b.height {
			return true
		}
	}

	return false
}




func ResultOf (sboard []string) (string,error) {

	//tile := NewTile(1,1, 'X')
	board := NewBoard(sboard)

	//tile.printTile()
	//board.printBoard()

	//board.returnTile(0,0).printTile()

	initGraph(board)
	//board.printBoard()

	winner := board.checkWinner()

	return winner, nil
}
