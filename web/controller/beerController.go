package controller

import (
	"beerSockets/service"
	"beerSockets/web/model"

	"github.com/google/uuid"
)

type Config struct {
	SavePath string
}

type BeerController struct {
	beerService service.Service
}

func NewBeerController(config Config) *BeerController {
	bs := service.BeerService{}
	bs.SetSaveDir(config.SavePath)
	return &BeerController{
		beerService: &bs,
	}
}

func (self *BeerController) SaveBeer(beer model.Beer) (string, error) {
	record, err := self.beerService.Save(beer)
	return record.Id.String(), err
}

func (self *BeerController) GetBeer(id string) (model.Beer, error) {
	guid, err := uuid.Parse(id)
	if err != nil {
		return model.Beer{}, err
	}
	return self.beerService.Get(guid)
}

func (self *BeerController) GetAllBeers() ([]model.Beer, error) {
	return self.beerService.GetAll()
}
