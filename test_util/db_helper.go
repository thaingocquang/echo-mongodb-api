package testutil

import (
	"context"
	"echo-mongodb-api/config"
	"echo-mongodb-api/model"
	"echo-mongodb-api/module/database"
	"echo-mongodb-api/util"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// ConnectTestDB ...
func ConnectTestDB() {
	envVars := config.GetEnv()

	// Connect
	cl, err := mongo.NewClient(options.Client().ApplyURI(envVars.Database.URI))
	if err != nil {
		log.Println(err)
		log.Fatal("Cannot connect to database:", envVars.Database.URI)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = cl.Connect(ctx); err != nil {
		log.Println(err)
	}
	db := cl.Database(envVars.Database.TestName)
	database.SetDB(db)
	fmt.Println("Database Connected to", envVars.Database.TestName)
}

// ClearDB ...
func ClearDB() {
	database.PlayerCol().DeleteMany(context.Background(), bson.M{})
}

// HelperPlayerCreateFake ...
func HelperPlayerCreateFake() {
	var (
		playerCol = database.PlayerCol()
		ctx       = context.Background()
	)

	Player.Password = util.HashPassword(Player.Password)

	//Insert
	_, err := playerCol.InsertOne(ctx, Player)
	if err != nil {
		log.Println(err)
	}
}
