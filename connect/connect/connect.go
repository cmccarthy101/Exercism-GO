package connect


func ResultOf (board []string) (bool,error) {

	// player O plays top to bottom
	// player X plays left to right

	// a rhombus board

	//". X X . .",
	//" X . X . X",
	//"  . X . X .",
	//"   . X X . .",
	//"    O O O O O"


	// similar coordinates to chess, a3, b4 etc.
	// every square is connected to 6 other squares (if not an edge piece)
	// for example b2 is connected to
	// b1, b3, c1, c2, a2, c4
	// so for a given [x][i] it is connected to
	//	  		[x][i-1]	[x+1][i-1]
	//	[x-1][i]		[x][i]   	[x+1][i]
	//			[x-1][i+1]	[x][i+1]	[x+1][i+1]

	//". X X . .",
	//"X . X . X",
	//". X . X .",
	//". X X . .",
	//"O O O O O"


	// 'X' = 88
	// 'O' = 79
	// '.' = 46

	// convert board into a a graph paths are only there if same token on squares

	// use some maybe bfs or dfs on it's own
	// is any node on the right in any node on the lefts routing table

	// make routing table

	// first, make graph.
	// how the fuck do you make a graph again?
	//





	return true, nil
}
