package utils

import (
	"uk-towns/pkg/entities"
	"uk-towns/pkg/store"

	// "reflect"
	"fmt"
	"time"
	// "sync"
	"encoding/csv"
	"os"
	"strconv"
)

func CsvReadData(fileName string, storage store.Storage) error {
	csvFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, line := range csvLines{
		var latitude float64
		var longitude float64
		latitude, err := strconv.ParseFloat(line[7], 32)
		if err != nil {
			fmt.Println(err)
		}
		longitude, err = strconv.ParseFloat(line[8], 32)
		if err != nil {
			fmt.Println(err)
		}

		town := &entities.Town{
			Name:      line[1],
			Latitude:  latitude,
			Longitude: longitude,
			CreatedAt: time.Now(),
		}

		storage.CreateTown(town)

		continue
	}
	return nil
}