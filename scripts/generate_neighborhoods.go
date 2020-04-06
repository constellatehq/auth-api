// package scripts

// import (
// 	"github.com/constellatehq/auth-api/model"
// )

// const (
// 	NEW_YORK_NEIGHBORHOODS_URL = "https://data.cityofnewyork.us/resource/xyye-rtrs.json"
// )

// type Geom struct {
// 	Type        string    `json:"type"`
// 	Coordinates []float64 `json:"coordinates"`
// }

// type NeighborhoodResponse struct {
// 	TheGeom   Geom   `json:"the_geom"`
// 	ObjectId  int    `json:"objectid"`
// 	Name      string `json:"name"`
// 	Stacked   string `json:"stacked"`
// 	Annoline1 string `json:"annoline1"`
// 	Annoline2 string `json:"annoline2"`
// 	Annoline3 string `json:"annoline3"`
// 	Annoangle string `json:"annoangle"`
// 	Borough   string `json:"borough"`
// }

// func getNeighborhoods() (model.Response, error) {

// }

// func addNeighborhoodsToDB() {
// 	response, err := getNeighborhoods()
// 	if err != nil {
// 		return nil, fmt.Sprintf("Failed to get neighborhoods:%s", err.Error())
// 	}

// 	neighborhoods := []model.Neighborhood

// 	for district := range response {
// 		neighborhood := &model.Neighborhood{}
// 		neighborhood.City = district["borough"]
// 		neighborhood.District = district["name"]
// 		neighborhood.State = "New York"
// 		neighborhood.Country = "United States"
// 		coordinates := district["the_geom"]["coordinates"]
// 		neighborhood.Latitude = coordinates[0]
// 		neighborhood.Longitude = coordinates[1]
// 	}

// }

