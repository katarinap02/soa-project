package model

import (
	"encoding/json"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type KeyPoint struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TourID      primitive.ObjectID `bson:"tourId" json:"tourId"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Latitude    float64            `bson:"latitude" json:"latitude"`
	Longitude   float64            `bson:"longitude" json:"longitude"`
	ImageURL    string             `bson:"imageUrl" json:"imageUrl"`
}

type KeyPoints []*KeyPoint

func (kp *KeyPoint) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(kp)
}

func (kp *KeyPoint) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(kp)
}

func (kps *KeyPoints) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(kps)
}

func (kps *KeyPoints) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(kps)
}
