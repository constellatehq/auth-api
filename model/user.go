package model

import (
	"fmt"
	"time"
)

type Gender int

const (
	NullGender Gender = iota
	Male
	Female
	Other
)

func (g Gender) String() string {
	return [...]string{"", "male", "female", "other"}[g]
}

type JobType int

const (
	NullJobType JobType = iota
	Standard
	WorkFromHome
	SelfEmployed
)

func (j JobType) String() string {
	return [...]string{"", "standard", "work_from_home", "self_employed"}[j]
}

type User struct {
	Id              string     `db:"id" json:"id"`
	FacebookId      *string    `db:"facebook_id" json:"facebook_id,omitempty"`
	GoogleId        *string    `db:"google_id" json:"google_id,omitempty"`
	InstagramId     *string    `db:"instagram_id" json:"instagram_id,omitempty"`
	SpotifyId       *string    `db:"spotify_id" json:"spotify_id,omitempty"`
	FirstName       string     `db:"first_name" json:"first_name"`
	LastName        string     `db:"last_name" json:"last_name"`
	Email           string     `db:"email" json:"email"`
	Birthday        *time.Time `db:"birthday" json:"birthday,omitempty"`
	Gender          int        `db:"gender" json:"gender,omitempty"` // Enum Gender
	Onboarded       bool       `db:"onboarded" json:"onboarded"`
	PermissionLevel int        `db:"permission_level" json:"permission_level"`
	EmailVerified   bool       `db:"email_verified" json:"email_verified"`
	CreatedAt       *time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt       *time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetFacebookId() string {
	if m != nil {
		return *m.FacebookId
	}
	return ""
}

func (m *User) GetGoogleId() string {
	if m != nil {
		return *m.GoogleId
	}
	return ""
}

func (m *User) GetInstagramId() string {
	if m != nil {
		return *m.InstagramId
	}
	return ""
}

func (m *User) GetSpotifyId() string {
	if m != nil {
		return *m.SpotifyId
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetBirthday() time.Time {
	if m != nil {
		return *m.Birthday
	}
	return time.Time{}
}

func (m *User) GetGender() string {
	if m != nil {
		return Gender(m.Gender).String()
	}
	return ""
}

func (m *User) GetOnboarded() bool {
	if m != nil {
		return m.Onboarded
	}
	return false
}

func (m *User) GetPermissionLevel() int {
	if m != nil {
		return m.PermissionLevel
	}
	return 0
}

func (m *User) GetEmailVerified() bool {
	if m != nil {
		return m.EmailVerified
	}
	return false
}

func (m *User) GetCreatedAt() time.Time {
	if m != nil {
		return *m.CreatedAt
	}
	return time.Time{}
}

func (m *User) GetUpdatedAt() time.Time {
	if m != nil {
		return *m.UpdatedAt
	}
	return time.Time{}
}

func (user *User) IsValid() *ErrorResponse {
	// if len(user.Id) != 26 {
	// 	return InvalidUserError("id", user.Id)
	// }
	if user.FirstName == "" {
		return InvalidUserError("first_name", user.FirstName)
	}
	if user.LastName == "" {
		return InvalidUserError("last_name", user.LastName)
	}
	if user.Email == "" {
		return InvalidUserError("email", user.Email)
	}

	return nil
}

func InvalidUserError(fieldName string, userId string) *ErrorResponse {
	errorMessage := fmt.Sprintf("Unable to validate %s for user %s", fieldName, userId)

	return NewErrorResponse("Invalid User Error", errorMessage, nil)
}

// Relation types

type UserPreference struct {
	Id     string `db:"id" json:"id"`
	UserId string `db:"user_id" json:"user_id"`
	// Apartment info
	Budget   float32    `db:"budget" json:"budget"`
	MoveIn   *time.Time `db:"move_in" json:"move_in"`
	Duration string     `db:"duration" json:"duration"`
	// Lifestyle info
	JobType           int    `db:"job_type" json:"job_type"` // Enum - standard (9 to 5), wfh, self_employed
	JobTitle          string `db:"job_title" json:"job_title"`
	SleepWakeSchedule int    `db:"sleep_wake_schedule" json:"sleep_wake_schedule"` // Enum - early_bird, night_owl
	CleaningSchedule  int    `db:"cleaning_schedule" json:"cleaning_schedule"`     // Enum - weekly, bimonthly, monthly

	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

func (m *UserPreference) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserPreference) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserPreference) GetBudget() float32 {
	if m != nil {
		return m.Budget
	}
	return 0
}

func (m *UserPreference) GetMoveIn() time.Time {
	if m != nil {
		return *m.MoveIn
	}
	return time.Time{}
}

func (m *UserPreference) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

func (m *UserPreference) GetJobType() int {
	if m != nil {
		return m.JobType
	}
	return 0
}

func (m *UserPreference) GetJobTitle() string {
	if m != nil {
		return m.JobTitle
	}
	return ""
}

func (m *UserPreference) GetCreatedAt() time.Time {
	if m != nil {
		return *m.CreatedAt
	}
	return time.Time{}
}

func (m *UserPreference) GetUpdatedAt() time.Time {
	if m != nil {
		return *m.UpdatedAt
	}
	return time.Time{}
}

func (m *UserPreference) IsValid(fieldName string, userPreferenceId string) *ErrorResponse {
	if len(m.Id) != 26 {
		return InvalidUserError("id", m.Id)
	}
	if m.UserId == "" {
		return InvalidUserError("first_name", m.UserId)
	}
	if m.Budget < 25.0 {
		return InvalidUserError("budget", fmt.Sprintf("%f", m.Budget))
	}
	if m.Duration == "" {
		return InvalidUserError("duration", m.Duration)
	}

	return nil
}

func InvalidUserPreferenceError(fieldName string, userPreferenceId string) *ErrorResponse {
	errorMessage := fmt.Sprintf("Unable to validate %s for user %s", fieldName, userPreferenceId)

	return NewErrorResponse("Invalid User Error", errorMessage, nil)
}

type UserRoommatePreference struct {
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

type UserPhoto struct {
	Id     string  `db:"id" json:"id"`
	UserId string  `db:"user_id" json:"user_id"`
	Photos []Photo `db:"photos" json:"photos"`
}

type UserNeighborhood struct {
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
