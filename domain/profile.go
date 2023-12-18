package domain

import "github.com/google/uuid"

type Profile struct {
	ID          uuid.UUID   `json:"id"`
	UserID      uuid.UUID   `json:"user_id"`
	Bio         Bio         `json:"bio"`
	Preferences Preferences `json:"preferences"`
}

type Bio struct {
	AboutMe   string `json:"about_me"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	City      string `json:"city"`
	Latitude  int    `json:"latitude"`
	Longitude int    `json:"longitude"`
}

type Preferences struct {
	Gender   string `json:"gender"`
	AgeMin   int    `json:"age_min"`
	AgeMax   int    `json:"age_max"`
	Distance int    `json:"distance"`
}
