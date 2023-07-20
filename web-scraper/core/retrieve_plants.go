package core

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	LatinNamePrefix  = "Latin Name:"
	FamilyNamePrefix = "Family Name:"
)

type Plant struct {
	Name        string `json:"name"`
	Safety      string `json:"safety"`
	LatinName   string `json:"latinName"`
	FamilyName  string `json:"familyName"`
	Description string `json:"description"`
}

type SectionPlants struct {
	Section string  `json:"section"`
	Plants  []Plant `json:"plants"`
}

var httpClient = &http.Client{
	Timeout: time.Second * 10,
}

func GetPlantData(plantUrl string) (Plant, error) {
	var plant Plant

	res, err := httpClient.Get(plantUrl)
	if err != nil {
		return plant, err
	}
	defer CloseBody(res.Body)

	if res.StatusCode != 200 {
		return plant, errors.New("status code error")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return plant, err
	}

	plant.Name = doc.Find("h1").First().Text()

	safetyClasses := []string{".greensign", ".redsign", ".orangesign1", ".orangesign2"}
	for _, className := range safetyClasses {
		safetyElement := doc.Find(className)
		if safetyElement.Length() > 0 {
			plant.Safety = strings.TrimSpace(safetyElement.Text())
			break
		}
	}

	doc.Find("#plantinfoouter ul li").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		switch i {
		case 1:
			plant.LatinName = strings.TrimPrefix(text, LatinNamePrefix)
		case 2:
			plant.FamilyName = strings.TrimPrefix(text, FamilyNamePrefix)
		}
	})

	ulSelection := doc.Find("#plantinfoouter ul")
	plant.Description = CleanupDescription(strings.TrimSpace(ulSelection.Parent().Next().Text()))

	return plant, nil
}

func GetPlants(sectionUrl string) ([]Plant, error) {
	var plants []Plant

	res, err := httpClient.Get(sectionUrl)
	if err != nil {
		return plants, err
	}
	defer CloseBody(res.Body)

	if res.StatusCode != 200 {
		return plants, errors.New("status code error")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return plants, err
	}

	doc.Find(".plantbox").Each(func(i int, s *goquery.Selection) {
		plantUrl, _ := s.Find(".dbseemore").First().Attr("href")

		plantData, err := GetPlantData(plantUrl)
		if err != nil {
			return
		}

		plants = append(plants, plantData)
	})

	return plants, nil
}

func Execute() ([]SectionPlants, error) {
	var sectionsPlants []SectionPlants

	res, err := httpClient.Get("https://www.thetortoisetable.org.uk")
	if err != nil {
		return nil, err
	}
	defer CloseBody(res.Body)

	if res.StatusCode != 200 {
		return nil, errors.New("status code error")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	doc.Find(".homepagebox").Each(func(i int, s *goquery.Selection) {
		sectionUrl, _ := s.Find(".boxpic a").First().Attr("href")
		sectionName := strings.ToUpper(s.Find(".boxtitle").First().Text())

		log.Println("Processing section:", sectionName)
		plants, err := GetPlants(sectionUrl)
		if err != nil {
			log.Println("Error getting plants for section", sectionName, ":", err)
			return
		}

		sectionPlants := SectionPlants{
			Section: sectionName,
			Plants:  plants,
		}
		sectionsPlants = append(sectionsPlants, sectionPlants)
	})

	return sectionsPlants, nil
}

func CloseBody(body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		fmt.Println("Error closing the body:", err)
	}
}
