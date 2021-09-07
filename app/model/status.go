package model

import "math/rand"

type StatusHealthcheck uint8

func (h StatusHealthcheck) String() string {
	return healthcheckTbl[h]
}

const (
	StatusOK StatusHealthcheck = iota + 1
	StatusAltered
	StatusDown
)

var healthcheckTbl = [...]string{
	0:             "unknown",
	StatusOK:      "ok",
	StatusAltered: "altered",
	StatusDown:    "down",
}

func HealthcheckStatusRandom() string {
	res := rand.Intn(len(healthcheckTbl))
	if res == 0 {
		res++
	}
	return healthcheckTbl[res]
}
