package util

import (
	"context"
	"echo-mongodb-api/model"
	"echo-mongodb-api/module/database"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	// PlayerIDString for test
	PlayerIDString = "5f24d45125ea51bc57a8285a"

	// PlayerID for test
	PlayerID, _ = primitive.ObjectIDFromHex(PlayerIDString)

	// Player for test
	Player = model.PlayerBSON{
		ID:        PlayerID,
		Name:      "tnq",
		Email:     "tnq@gmail.com",
		Password:  "123456",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
)

// HelperPlayerCreateFake ...
func HelperPlayerCreateFake() {
	var (
		playerCol = database.PlayerCol()
		ctx       = context.Background()
	)

	Player.Password = HashPassword(Player.Password)

	//Insert
	_, err := playerCol.InsertOne(ctx, Player)
	if err != nil {
		log.Println(err)
	}
}

// ClearDB ...
func ClearDB() {
	database.PlayerCol().DeleteMany(context.Background(), bson.M{})
}
