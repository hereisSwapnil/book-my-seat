package httptransport

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/hereisSwapnil/book-my-seat/internal/domain"
	"github.com/hereisSwapnil/book-my-seat/internal/service"
)

type Server struct {
	SeatService *service.SeatService
}

func (s *Server) ListSeats(w http.ResponseWriter, r *http.Request) {
	statuses := s.SeatService.ListSeats()

	type seatDTO struct {
		ID     int    `json:"id"`
		Status string `json:"status"`
	}

	resp := make([]seatDTO, 0, len(statuses))
	for id, st := range statuses {
		status := "available"
		if st == domain.Held {
			status = "held"
		} else if st == domain.Booked {
			status = "booked"
		}

		resp = append(resp, seatDTO{ID: id, Status: status})
	}

	json.NewEncoder(w).Encode(resp)
}

func (s *Server) HoldSeat(w http.ResponseWriter, r *http.Request) {
	seatID, _ := strconv.Atoi(r.URL.Query().Get("seat"))
	user := domain.UserID(r.URL.Query().Get("user"))

	err := s.SeatService.HoldSeat(user, seatID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) UnholdSeat(w http.ResponseWriter, r *http.Request) {
	seatID, _ := strconv.Atoi(r.URL.Query().Get("seat"))
	user := domain.UserID(r.URL.Query().Get("user"))

	err := s.SeatService.UnholdSeat(user, seatID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) ConfirmSeat(w http.ResponseWriter, r *http.Request) {
	seatID, _ := strconv.Atoi(r.URL.Query().Get("seat"))
	user := domain.UserID(r.URL.Query().Get("user"))
	bookingID := domain.BookingID(r.URL.Query().Get("booking"))

	err := s.SeatService.ConfirmBooking(seatID, user, bookingID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
}
