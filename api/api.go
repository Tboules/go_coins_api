package api

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	Code    int
	Balance int
}

type Error struct {
	Code    int
	Message string
}
