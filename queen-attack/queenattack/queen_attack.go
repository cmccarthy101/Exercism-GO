package queenattack

import (
	"errors"
	"math"
)

func onSameSquare (wPos string, bPos string) bool {

	if wPos == bPos {
		return true
	} else {
		return false
	}
}

func pieceOnInvalidSquare(pos string) (invalid bool) {
	if len(pos) >2 {
		return true
	}
	col := pos[0]
	row := pos[1]

	ColOutOfRange := col < 'a' || col > 'h'
	RowOutOfRange := row < '1' || row > '8'

	if ColOutOfRange || RowOutOfRange {
		return true
	}

	return false
}

func pieceCanAttack(wPos string, bPos string) bool {
	return piecesOnSameRowOrCol(wPos, bPos) || piecesOnSameDiagonal(wPos, bPos)
}

func piecesOnSameRowOrCol(wPos string, bPos string) bool {
	wCol := wPos[0]
	wRow := wPos[1]
	bCol := bPos[0]
	bRow := bPos[1]

	if wCol == bCol || wRow == bRow {
		return true
	}

	return false
}

func piecesOnSameDiagonal (wPos string, bPos string) bool {

	wColNum := float64(wPos[0])
	bColNum := float64(bPos[0])
	wRowNum := float64(wPos[1])
	bRowNum := float64(bPos[1])

	if  math.Abs(wColNum-bColNum) == math.Abs(wRowNum-bRowNum) {
		return true
	}

	return false
}

func CanQueenAttack(wQueen string, bQueen string) (bool, error) {

	if onSameSquare(wQueen, bQueen) {
		return false, errors.New("same square")
	}
	if pieceOnInvalidSquare(wQueen) || pieceOnInvalidSquare(bQueen) {
		return false, errors.New("invalid coordinates")
	}
	if pieceCanAttack(wQueen, bQueen) {
		return true, nil
	}

	return false, nil
}
