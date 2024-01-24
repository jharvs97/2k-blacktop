package database

import (
	"database/sql"

	"github.com/jharvs97/2k-blacktop/model"
	_ "modernc.org/sqlite"
)

func GetConfigByName(name string) (model.TeamConfig, error) {
	row := db.QueryRow(
		"SELECT * FROM team_config WHERE name = :name",
		sql.Named("name", name))

	err := row.Err()
	if err != nil {
		return model.TeamConfig{}, err
	}

	var tc model.TeamConfig
	var id int
	err = row.Scan(&id, &tc.Name, &tc.NumGuards, &tc.NumWings, &tc.NumBigs, &tc.PlayerCount)

	if err != nil {
		return model.TeamConfig{}, err
	}

	return tc, nil
}

func GetConfigsByPlayerCount(playerCount int) ([]model.TeamConfig, error) {
	rows, err := db.Query(
		"SELECT * FROM team_config WHERE player_count = :player_count",
		sql.Named("player_count", playerCount))

	if err != nil {
		return nil, err
	}

	var configs []model.TeamConfig

	for rows.Next() {
		var tc model.TeamConfig
		var id int
		err = rows.Scan(&id, &tc.Name, &tc.NumGuards, &tc.NumWings, &tc.NumBigs, &tc.PlayerCount)
		if err != nil {
			return nil, err
		}

		configs = append(configs, tc)
	}

	return configs, nil
}

func GetTeamsForConfig(config model.TeamConfig) (model.Teams, error) {
	teams := model.Teams{}

	team1, _ := GetPlayersForConfig(config)
	team2, _ := GetPlayersForConfig(config)

	teams.Team1 = team1
	teams.Team2 = team2

	return teams, nil
}

func GetPlayersForConfig(config model.TeamConfig) ([]model.Player, error) {
	players := make([]model.Player, 0)

	guards, err := GetNRandomPlayers(config.NumGuards, model.Guard)
	if err != nil {
		return nil, err
	}
	players = append(players, guards...)

	wings, err := GetNRandomPlayers(config.NumWings, model.Wing)
	if err != nil {
		return nil, err
	}
	players = append(players, wings...)

	bigs, err := GetNRandomPlayers(config.NumBigs, model.Big)
	if err != nil {
		return nil, err
	}
	players = append(players, bigs...)

	return players, nil
}

func GetNRandomPlayers(n int, position model.Position) ([]model.Player, error) {
	rows, err := db.Query(
		"SELECT * FROM player WHERE position_id = :position ORDER BY RANDOM() LIMIT :n",
		sql.Named("position", position),
		sql.Named("n", n))

	if err != nil {
		return nil, err
	}

	var players []model.Player

	for rows.Next() {
		var p model.Player
		var id int
		err = rows.Scan(&id, &p.Name, &p.Team, &p.Position)
		if err != nil {
			return nil, err
		}

		players = append(players, p)
	}

	return players, nil
}

var db *sql.DB

func Init() error {
	var err error
	db, err = sql.Open("sqlite", "./blacktop.db")
	if err != nil {
		return err
	}

	return nil
}
