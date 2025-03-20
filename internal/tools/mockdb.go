package tools

import "time"

type mockDB struct {
}

// mock data

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "1234",
		Username:  "alex",
	},
	"john": {
		AuthToken: "12345",
		Username:  "john",
	},
	"maria": {
		AuthToken: "11111",
		Username:  "maria",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    20,
		Username: "alex_1",
	},
	"john": {
		Coins:    100,
		Username: "john_1",
	},
	"maria": {
		Coins:    50,
		Username: "maria_1",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)
	clientData := LoginDetails{}
	clientData, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)
	clientData := CoinDetails{}
	clientData, ok := mockCoinDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetUpDatabase() error {
	return nil
}
