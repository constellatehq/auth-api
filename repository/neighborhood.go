package repository

import (
	"fmt"

	"github.com/constellatehq/auth-api/model"
	gsb "github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

type NeighborhoodInterface interface {
}
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

func CreateNeighborhood(db *sqlx.DB, neighborhood model.Neighborhood) error {

	ib := gsb.PostgreSQL.NewInsertBuilder()

	ib.InsertInto("neighborhoods")
	ib.Cols("country", "state", "city", "district", "latitude", "longitude", "zip_code")
	ib.Values(neighborhood.Country, neighborhood.State, neighborhood.City, neighborhood.District, neighborhood.Latitude, neighborhood.Longitude, neighborhood.ZipCode)

	// Execute the query.
	sql, args := ib.Build()
	fmt.Printf("%s\n%s\n", sql, args)

	tx := db.MustBegin()
	tx.MustExec(sql, args...)
	tx.Commit()
	// return

	return nil
}
