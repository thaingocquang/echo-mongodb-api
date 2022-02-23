package dao

import (
	"context"
	"echo-mongodb-api/module/database"

	"go.mongodb.org/mongo-driver/bson"
)

func CheckEmailExisted(email string) (bool, error) {
	var (
		userCol = database.PlayerCol()
		ctx     = context.Background()
	)

	count, err := userCol.CountDocuments(ctx, bson.M{"email": email})
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
