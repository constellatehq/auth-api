package model

type Photo struct {
	Id        string `db:"id" json:"id"`
	MediaType string `db:"media_type" json:"media_type"`
	MediaUrl  string `db:"media_link" json:"media_link"`
	Permalink string `db:"permalink" json:"permalink"`
}
