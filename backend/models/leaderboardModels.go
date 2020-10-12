package models

type LeaderboardEntry struct {
	Player_name string `json:"player_name"`
	Play_time   int    `json:"play_time"`
}

type Leaderboard struct {
	Entries []LeaderboardEntry `json:"entries"`
	Pages   int                `json:"pages"`
}

func GetLeaderboard(limit int, offset int) Leaderboard {
	var leaderboard Leaderboard

	result, err := Db.Query(`SELECT player_name,
							TIMESTAMPDIFF(SECOND,create_date,finish_date) win_time
							FROM Games
							WHERE winner = 1
							Order by win_time
							LIMIT ? OFFSET ?
							`, limit, offset)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var entry LeaderboardEntry
		err := result.Scan(&entry.Player_name, &entry.Play_time)
		if err != nil {
			panic(err.Error())
		}
		leaderboard.Entries = append(leaderboard.Entries, entry)
	}

	result, err = Db.Query(`SELECT CEIL(COUNT(*)/?) pages from Games
							WHERE winner = 1
							`, limit)

	for result.Next() {
		err := result.Scan(&leaderboard.Pages)
		if err != nil {
			panic(err.Error())
		}
	}

	return leaderboard

}
