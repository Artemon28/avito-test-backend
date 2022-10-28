package internal

type transferRequest struct {
	Fromuserid int `json:"fromuserid"`
	Touserid   int `json:"touserid"`
	Serviceid  int `json:"serviceid"`
	Orderid    int `json:"orderid"`
	Amount     int `json:"amount"`
}

type bookRequest struct {
	Id     int `json:"id"`
	Amount int `json:"amount"`
}
