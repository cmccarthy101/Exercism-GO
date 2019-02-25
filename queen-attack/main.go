package main

import (
	"./queenattack"
	"fmt"
)
func main() {

	attack, err := queenattack.CanQueenAttack("d1", "g4")

	fmt.Println(attack, err)


}