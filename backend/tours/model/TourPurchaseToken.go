package model

import (
	"encoding/json"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TourPurchaseToken struct {
	TourID primitive.ObjectID `bson:"tourId" json:"tourId"`
	Token  string             `bson:"token" json:"token"`
	UserID string             `bson:"userId" json:"userId"`
}


// Lista tokena
type TourPurchaseTokens []*TourPurchaseToken

// Serijalizacija u JSON
func (t *TourPurchaseToken) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(t)
}

func (t *TourPurchaseToken) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(t)
}

func (tt *TourPurchaseTokens) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(tt)
}

func (tt *TourPurchaseTokens) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(tt)
}
