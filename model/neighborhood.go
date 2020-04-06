package model

type Neighborhood struct {
	Id        string  `json:"id"`
	Country   string  `json:"country"`
	State     string  `json:"state"`
	City      string  `json:"city"`
	District  string  `json:"district"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	ZipCode   string  `json:"zip_code"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
