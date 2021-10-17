package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type Reservation struct {
	Date          string `json:"date"`
	Duration      int    `json:"duration"`
	SeatNumber    int    `json:"seatNumber"`
	FullName      string `json:"fullName"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	NumberOfSeats int    `json:"numberOfSeats"`
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
