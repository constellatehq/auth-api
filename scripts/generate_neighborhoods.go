package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/constellatehq/auth-api/model"
)

const (
	GET_NEW_YORK_NEIGHBORHOODS_URL = "https://data.cityofnewyork.us/resource/xyye-rtrs.json"
)

type Geom struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type NeighborhoodResponseObject struct {
	TheGeom   Geom   `json:"the_geom"`
	ObjectId  int    `json:"objectid"`
	Name      string `json:"name"`
	Stacked   string `json:"stacked"`
	Annoline1 string `json:"annoline1"`
	Annoline2 string `json:"annoline2"`
	Annoline3 string `json:"annoline3"`
	Annoangle string `json:"annoangle"`
	Borough   string `json:"borough"`
}

func main() {
	// addNeighborhoodsToDB()

}

func getNeighborhoods() ([]NeighborhoodResponseObject, error) {
	neighborhoodResponseObject := []NeighborhoodResponseObject{}
	response, err := Get(GET_NEW_YORK_NEIGHBORHOODS_URL)
	if err != nil {
		return nil, fmt.Errorf("Failed to get neighborhoods: %s", err.Error())
	}
	json.Unmarshal(response, &neighborhoodResponseObject)
	return neighborhoodResponseObject, nil
}

func addNeighborhoodsToDB() error {
	response, err := getNeighborhoods()
	if err != nil {
		return err
	}

	for index, district := range response {
		index += 1
		neighborhood := &model.Neighborhood{}
		neighborhood.City = district.Borough
		neighborhood.District = district.Name
		neighborhood.State = "New York"
		neighborhood.Country = "United States"
		coordinates := district.TheGeom.Coordinates
		neighborhood.Latitude = coordinates[0]
		neighborhood.Longitude = coordinates[1]
		fmt.Printf("%+v\n", neighborhood)
	}

	return nil
}

func ResponseToByte(response *http.Response) ([]byte, error) {
	defer response.Body.Close()
	buf := &bytes.Buffer{}
	_, err := io.Copy(buf, response.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed reading response body: %s", err.Error())
	}

	return buf.Bytes(), nil
}

func Get(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Get request failed: %s", err.Error())
	}

	body, err := ResponseToByte(response)

	if response.StatusCode != 200 {
		return body, fmt.Errorf("%s", response.Status)
	}

	return body, err
}
