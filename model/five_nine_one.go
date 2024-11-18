package model

import (
	"net/http"
	"sync"
)

// FiveN1 is the representation page information.
type FiveN1 struct {
	records    int
	TotalPages int
	queryURL   string
	isExecuted bool
	RentList   HouseInfoCollection
	wg         sync.WaitGroup
	rw         sync.RWMutex
	client     *http.Client
	cookie     *http.Cookie
}

type HouseInfoCollection map[int][]*HouseInfo

// NewFiveN1 create a `FiveN1` with default value.
func NewFiveN1(url string) *FiveN1 {
	return &FiveN1{
		queryURL: url,
		RentList: make(map[int][]*HouseInfo),
		client:   &http.Client{},
		cookie: &http.Cookie{
			Name:  "urlJumpIp",
			Value: "1",
		},
	}
}
