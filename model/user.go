package model

import (
	"time"
)

type User struct {
	Id              string    `db:"id" json:"id"`
	FacebookId      string    `db:"facebook_id" json:"facebook_id"`
	GoogleId        string    `db:"google_id" json:"google_id"`
	InstagramId     string    `db:"instagram_id" json:"instagram_id"`
	SpotifyId       string    `db:"spotify_id" json:"spotify_id"`
	FirstName       string    `db:"first_name" json:"first_name"`
	LastName        string    `db:"last_name" json:"last_name"`
	Email           string    `db:"email" json:"email"`
	Birthday        time.Time `db:"birthday" json:"birthday"`
	Gender          string    `db:"gender" json:"gender"` // Enum - Male, Female, Other
	Onboarded       bool      `db:"onboarded" json:"onboarded"`
	PermissionLevel int64     `db:"permission_level" json:"permission_level"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
}

// Relation types

type UserPreferences struct {
	Id     string `db:"id" json:"id"`
	UserId string `db:"user_id" json:"user_id"`
	// Apartment info
	Budget   string    `db:"budget" json:"budget"`
	MoveIn   time.Time `db:"move_in" json:"move_in"`
	Duration string    `db:"duration" json:"duration"`
	RoomType string    `db:"room_type" json:"room_type"` // Enum - master, flex, normal
	// Lifestyle info
	JobType           string `db:"job_type" json:"job_type"` // Enum - standard (9 to 5), wfh, self_employed
	JobTitle          string `db:"job_title" json:"job_title"`
	SleepWakeSchedule string `db:"sleep_wake_schedule" json:"sleep_wake_schedule"` // Enum - early_bird, night_owl
	CleaningSchedule  string `db:"cleaning_schedule" json:"cleaning_schedule"`     // Enum - weekly, bimonthly, monthly

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type UserRoommatePreferences struct {
	Id                   string `db:"id" json:"id"`
	UserId               string `db:"user_id" json:"user_id"`
	RoommateRelationship string `db:"roommate_relationship" json:"roommate_relationship"` // Enum - coexistence, friendship, mutual
	RoommateAmount       string `db:"roommate_amount" json:"roommate_amount"`
	Gender               string `db:"gender" json:"gender"`
	Age                  string `db:"age" json:"age"`
	JobType              string `db:"job_type" json:"job_type"`
	SleepWakeSchedule    string `db:"sleep_wake_schedule" json:"sleep_wake_schedule"` // Enum - early_bird, night_owl
	CleaningSchedule     string `db:"cleaning_schedule" json:"cleaning_schedule"`     // Enum - weekly, bimonthly, monthly

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type UserPhotos struct {
	Id     string  `db:"id" json:"id"`
	UserId string  `db:"user_id" json:"user_id"`
	Photos []Photo `db:"photos" json:"photos"`
}

type UserNeighborhoods struct {
	Id            string         `db:"id" json:"id"`
	UserId        string         `db:"user_id" json:"user_id"`
	Neighborhoods []Neighborhood `db:"neighborhoods" json:"neighborhoods"`
}

type UserMusic struct {
	Id         string   `db:"id" json:"id"`
	UserId     string   `db:"user_id" json:"user_id"`
	TopTracks  []Track  `db:"top_tracks" json:"top_tracks"`
	TopArtists []Artist `db:"top_artists" json:"top_artists"`
}
