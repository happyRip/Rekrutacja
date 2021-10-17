package waiter

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Table struct {
	Number   int `json:"number"`
	MinSeats int `json:"minNumberOfSeats"`
	MaxSeats int `json:"maxNumberOfSeats"`
}

type ListOfTables struct {
	Tables []Table `json:"tables"`
}

func (t *ListOfTables) GetSeatsFromFile(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &t)
	if err != nil {
		return err
	}
	return nil
}

func (t ListOfTables) GetTables(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := struct {
		Status    string `json:"status"`
		MinSeats  int    `json:"min_seats"`
		StartDate string `json:"start_date"`
		Duration  int    `json:"duration"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		log.Println(err)
		return
	}

	var tables []Table
	for _, v := range t.Tables {
		if params.MinSeats <= v.MaxSeats && params.MinSeats >= v.MinSeats {
			tables = append(tables, v)
		}
	}
	json.NewEncoder(w).Encode(tables)
	return
}
