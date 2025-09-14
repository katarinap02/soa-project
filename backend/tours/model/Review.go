package model

import (
	"encoding/json"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TourID      primitive.ObjectID `bson:"tourId,omitempty" json:"tourId"`
	TouristID   primitive.ObjectID `bson:"touristId,omitempty" json:"touristId"`
	Rating      int                `bson:"rating" json:"rating"`
	Comment     string             `bson:"comment" json:"comment"`
	VisitDate   time.Time          `bson:"visitDate" json:"visitDate"`
	CommentDate time.Time          `bson:"commentDate" json:"commentDate"`
	Images      []string           `bson:"images,omitempty" json:"images"`
}

type Reviews []*Review

// JSON serijalizacija
func (r *Review) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(r)
}

func (r *Review) FromJSON(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(r)
}

func (rv *Reviews) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(rv)
}

func (rv *Reviews) FromJSON(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(rv)
}
