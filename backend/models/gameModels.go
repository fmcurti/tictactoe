package models

import (
	"crypto/rand"
	b64 "encoding/base64"
	"strings"
	"time"
)

var winningLines = [][3]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {2, 4, 6}}

type board []string

type Game struct {
	Game_key     string `json:"game_key"`
	Player_name  string `json:"player_name"`
	Difficulty   string `json:"difficulty"`
	Current_turn int8   `json:"current_turn"`
	Board        board  `json:"board"`
	Winner       int8   `json:"winner"`
}

func (b board) isValidMove(cell int) bool {
	return b[cell] == "E"
}

func (b board) boardToString() string {

	var res string
	for _, cell := range b {
		res += cell + "|"
	}
	return res[0 : len(res)-1]

}

func boardFromString(strBoard string) board {
	return strings.Split(strBoard, "|")
}

func newGameId() string {
	b := make([]byte, 9)
	rand.Read(b)
	return b64.StdEncoding.EncodeToString(b)
}

func NewGame(player string, difficulty string) *Game {
	g := new(Game)
	g.Game_key = newGameId()
	g.Player_name = player
	g.Difficulty = difficulty
	g.Current_turn = 0
	g.Winner = 0

	stmt, err := Db.Prepare("INSERT INTO Games(game_key,player_name,difficulty,create_date) VALUES(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(g.Game_key, g.Player_name, g.Difficulty, time.Now())

	if err != nil {
		panic(err.Error())
	}

	return g
}

func GetGame(game_key string) (*Game, bool) {
	g := new(Game)
	var boardString string
	found := false
	result, err := Db.Query("SELECT game_key,player_name,difficulty,current_turn,board,winner FROM Games WHERE game_key = ?", game_key)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		err := result.Scan(&g.Game_key, &g.Player_name, &g.Difficulty, &g.Current_turn, &boardString, &g.Winner)
		if err != nil {
			panic(err.Error())
		}
		g.Board = boardFromString(boardString)
		found = true
	}

	return g, found
}

func (g Game) thereIsTie() bool {

	for _, cell := range g.Board {
		if cell == "E" {
			return false
		}
	}
	return true
}

func (g *Game) CheckForWinners() bool {

	board := g.Board
	for _, l := range winningLines {

		if board[l[0]] != "E" && board[l[0]] == board[l[1]] && board[l[0]] == board[l[2]] {
			g.Winner = g.Current_turn + 1
			return true
		}
	}

	if g.thereIsTie() {
		g.Winner = 3
		return true
	}

	return false
}

func (g Game) update() {

	stmt, err := Db.Prepare(`UPDATE Games SET 
							Current_turn = ?,
							Board        = ?
							WHERE game_key = ?`)
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(g.Current_turn, g.Board.boardToString(), g.Game_key)

	if err != nil {
		panic(err.Error())
	}

}

func (g Game) closeGame() {

	stmt, err := Db.Prepare(`UPDATE Games SET 
							Winner 		 = ?,
							Board 		 = ?,
							finish_date = NOW()
							WHERE game_key = ?`)
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(g.Winner, g.Board.boardToString(), g.Game_key)

	if err != nil {
		panic(err.Error())
	}

}

func (g *Game) endTurn() {
	g.Current_turn = 1 - g.Current_turn
}

func (g Game) gameFinished() bool {
	return g.Winner != 0
}

func (g *Game) MakeMove(cell int, save bool) bool {

	if g.Board.isValidMove(cell) && !g.gameFinished() {
		if g.Current_turn == 0 {
			g.Board[cell] = "X"
		} else {
			g.Board[cell] = "O"
		}

		if g.CheckForWinners() {
			if save {
				g.closeGame()
			}

		} else {
			g.endTurn()
			if save {
				g.update()
			}
		}
		return true

	} else {
		return false
	}

}
