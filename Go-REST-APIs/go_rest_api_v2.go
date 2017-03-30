package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"sort"
)

var db *sql.DB

type City struct {
	CityName string
}

type Cities []City

type Rawdata struct {
	Citynames Cities
}

func (slice Cities) Len() int {
	return len(slice)
}

func (slice Cities) Less(i, j int) bool {
	return slice[i].CityName < slice[j].CityName
}

func (slice Cities) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func getPayloadComponents() []City {
	var (
		Name   string
		cities Cities
	)
	rows, err := db.Query("select Name from City")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&Name)
		someCity := City{Name}

		cities = append(cities, someCity)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	sort.Sort(cities)
	return cities
}

func getJSONPayload() ([]byte, error) {
	cities := getPayloadComponents()
	rd := Rawdata{cities}
	return json.MarshalIndent(rd, "", "  ")
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := getJSONPayload()
	if err != nil {
		fmt.Println("Error!")
	}

	fmt.Fprintf(w, string(resp))
}

func main() {
	db, _ = sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/world")
	db.SetMaxIdleConns(25)
	db.SetMaxOpenConns(25)
	defer db.Close()

	http.HandleFunc("/", pageHandler)
	http.ListenAndServe("localhost:8080", nil)
}
