package service

import (
	"beerSockets/web/model"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type BeerService struct {
	saveDir string
}

func (self *BeerService) Save(beer model.Beer) (model.Beer, error) {
	if beer.Id == [16]byte{} {
		beer.Id = uuid.New()
	} else {
		old, err := self.Get(beer.Id)
		if err == nil {
			beer.LastUpdate = time.Now()
			beer.CreationDate = old.CreationDate
		}
	}
	content, err := json.Marshal(beer)
	if err != nil {
		return model.Beer{}, err
	}
	err = SaveFile(self.saveDir, beer.Id.String(), string(content))
	if err != nil {
		return model.Beer{}, err
	}
	return beer, nil
}

func (self *BeerService) GetAll() ([]model.Beer, error) {
	files, err := ReadAllFile(self.saveDir)
	if err != nil {
		return nil, err
	}
	ba := [100]model.Beer{}
	res := ba[:0]
	for _, elem := range files {
		beer := model.Beer{}
		err = json.Unmarshal([]byte(elem.content), &beer)
		if err != nil {
			return nil, err
		}
		res = append(res, beer)
	}
	return res, nil
}

func (self *BeerService) Get(id uuid.UUID) (model.Beer, error) {
	if id == [16]byte{} {
		return model.Beer{}, errors.New("Invalid id")
	}
	res, err := ReadFile(self.saveDir, id.String())
	if err != nil {
		return model.Beer{}, err
	}
	beer := model.Beer{}
	err = json.Unmarshal([]byte(res.content), &beer)
	if err != nil {
		return model.Beer{}, err
	}
	return beer, nil
}

func (self *BeerService) SetSaveDir(dir string) {
	self.saveDir = dir
}
