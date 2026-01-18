package main

import (
	"fmt"
	"net/http"

	"github.com/hereisSwapnil/book-my-seat/internal/service"
	httptransport "github.com/hereisSwapnil/book-my-seat/internal/transport/http"
)

func main() {
	seatService := service.NewSeatService(100)

	server := &httptransport.Server{
		SeatService: seatService,
	}

	httptransport.RegisterRoutes(server)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", enableCORS(http.DefaultServeMux))
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
