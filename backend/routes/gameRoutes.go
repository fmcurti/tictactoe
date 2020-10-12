package routes

import (
	"encoding/json"
	"net/http"
	"tictactoeapi/models"
)

type NewGameRequest struct {
	Name       string `json:"Name"`
	Difficulty string `json:"Difficulty"`
}

type TicTacToeMove struct {
	Game_key string `json:"game_key"`
	Cell     int    `json:"cell"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	switch r.Method {
	case "POST":

		var p NewGameRequest

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		game := models.NewGame(p.Name, p.Difficulty)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"game_key": "` + game.Game_key + `"}`))
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(`{"message": "Method not implemented"}`))
	}
}

func GetGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		enableCors(&w)
		game_key := r.URL.Query().Get("game_key")

		if game_key != "" {
			g, found := models.GetGame(game_key)
			if found {
				jsonRes, _ := json.Marshal(g)
				w.WriteHeader(http.StatusOK)
				w.Write(jsonRes)
				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Game Not Found"}`))
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(`{"message": "Method not implemented"}`))
	}
}

func MakeMove(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	switch r.Method {
	case "POST":
		enableCors(&w)
		var m TicTacToeMove

		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if m.Cell < 0 || m.Cell > 8 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message": "Invalid Cell"}`))
		}
		g, _ := models.GetGame(m.Game_key)
		if g.Current_turn == 0 {
			if g.MakeMove(m.Cell, true) {
				jsonRes, _ := json.Marshal(m)
				w.WriteHeader(http.StatusOK)
				w.Write(jsonRes)
				return
			}
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "Invalid Move"}`))

	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(`{"message": "Method not implemented"}`))
	}
}
