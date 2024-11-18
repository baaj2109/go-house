package model

type HouseInfo struct {
	Preview    string `json:"preview"`
	Title      string `json:"title"`
	URL        string `json:"url"`
	Address    string `json:"address"`
	RentType   string `json:"rentType"`
	OptionType string `json:"optionType"`
	Ping       string `json:"ping"` // a.k.a 坪數
	Floor      string `json:"floor"`
	Price      string `json:"price"`
	IsNew      bool   `json:"isNew"`
}

func NewHouseInfo() *HouseInfo {
	return &HouseInfo{}
}