package service

import (
	"beerSockets/web/model"

	"github.com/google/uuid"
)

type Service interface {
	Save(beer model.Beer) (model.Beer, error)
	GetAll() ([]model.Beer, error)
	Get(id uuid.UUID) (model.Beer, error)
}
