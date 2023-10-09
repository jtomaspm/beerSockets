package controller

import (
	"beerSockets/web/model"
	"encoding/json"
)

type ControllerWrapper struct {
	beerController *BeerController
}

func NewControllerWrapper(config Config) *ControllerWrapper {
	return &ControllerWrapper{
		beerController: NewBeerController(config),
	}
}

func (self *ControllerWrapper) handleError(err error) []byte {
	return []byte("ERROR\n" + err.Error())
}

func (self *ControllerWrapper) SAVE(payload []byte) []byte {
	beer := model.Beer{}
	err := json.Unmarshal(payload, &beer)
	if err != nil {
		return self.handleError(err)
	}
	res, err := self.beerController.SaveBeer(beer)
	if err != nil {
		return self.handleError(err)
	}
	return []byte("SUCCESS\n" + res)
}

func (self *ControllerWrapper) GET(payload []byte) []byte {
	id := string(payload)
	res, err := self.beerController.GetBeer(id)
	if err != nil {
		return self.handleError(err)
	}
	json, err := json.Marshal(res)
	if err != nil {
		return self.handleError(err)
	}
	return json
}

func (self *ControllerWrapper) GETALL(_ []byte) []byte {
	res, err := self.beerController.GetAllBeers()
	if err != nil {
		return self.handleError(err)
	}
	json, err := json.Marshal(res)
	if err != nil {
		return self.handleError(err)
	}
	return json
}
