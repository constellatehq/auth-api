package model

type Photo struct {
	Id  string `db:"id" json:"id"`
	Url string `db:"url" json:"url"`
}
