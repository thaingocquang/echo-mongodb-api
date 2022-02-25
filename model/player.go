package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// PlayerBSON ...
	PlayerBSON struct {
		ID        primitive.ObjectID `bson:"_id"`
		Name      string             `bson:"name"`
		Email     string             `bson:"email"`
		Password  string             `bson:"password"`
		CreatedAt time.Time          `bson:"createdAt"`
		UpdatedAt time.Time          `bson:"updatedAt"`
	}

	// PlayerDetail ...
	PlayerDetail struct {
		ID        primitive.ObjectID `bson:"_id" json:"_id"`
		Name      string             `json:"name"`
		Email     string             `json:"email"`
		Password  string             `json:"password"`
		CreatedAt time.Time          `json:"createdAt"`
		UpdatedAt time.Time          `json:"updatedAt"`
	}
)
