package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"database-example/model"
	"database-example/service"
)

type KeyTour struct{}

type ToursHandler struct {
	logger  *log.Logger
	service *service.TourService
}

// Konstruktor za ToursHandler
func NewToursHandler(l *log.Logger, s *service.TourService) *ToursHandler {
	return &ToursHandler{logger: l, service: s}
}

func (h *ToursHandler) MiddlewareTourDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Čitaj telo zahteva
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read request body", http.StatusBadRequest)
			h.logger.Printf("Error reading body: %v", err)
			return
		}

		// Loguj telo u konzolu
		h.logger.Printf("Received JSON body: %s", string(bodyBytes))

		// Vrati body da može da se ponovo koristi
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// Deserijalizacija u Tour
		tour := &model.Tour{}
		if err := json.Unmarshal(bodyBytes, tour); err != nil {
			var msg string

			// Detaljnija obrada grešaka
			switch e := err.(type) {
			case *json.SyntaxError:
				msg = fmt.Sprintf("JSON syntax error at byte offset %d: %v", e.Offset, e.Error())
			case *json.UnmarshalTypeError:
				msg = fmt.Sprintf("JSON type error: field '%s', expected %v but got %v at offset %d",
					e.Field, e.Type, e.Value, e.Offset)
			default:
				msg = fmt.Sprintf("Unable to decode JSON: %v", err)
			}

			http.Error(w, msg, http.StatusBadRequest)
			h.logger.Println(msg)
			return
		}

		// Dodaj tour u kontekst
		ctx := context.WithValue(r.Context(), KeyTour{}, tour)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *ToursHandler) GetAllTours(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("INFO: GetAllTours handler hit!") // ADD THIS LINE
	tours, err := h.service.GetAllTours(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logger.Printf("Error getting all tours: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tours); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		h.logger.Printf("Error encoding tours response: %v", err)
	}
}

func (h *ToursHandler) CreateTour(w http.ResponseWriter, r *http.Request) {

	_, span := tp.Tracer(serviceName).Start(r.Context(), "CreateTour")
	defer span.End()

	tour, ok := r.Context().Value(KeyTour{}).(*model.Tour)
	if !ok {
		http.Error(w, "Tour not found in context", http.StatusInternalServerError)
		h.logger.Println("Tour object not found in context for creation")
		span.RecordError(fmt.Errorf("tour object not found in context"))
		return
	}

	// Uzmi authorID iz tela (već je deserializovan u tour)
	if tour.AuthorID == "" {
		http.Error(w, "AuthorID not provided in tour body", http.StatusBadRequest)
		h.logger.Println("AuthorID missing in tour creation request body")
		span.RecordError(fmt.Errorf("authorID missing"))
		return
	}

	span.AddEvent("Creating tour in service")
	if err := h.service.CreateTour(r.Context(), tour, tour.AuthorID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logger.Printf("Error creating tour: %v", err)
		span.RecordError(err)
		return
	}
	span.AddEvent("Tour created successfully")

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Tour created successfully"})
}

// Ostale metode: GetTourByID, UpdateTour, DeleteTour...
func (h *ToursHandler) GetTourByID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("Not Implemented"))
}
func (h *ToursHandler) UpdateTour(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("Not Implemented"))
}
func (h *ToursHandler) DeleteTour(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("Not Implemented"))
}

func (h *ToursHandler) GetToursByAuthor(w http.ResponseWriter, r *http.Request) {
	authorID := r.URL.Query().Get("authorId")
	if authorID == "" {
		http.Error(w, "authorId is required", http.StatusBadRequest)
		return
	}

	tours, err := h.service.GetToursByAuthor(r.Context(), authorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logger.Printf("Error fetching tours by author: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if tours == nil {
		tours = []*model.Tour{} // vrati prazan niz umesto null
	}
	json.NewEncoder(w).Encode(tours)
}
