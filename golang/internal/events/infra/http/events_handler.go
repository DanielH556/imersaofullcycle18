package http

import (
	"encoding/json"
	"net/http"

	"github.com/DanielH556/imersaofullcycle18/golang/internal/events/usecase"
)

type EventsHandler struct {
	listEventsUseCase  *usecase.ListEventsUseCase
	getEventUseCase    *usecase.GetEventUseCase
	createEventUseCase *usecase.CreateEventUseCase
	buyTicketsUseCase  *usecase.BuyTicketsUseCase
	createSpotsUseCase *usecase.CreateSpotsUseCase
	listSpotsUseCase   *usecase.ListSpotsUseCase
}

func NewEventsHandler(
	listEventsUseCase *usecase.ListEventsUseCase,
	getEventUseCase *usecase.GetEventUseCase,
	createEventUseCase *usecase.CreateEventUseCase,
	buyTicketsUseCase *usecase.BuyTicketsUseCase,
	createSpotsUseCase *usecase.CreateSpotsUseCase,
	listSpotsUseCase *usecase.ListSpotsUseCase,
) *EventsHandler {
	return &EventsHandler{
		listEventsUseCase:  listEventsUseCase,
		getEventUseCase:    getEventUseCase,
		createEventUseCase: createEventUseCase,
		buyTicketsUseCase:  buyTicketsUseCase,
		createSpotsUseCase: createSpotsUseCase,
		listSpotsUseCase:   listSpotsUseCase,
	}
}

func (h *EventsHandler) ListEvents(w http.ResponseWriter, r *http.Request) {
	output, err := h.listEventsUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (h *EventsHandler) GetEvent(w http.ResponseWriter, r *http.Request) {
	eventID := r.PathValue("eventID")
	input := usecase.GetEventInputDTO{ID: eventID}

	output, err := h.getEventUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (h *EventsHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateEventInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.createEventUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (h *EventsHandler) ListSpots(w http.ResponseWriter, r *http.Request) {
	eventID := r.PathValue("eventID")
	input := usecase.ListSpotsInputDTO{EventID: eventID}

	output, err := h.listSpotsUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (h *EventsHandler) BuyTickets(w http.ResponseWriter, r *http.Request) {
	var input usecase.BuyTicketsInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.buyTicketsUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (h *EventsHandler) CreateSpots(w http.ResponseWriter, r *http.Request) {
	eventID := r.PathValue("eventID")
	var input usecase.CreateSpotsInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	input.EventID = eventID

	output, err := h.createSpotsUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

// writeErrorResponse writes an error response in JSON format
func (h *EventsHandler) writeErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Message: message})
}

// ErrorResponse represents the structure of an error response
type ErrorResponse struct {
	Message string `json:"message"`
}

// CreateSpotsRequest represents the input for creating spots.
type CreateSpotsRequest struct {
	NumberOfSpots int `json:"number_of_spots"`
}
