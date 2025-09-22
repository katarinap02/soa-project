package model

import (
	"encoding/json"
	"io"
	"time"
	 "go.mongodb.org/mongo-driver/bson/primitive"
)

type Tour struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description,omitempty" json:"description"`
	Price       float64            `bson:"price,omitempty" json:"price"`
	Difficulty  string             `bson:"difficulty,omitempty" json:"difficulty"`
	Tags        []string           `bson:"tags,omitempty" json:"tags"`
	Status      string             `bson:"status,omitempty" json:"status"`
	AuthorID    string             `bson:"authorId,omitempty" json:"authorId"`
	ArchiveDate *time.Time `bson:"archiveDate,omitempty" json:"archiveDate,omitempty"`
	PublishDate *time.Time `bson:"publishDate,omitempty" json:"publishDate,omitempty"`
	Durations map[string]int `bson:"durations,omitempty"` // walking, bicycle, car
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
