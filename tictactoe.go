package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var game [9]string
	var name string
	var row, col int
	var isvalid bool

	fmt.Print("Masukkan Nama Anda: ")
	fmt.Scanln(&name)
	fmt.Printf("Selamat Datang Di Game TicTacToe %s\n\n", name)

	/*for i := range game {
		game[i] = "x"
	}
	printGame(game)*/
	resetGame(&game)

	for row != 99 {
		printGame(game)
		for !isvalid {
			fmt.Printf("%s Sedang Jalan. Gunakan Format Row(int) dan Column(int): ", name)
			fmt.Scanln(&row, &col)
			isvalid = playerMove(&game, row, col)

			if !isvalid {
				fmt.Printf("Coba Lagi!\n\n")
			}
		}
		isvalid = false
		row = checkWin(game, name)
		if row == 99 {
			break
		}

		compRandomMove(&game)
		fmt.Printf("Comp Sedang Jalan...\n\n")
		row = checkWin(game, name)
	}

	/*for i := 0; i < 9; i++ {
		compRandomMove(&game)
		printGame(game)
	}*/

	return
}

func resetGame(game *[9]string) {
	for i := range game {
		game[i] = " "
	}
}

func printGame(game [9]string) {
	row := 0
	for i := range game {
		if row >= 3 {
			row = 0
			fmt.Printf("\n---------\n")
		}
		switch {
		case i == 0:
			fmt.Printf("%s |", game[i])
		case i%3 == 1:
			fmt.Printf(" %s ", game[i])
		case i%3 == 2:
			fmt.Printf("| %s", game[i])
		case i%3 == 0:
			fmt.Printf("%s |", game[i])
		}

		row++
	}
	fmt.Print("\n\n")
}

func playerMove(game *[9]string, row, col int) bool {
	var move int
	switch {
	case row == 1:
		move = 3*(row-1) + col - 1
	case row == 2:
		move = 3*(row-1) + col - 1
	case row == 3:
		move = 3*(row-1) + col - 1
	}

	if game[move] == " " {
		game[move] = "x"
		return true
	}
	return false
}

func compRandomMove(game *[9]string) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	var move int
	for {
		move = r.Intn(9)
		if game[move] == " " {
			game[move] = "o"
			break
		}
	}
}

func checkWin(game [9]string, name string) int {
	isfull := true
	for i := range game {
		if game[i] == " " {
			isfull = false
		}
	}
	if isfull {
		fmt.Printf("Tidak Ada Pemenang")
		return 99
	}

	switch {
	case game[0] == "x" && game[1] == "x" && game[2] == "x":
		fmt.Printf("%s Menang!!!\n", name)
		printGame(game)
		return 99
	case game[3] == "x" && game[4] == "x" && game[5] == "x":
		fmt.Printf("%s Menang!!!\n", name)
		printGame(game)
		return 99
	case game[6] == "x" && game[7] == "x" && game[8] == "x":
		fmt.Printf("%s Menang!!!\n", name)
		printGame(game)
		return 99
	case game[0] == "o" && game[1] == "o" && game[2] == "o":
		fmt.Printf("Comp Menang, Coba Lagi\n")
		printGame(game)
		return 99
	case game[3] == "o" && game[4] == "o" && game[5] == "o":
		fmt.Printf("Comp Menang, Coba Lagi\n")
		printGame(game)
		return 99
	case game[6] == "o" && game[7] == "o" && game[8] == "o":
		fmt.Printf("Comp Menang, Coba Lagi\n")
		printGame(game)
		return 99
	case game[0] == "x" && game[4] == "x" && game[8] == "x":
		fmt.Printf("%s Menang!!!\n", name)
		printGame(game)
		return 99
	case game[2] == "x" && game[4] == "x" && game[6] == "x":
		fmt.Printf("%s Menang!!!\n", name)
		printGame(game)
		return 99
	case game[0] == "o" && game[4] == "o" && game[8] == "o":
		fmt.Printf("Comp Menang, Coba Lagi\n")
		printGame(game)
		return 99
	case game[2] == "o" && game[4] == "o" && game[6] == "o":
		fmt.Printf("Comp Menang, Coba Lagi\n")
		printGame(game)
		return 99
	case game[0] == "x" && game[3] == "x" && game[6] == "x":
		fmt.Printf("%s Menang!!!\n", name)
		printGame(game)
		return 99
	case game[1] == "x" && game[4] == "x" && game[7] == "x":
		fmt.Printf("%s Menang!!!\n", name)
		printGame(game)
		return 99
	case game[2] == "x" && game[5] == "x" && game[8] == "x":
		fmt.Printf("%s Menang!!!\n", name)
		printGame(game)
		return 99
	case game[0] == "o" && game[3] == "o" && game[6] == "o":
		fmt.Printf("Comp Menang, Coba Lagi\n")
		printGame(game)
		return 99
	case game[1] == "o" && game[4] == "o" && game[7] == "o":
		fmt.Printf("Comp Menang, Coba Lagi\n")
		printGame(game)
		return 99
	case game[2] == "o" && game[5] == "o" && game[8] == "o":
		fmt.Printf("Comp Menang, Coba Lagi\n")
		printGame(game)
		return 99
	}
	return 0
}
