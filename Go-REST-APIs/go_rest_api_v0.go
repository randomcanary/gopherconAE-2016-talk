package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

type Rawdata struct {
	RandomInt int
}

func get_random_number() int {
	return rand.Intn(100)
}

func getJSONPayload() ([]byte, error) {
	randomNum := get_random_number()
	rd := Rawdata{randomNum}
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
	http.HandleFunc("/", pageHandler)
	http.ListenAndServe("localhost:8080", nil)
}
