package model

import (
	"encoding/json"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TourExecution struct {
	ID           primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
	TourID       string              `bson:"tourId" json:"tourId"`
	TouristID    string              `bson:"touristId" json:"touristId"`
	LastActivity time.Time           `bson:"lastActivity" json:"lastActivity"`
	Status       TourExecutionStatus `bson:"status" json:"status"`
}

type TourExecutions []*TourExecution

func (te *TourExecution) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(te)
}

func (te *TourExecution) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(te)
}

func (te *TourExecutions) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(te)
}

func (te *TourExecutions) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(te)
}
