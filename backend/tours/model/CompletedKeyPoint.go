package model

import (
	"encoding/json"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompletedKeyPoint struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	TourExecutionID string             `bson:"tourExecutionId" json:"tourExecutionId"`
	KeyPointID      string             `bson:"keyPointId" json:"keyPointId"`
	CompletedTime   time.Time          `bson:"completedTime" json:"completedTime"`
}

func (ckp *CompletedKeyPoint) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(ckp)
}

func (ckp *CompletedKeyPoint) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(ckp)
}
