package main

import (
	"time"
)

type Agent struct {
	ID          int       `json:"id"`
	UUID        string    `json:"uuid"`
	OS          string    `json:"os"`
	Location    string    `json:"location"`
	CreatedDate time.Time `json:"createddate"`
}

type Agents []Agent
