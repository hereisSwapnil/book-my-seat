# Book My Seat - Low Level Design (LLD)

## Project Overview
Book My Seat is a simple, concurrent seat booking system implemented in Go. It demonstrates core backend concepts including Clean Architecture, concurrent state management with Mutexes, and RESTful API design. The frontend is a lightweight HTML/JS application that visualizes real-time seat availability.

## Architecture
The project follows **Clean Architecture** principles to ensure separation of concerns:

- **Domain Layer (`internal/domain`)**: Contains core entities (`Seat`, `SeatStatus`) and business logic. It has no external dependencies.
- **Service Layer (`internal/service`)**: Orchestrates business use cases (`SeatService`). Handles concurrency and transaction-like operations.
- **Transport Layer (`internal/transport/http`)**: Manages HTTP handlers, routing, and request/response marshaling.
- **Frontend**: A vanilla JavaScript/HTML client consuming the backend APIs.

## Low Level Design (LLD)

### 1. Entities & Data Structures
**Seat Entity**
- `ID`: Unique identifier for the seat.
- `Status`: generic state (`Available`, `Held`, `Booked`).
- `HeldBy`: User ID holding the seat.
- `HeldUntil`: Expiration time for the hold.
- `BookingID`: Confirmation reference after booking.

**Seat Service**
- Uses `sync.RWMutex` to ensure thread-safe access to the shared `seats` slice.
- `seats`: In-memory storage of all `Seat` objects.

### 2. Concurrency Model
The system is designed to handle multiple concurrent reservation requests:
- **Locking**: A global `sync.RWMutex` in `SeatService` protects the seat inventory.
  - `RLock()` is used for reading seat status (ListSeats).
  - `Lock()` is used for state-modifying operations (Hold, Unhold, Confirm).
- **Optimistic/Pessimistic Locking**: The system effectively uses pessimistic locking at the service level (mutex) to prevent race conditions (e.g., two users booking the same seat).

### 3. API Specification

| Method | Endpoint    | Query Params                | Description                          |
|--------|-------------|-----------------------------|--------------------------------------|
| GET    | `/seats`    | -                           | Returns list of all seats with status|
| GET    | `/hold`     | `seat={id}`, `user={id}`    | Temporarily holds a seat             |
| GET    | `/unhold`   | `seat={id}`, `user={id}`    | Releases a held seat                 |
| GET    | `/confirm`  | `seat={id}`, `user={id}`, `booking={id}` | Confirms booking for a held seat |

### 4. State Machine (Seat Status)
A generic state machine governs seat transitions:
1. **Available** → **Held** (via `HoldSeat`)
2. **Held** → **Booked** (via `ConfirmBooking` if User matches)
3. **Held** → **Available** (via `UnholdSeat` or Timeout)
4. **Booked** → *Final State* (No further transitions in this MVP)

## Setup & Running

### Prerequisites
- Go 1.20+
- Git

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/hereisSwapnil/book-my-seat.git
   cd book-my-seat
   ```

2. Run the server:
   ```bash
   go run cmd/server/main.go
   ```

3. Open the application:
   Visit `http://localhost:8080/index.html` in your browser. Note: You may need to serve the frontend file or simply open it directly if CORS is enabled, though serving via the Go server is recommended for production.

   *Currently, `index.html` can be opened directly in a browser as the server supports CORS.*

## Technologies
- **Language**: Go (Golang)
- **Transport**: Standard `net/http`
- **Frontend**: HTML5, CSS3, Vanilla JavaScript (ES6+)
