package main

import (
	"fmt"
	"time"
	"uk-towns/pkg/entities"
	"uk-towns/pkg/store"
	"uk-towns/pkg/utils"
)

func main() {
	db := store.DataBaseConnect()
	townStore := store.NewStorage(db)

	err := utils.CsvReadData("uk-towns-sample.csv", townStore)
	if err != nil {
		fmt.Println(err)
	}

	town := &entities.Town{
		Name:        "Foo",
		Latitude:    0.1,
		Longitude:   0.1,
		CreatedAt: time.Now(),
	}
	town2 := &entities.Town{
		Name:        "Bar",
		Latitude:    0.2,
		Longitude:   0.2,
		CreatedAt: time.Now(),
	}
	err = townStore.CreateTown(town)
	if err != nil {
		fmt.Println(err)
	}
	err = townStore.CreateTown(town2)
	if err != nil {
		fmt.Println(err)
	}

	journey := &entities.Journey{
		Name:       "FooBar",
		FromTownID: town.ID,
		ToTownID:   town2.ID,
		Distance:   9.0,
		CreatedAt:  time.Now(),
	}
	err = townStore.CreateJourney(journey)
	if err != nil {
		fmt.Println(err)
	}
    outputJourney, err := townStore.GetJourney()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", outputJourney)
}