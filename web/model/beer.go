package model

import (
	"github.com/google/uuid"
	"time"
)

type Beer struct {
	Id           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	BeerType     string    `json:"beerType"`
	CreationDate time.Time `json:"creationDate"`
	LastUpdate   time.Time `json:"lastUpdate"`
}
