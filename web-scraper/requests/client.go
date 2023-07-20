package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/tonycallaghan/web-scraper/core"
	"net/http"
)

func PostPlant(plant core.Plant) error {
	plantJSON, _ := json.Marshal(plant)

	resp, err := http.Post("http://localhost:8080/plant", "application/json", bytes.NewBuffer(plantJSON))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("status code error")
	}

	return nil
}
