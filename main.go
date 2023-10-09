package main

import (
	"beerSockets/server"
	"beerSockets/web/controller"
	"beerSockets/web/model"
	"log"
	"time"

	"github.com/google/uuid"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Println(`
	|==================|
	|   Beer Sockets   |
	|  Program Started |
	|==================|
	`)

	serverConfig := server.ServerConfig{
		Host: "0.0.0.0",
		Port: "8080",
	}
	controllerConfig := controller.Config{
		SavePath: "/home/pop/code/temp/",
	}
	server := server.New(serverConfig, controllerConfig)
	log.Fatal(server.Run())
}

func testBeerController() {
	conf := controller.Config{}
	conf.SavePath = "/home/pop/code/temp/"
	c := controller.NewBeerController(conf)
	date := time.Now()
	id, err := uuid.Parse("68f325c2-bc09-4ba6-a204-f67da6e02ea9")
	handleError(err)
	beer := model.Beer{
		Id:           id,
		Name:         "Sagres",
		BeerType:     "LAGGER",
		LastUpdate:   date,
		CreationDate: date,
	}
	log.Println(beer)
	log.Println("########## SAVE BEER ##########")
	res1, err := c.SaveBeer(beer)
	handleError(err)
	log.Println(res1)

	log.Println("########## GET BEER ##########")
	res2, err := c.GetBeer(res1)
	handleError(err)
	log.Println(res2)

	log.Println("########## GET ALL BEERS ##########")
	res3, err := c.GetAllBeers()
	handleError(err)
	for _, elem := range res3 {
		log.Println(elem)
	}
}
