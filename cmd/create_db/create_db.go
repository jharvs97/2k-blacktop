package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jharvs97/2k-blacktop/model"
	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./blacktop.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	createDb(db)

	importData(db)
}

func createDb(db *sql.DB) {
	file, err := os.Open("./cmd/create_db/create.sql")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	q := string(bytes)

	res, err := db.Exec(q)
	if err != nil {
		panic(err)
	}

	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("Finished creating db, %d rows affected\n", rowsAffected)
}

func importData(db *sql.DB) {
	players := readCsv("./cmd/create_db/2015-2016.csv")

	q := "INSERT INTO player (name, team, position_id) VALUES "

	for _, g := range players.guards {
		q += fmt.Sprintf("('%s', '%s', 1),\n", g.Name, g.Team)
	}

	q += "\n"

	for _, w := range players.wings {
		q += fmt.Sprintf("('%s', '%s', 2),\n", w.Name, w.Team)
	}

	q += "\n"

	for i, b := range players.bigs {
		q += fmt.Sprintf("('%s', '%s', 3)", b.Name, b.Team)
		if i < len(players.bigs)-1 {
			q += ",\n"
		}
	}

	q += ";"

	res, err := db.Exec(q)
	if err != nil {
		panic(err)
	}
	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("Finished importing data, %d rows affected\n", rowsAffected)
}

type players struct {
	guards []model.Player
	wings  []model.Player
	bigs   []model.Player
}

func readCsv(filePath string) players {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	res := players{}
	res.guards = make([]model.Player, 0, 100)
	res.wings = make([]model.Player, 0, 100)
	res.bigs = make([]model.Player, 0, 100)

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	for _, row := range data[1:] {

		name := strings.ReplaceAll(row[1], "'", "")
		pos := row[2]
		team := row[4]

		p := model.Player{Name: name, Team: team}

		switch pos {
		case "PG":
			fallthrough
		case "PG-SG":
			res.guards = append(res.guards, p)
		case "SG":
			fallthrough
		case "SF":
			fallthrough
		case "SG-SF":
			fallthrough
		case "SF-PF":
			res.wings = append(res.wings, p)
		case "PF":
			fallthrough
		case "PF-C":
			fallthrough
		case "C":
			res.bigs = append(res.bigs, p)
		default:
			fmt.Printf("Couldn't match %s to a position\n", pos)
		}
	}

	return res
}
