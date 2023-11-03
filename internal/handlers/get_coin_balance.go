package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Tboules/go_coins_api/api"
	"github.com/Tboules/go_coins_api/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func getCoinBalance(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("calling get coin balance")
	var params = api.CoinBalanceParams{}
	decoder := schema.NewDecoder()
	var err error

	//grab username from request
	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	//instantiate db instance
	db, err := tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	//grab coins from db
	userCoins, coinErr := (*db).GetUserCoins(params.Username)
	if err != nil {
		log.Error(coinErr)
		api.InternalErrorHandler(w)
		return
	}

	//send response
	res := api.CoinBalanceResponse{
		Code:    http.StatusOK,
		Balance: (*userCoins).Coins,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}
