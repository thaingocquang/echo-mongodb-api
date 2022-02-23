package service

import (
	"echo-mongodb-api/dao"
	"echo-mongodb-api/model"
	"errors"
)

// PlayerCreate ...
func PlayerCreate(body model.PlayerCreateBody) (model.PlayerBSON, error) {
	var (
		player = body.ConvertToBSON()
	)

	//Create user
	doc, err := dao.PlayerCreate(player)
	if err != nil {
		err = errors.New("can not create user")
		return doc, err
	}

	return doc, err
}

func PlayerProfileByEmail(email string) (model.PlayerProfile, error) {
	doc, err := dao.PlayerProfileByEmail(email)
	return doc, err
}
