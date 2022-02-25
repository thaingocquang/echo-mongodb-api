package service

import (
	"echo-mongodb-api/dao"
	"echo-mongodb-api/model"
)

func PlayerProfileFindByID(ID string) (model.PlayerProfile, error) {
	doc, err := dao.PlayerProfileFindByID(ID)
	return doc, err
}
