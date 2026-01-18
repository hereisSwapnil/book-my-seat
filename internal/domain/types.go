package domain

type UserID string
type BookingID string

type SeatStatus int

const (
	Available SeatStatus = iota
	Held
	Booked
)
