package main

import (
	"github.com/tonycallaghan/web-scraper/core"
	"github.com/tonycallaghan/web-scraper/requests"
	"log"
)

func main() {
	sectionsPlants, err := core.Execute()
	if err != nil {
		log.Fatalf("Failed to execute core: %v", err)
	}

	for _, sectionPlants := range sectionsPlants {
		for _, plant := range sectionPlants.Plants {
			err := requests.PostPlant(plant)
			if err != nil {
				log.Fatalf("Failed to POST plant data: %v", err)
			}
		}
	}
}
