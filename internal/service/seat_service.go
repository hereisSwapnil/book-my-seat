package service

import (
	"errors"
	"sync"

	"github.com/hereisSwapnil/book-my-seat/internal/domain"
)

type SeatEntry struct {
	seat *domain.Seat
	mu   sync.Mutex
}

type SeatService struct {
	seats map[int]*SeatEntry
}

func NewSeatService(number_of_seats int) *SeatService {
	seats := make(map[int]*SeatEntry)
	for i := 0; i < number_of_seats; i++ {
		seats[i] = &SeatEntry{
			seat: &domain.Seat{
				ID: i,
			},
		}
	}
	return &SeatService{
		seats: seats,
	}
}

func (ss *SeatService) HoldSeat(userID domain.UserID, seatID int) error {
	entry, ok := ss.seats[seatID]
	if !ok {
		return errors.New("invalid seat")
	}

	entry.mu.Lock()
	defer entry.mu.Unlock()

	return entry.seat.Hold(userID)
}

func (ss *SeatService) UnholdSeat(userID domain.UserID, seatID int) error {
	entry, ok := ss.seats[seatID]
	if !ok {
		return errors.New("invalid seat")
	}

	entry.mu.Lock()
	defer entry.mu.Unlock()

	return entry.seat.Unhold(userID)
}

func (ss *SeatService) ConfirmBooking(seatID int, userId domain.UserID, bookingID domain.BookingID) error {
	entry, ok := ss.seats[seatID]
	if !ok {
		return errors.New("invalid seat")
	}

	entry.mu.Lock()
	defer entry.mu.Unlock()

	return entry.seat.Confirm(userId, bookingID)
}

func (ss *SeatService) ListSeats() map[int]domain.SeatStatus {
	seats := make(map[int]domain.SeatStatus)
	for id, entry := range ss.seats {
		entry.mu.Lock()
		defer entry.mu.Unlock()
		seats[id] = entry.seat.Status()
	}
	return seats
}
