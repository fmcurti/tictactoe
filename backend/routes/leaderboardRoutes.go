package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tictactoeapi/models"
)

func GetHighscores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		enableCors(&w)
		limit := r.URL.Query().Get("limit")
		offset := r.URL.Query().Get("offset")
		if offset != "" && limit != "" {
			limit, _ := strconv.Atoi(limit)
			offset, _ := strconv.Atoi(offset)
			leaderboard := models.GetLeaderboard(limit, offset)
			jsonRes, _ := json.Marshal(leaderboard)
			w.WriteHeader(http.StatusOK)
			w.Write(jsonRes)
			return
		}
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Invalid parameters"}`))
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(`{"message": "Method not implemented"}`))
	}
}
