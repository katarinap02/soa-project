package model

import (
	"encoding/json"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Korpa korisnika
type ShoppingCart struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID     string             `bson:"userId" json:"userId"`
	Items      []OrderItem        `bson:"items" json:"items"`
	TotalPrice float64            `bson:"totalPrice" json:"totalPrice"`
}

// Lista korpi
type ShoppingCarts []*ShoppingCart

// Serijalizacija u JSON
func (s *ShoppingCart) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(s)
}

func (s *ShoppingCart) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(s)
}

func (sc *ShoppingCarts) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(sc)
}

func (sc *ShoppingCarts) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(sc)
}
