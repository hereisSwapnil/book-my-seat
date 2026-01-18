package domain

import (
	"errors"
	"time"
)

type Seat struct {
	ID        int
	status    SeatStatus
	heldUntil time.Time
	heldBy    UserID
	bookingID BookingID
}

func (s *Seat) Hold(user UserID) error {
	if s.status != Available {
		return errors.New("seat is not available")
	}

	s.status = Held
	s.heldUntil = time.Now().Add(60 * time.Second)
	s.heldBy = user

	return nil
}

func (s *Seat) Unhold(user UserID) error {
	if s.status == Held && s.heldBy == user {
		s.status = Available
		return nil
	}
	return errors.New("seat not held by you")
}

func (s *Seat) Confirm(user UserID, bookingID BookingID) error {

	if s.status == Booked {
		return errors.New("seat is already booked")
	}

	if s.status != Held || s.heldBy != user {
		return errors.New("seat not held by you")
	}

	s.status = Booked
	s.bookingID = bookingID

	return nil
}

func (s *Seat) Status() SeatStatus {
	return s.status
}
