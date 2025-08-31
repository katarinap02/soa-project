package model

import (
	"encoding/json"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tour struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description,omitempty" json:"description"`
	Price       float64            `bson:"price,omitempty" json:"price"`
}

type Tours []*Tour

// Serijalizacija u JSON
func (t *Tour) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(t)
}

func (t *Tour) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(t)
}

func (t *Tours) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(t)
}

func (t *Tours) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(t)
}
