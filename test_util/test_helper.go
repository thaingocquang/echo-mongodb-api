package testutil

import (
	"bytes"
	"context"
	"echo-mongodb-api/config"
	"echo-mongodb-api/model"
	"echo-mongodb-api/module/database"
	"echo-mongodb-api/route"
	"echo-mongodb-api/util"
	"encoding/json"
	"io"
	"log"
	"time"

	"github.com/labstack/echo/v4"
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

	Player.Password = util.HashPassword(Player.Password)

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

func InitServer() *echo.Echo {
	config.Init()
	database.Connect()
	ClearDB()

	// new server
	e := echo.New()
	route.Route(e)

	return e
}

func BodyToReader(i interface{}) io.Reader {
	bodyMarshal, _ := json.Marshal(i)
	return bytes.NewReader(bodyMarshal)
}
