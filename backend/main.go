package main

import (
	"database/sql"
	"log"
	"math/rand"
	"net/http"
	"os"
	"tictactoeapi/models"
	"tictactoeapi/routes"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func make_best_move(g models.Game) {
	bestScore := -64000
	var bestMove int

	for i, cell := range g.Board {
		if cell == "E" {
			g.Board[i] = "O"
			score := minimax(g, false)
			g.Current_turn = 1 - g.Current_turn
			g.Board[i] = "E"
			if score > bestScore {
				bestScore = score
				bestMove = i
			}
		}
	}
	g.MakeMove(bestMove, true)

}

func minimax(g models.Game, maximise bool) int {

	g.CheckForWinners()
	if g.Winner > 0 {
		if g.Winner == 3 {
			return 0
		} else {
			if maximise {
				return -1
			} else {
				return 1
			}

		}
	}

	emptyCells := getEmptyCells(g)

	var scores []int
	for _, i := range emptyCells {
		if maximise {
			g.Board[i] = "O"
		} else {
			g.Board[i] = "X"
		}
		scores = append(scores, minimax(g, !maximise))
		g.Board[i] = "E"
	}

	min, max := MinMax(scores)

	if maximise {
		return max
	} else {
		return min
	}
}

func getEmptyCells(g models.Game) []int {
	var emptyCells []int
	for i, cell := range g.Board {
		if cell == "E" {
			emptyCells = append(emptyCells, i)
		}
	}
	return emptyCells
}

func playPendingGames() {
	for {
		result, err := models.Db.Query("SELECT game_key FROM Games WHERE current_turn = 1")

		if err != nil {
			panic(err.Error())
		}

		defer result.Close()

		var game_key string

		for result.Next() {
			err := result.Scan(&game_key)
			if err != nil {
				panic(err.Error())
			}

			g, _ := models.GetGame(game_key)

			if g.Difficulty == "Easy" {
				var emptyCells []int
				for i, cell := range g.Board {
					if cell == "E" {
						emptyCells = append(emptyCells, i)
					}
				}
				randomCell := rand.Intn(len(emptyCells))
				g.MakeMove(emptyCells[randomCell], true)
			} else {
				make_best_move(*g)
			}

		}
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	models.Db, models.Err = sql.Open("mysql", os.Getenv("JAWSDB_MARIA_URL"))
	if models.Err != nil {
		panic(models.Err.Error())
	}

	models.InitializeDB()

	defer models.Db.Close()

	go playPendingGames()

	http.HandleFunc("/game/new", routes.NewGame)
	http.HandleFunc("/game/get", routes.GetGame)
	http.HandleFunc("/game/make_move", routes.MakeMove)

	http.HandleFunc("/highscores", routes.GetHighscores)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
