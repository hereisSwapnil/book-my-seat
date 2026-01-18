package httptransport

import "net/http"

func RegisterRoutes(server *Server) {
	http.HandleFunc("/seats", server.ListSeats)
	http.HandleFunc("/hold", server.HoldSeat)
	http.HandleFunc("/unhold", server.UnholdSeat)
	http.HandleFunc("/confirm", server.ConfirmSeat)
}
