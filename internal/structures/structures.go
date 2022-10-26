package structures

import "time"

type User struct {
	Id         int `json:"id"`
	Amount     int `json:"amount"`
	Bookamount int `json:"bookamount"`
}

type Order struct {
	Id          int       `json:"id"`
	Fromuserid  int       `json:"fromuserid"`
	Touserid    int       `json:"touserid"`
	Serviceid   int       `json:"serviceid"`
	Orderid     int       `json:"orderid"`
	Amount      int       `json:"amount"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}
