package portier

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const LayoutISO = "2006-01-02T15:04:05-0700"

var Reservations ListOfReservations

func init() {
	if err := Reservations.getBookingsFromFile("bookings.json"); err != nil {
		log.Fatal(err)
	}
}

type Reservation struct {
	Date          string `json:"date"`
	Duration      int    `json:"duration"`
	SeatNumber    int    `json:"seatNumber"`
	FullName      string `json:"fullName"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	NumberOfSeats int    `json:"numberOfSeats"`
	ID            string `json:"id"`
}

func (r Reservation) isWithinTimeFrame(check ...time.Time) bool {
	start, err := time.Parse(LayoutISO, r.Date)
	if err != nil {
		log.Println(err)
		return false
	}
	end := start.Add(time.Minute * time.Duration(r.Duration))

	for _, c := range check {
		if start.Before(c) && end.After(c) {
			return true
		} else if start.Equal(c) || end.Equal(c) {
			return true
		}
	}
	return false
}

func (r Reservation) IsTableOccupied(start, end time.Time) bool {
	s, err := time.Parse(LayoutISO, r.Date)
	if err != nil {
		log.Println(err)
		return false
	}
	e := start.Add(time.Minute * time.Duration(r.Duration))

	if r.isWithinTimeFrame(start, end) || start.Before(s) && end.After(e) {
		return true
	}
	return false
}

type ListOfReservations struct {
	Bookings []Reservation `json:"bookings"`
}

func (b *ListOfReservations) NewReservation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params Reservation
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		log.Println(err)
		return
	}
}

func (b *ListOfReservations) GetReservations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := struct {
		Date string `json:"date"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		log.Println(err)
		return
	}
}

func (b *ListOfReservations) CancelReservation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// id, err := strconv.Atoi(mux.Vars(r)["id"])
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println("id:", id)
}

func (b *ListOfReservations) DeleteReservation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// id, err := strconv.Atoi(mux.Vars(r)["id"])
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println("id:", id)
}

func (b *ListOfReservations) getBookingsFromFile(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &b); err != nil {
		return err
	}
	return nil
}
