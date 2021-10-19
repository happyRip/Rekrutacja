package waiter

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/happyRip/Rekrutacja/backend/portier"
)

const LayoutISO = "2006-01-02T15:04:05-0700"

var Tables ListOfTables

func init() {
	if err := Tables.getSeatsFromFile("seats.json"); err != nil {
		log.Fatal(err)
	}
}

type Table struct {
	Number   int `json:"number"`
	MinSeats int `json:"minNumberOfSeats"`
	MaxSeats int `json:"maxNumberOfSeats"`
}

type ListOfTables struct {
	Tables []Table `json:"tables"`
}

func (t *ListOfTables) getSeatsFromFile(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	return nil
}

func (t ListOfTables) GetTables(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reservations := &portier.Reservations

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

	isSeatAmountValid := func(min, max int) bool {
		if params.MinSeats <= max && params.MinSeats >= min {
			return true
		}
		return false
	}

	var tables []Table
	for _, table := range t.Tables {
		if isSeatAmountValid(table.MinSeats, table.MaxSeats) {
			tables = append(tables, table)
		}
	}

	var unavailableTables []int
	start, err := time.Parse(portier.LayoutISO, params.StartDate)
	if err != nil {
		log.Println(err)
		return
	}
	end := start.Add(time.Minute * time.Duration(params.Duration))
	for _, reservation := range reservations.Bookings {
		if reservation.IsTableOccupied(start, end) {
			unavailableTables = append(unavailableTables,
				reservation.SeatNumber,
			)
		}
	}
	for i, v := range tables {
		for j, u := range unavailableTables {
			if v.Number == u {
				tables = append(tables[:i], tables[i+1:]...)
				unavailableTables = append(
					unavailableTables[:j],
					unavailableTables[j+1:]...,
				)
				break
			}
		}
	}

	// TODO:
	// * check if params are valid
	// * send email to user
	json.NewEncoder(w).Encode(tables)
}
