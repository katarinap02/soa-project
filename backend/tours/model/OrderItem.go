package model

import (
	"encoding/json"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Stavka u korpi
type OrderItem struct {
	TourID   primitive.ObjectID `bson:"tourId" json:"tourId"`
	TourName string             `bson:"tourName" json:"tourName"`
	Price    float64            `bson:"price" json:"price"`
}

// Lista stavki
type OrderItems []*OrderItem

// Serijalizacija u JSON
func (o *OrderItem) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(o)
}

func (o *OrderItem) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(o)
}

func (oi *OrderItems) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(oi)
}

func (oi *OrderItems) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(oi)
}
