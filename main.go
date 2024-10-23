package main

import (
	"fmt"
	"os"

	"github.com/notnil/chess"
)

func main() {
	scanFile()
}

func scanFile() {

	f, err := os.Open("lichess.pgn")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := chess.NewScanner(f)
	for scanner.Scan() {
		game := scanner.Next()
		fmt.Println(game.GetTagPair("Site"))
		fmt.Println(game.GetTagPair("ECO"))

		// Output:
		//&{Site https://lichess.org/LDtVpbOE}
		//&{ECO A00}
	}

}
