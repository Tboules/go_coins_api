package tools

import (
	log "github.com/sirupsen/logrus"
)

// db collections
type LoginDetails struct {
	Username string
	Token    string
}

type CoinDetails struct {
	Coins    int64
	Username string
}

type DbInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) (*CoinDetails, error)
	SetupDb() error
}

func NewDatabase() (*DbInterface, error) {
	var db DbInterface = &mockDB{}
	var err error = db.SetupDb()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &db, nil
}
