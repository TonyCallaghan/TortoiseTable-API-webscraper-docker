package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tonycallaghan/web-scraper/core"
)

func PostPlant(plant core.Plant) error {
	plantJSON, _ := json.Marshal(plant)

	resp, err := http.Post("http://api:8080/plant", "application/json", bytes.NewBuffer(plantJSON))

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error reading response body:", err)
		} else {
			bodyString := string(bodyBytes)
			log.Println("Response status code:", resp.StatusCode)
			log.Println("Response body:", bodyString)
		}
		return errors.New("status code error")
	}

	return nil
}
