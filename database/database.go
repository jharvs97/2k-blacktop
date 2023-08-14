package database

import (
	"database/sql"
	"fmt"

	"github.com/jharvs97/2k-blacktop/model"
	_ "github.com/mattn/go-sqlite3"
)

func GetConfigByName(name string) (model.TeamConfig, error) {
	fmt.Println(name)
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

	fmt.Println(tc)

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

func GetNRandomPlayers(n int, position model.Position) ([]model.Player, error) {
	fmt.Println(position)
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
	db, err = sql.Open("sqlite3", "./blacktop.db")
	if err != nil {
		return err
	}

	return nil
}
