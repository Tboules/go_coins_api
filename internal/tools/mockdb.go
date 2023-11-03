package tools

import (
	"errors"
	"time"
)

type mockDB struct {
}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		Token:    "123ABC",
		Username: "alex",
	},
	"jason": {
		Token:    "456DEF",
		Username: "jason",
	},
	"marie": {
		Token:    "789GHI",
		Username: "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"jason": {
		Coins:    200,
		Username: "jason",
	},
	"marie": {
		Coins:    250,
		Username: "marie",
	},
}

func (*mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	data, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}

	return &data
}

func (*mockDB) GetUserCoins(username string) (*CoinDetails, error) {
	time.Sleep(time.Second * 1)

	data, ok := mockCoinDetails[username]
	if !ok {
		return nil, errors.New("No coins could be found for that user")
	}

	return &data, nil

}

func (*mockDB) SetupDb() error {
	return nil
}
