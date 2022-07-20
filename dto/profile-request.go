package dto

import "time"

type ProfileRequest struct {
	WantedJobTitle   string    `json:"wantedJobTitle"`
	FirstName        string    `json:firstName`
	LastName         string    `json:"lastName"`
	Email            string    `json:"email"`
	Phone            string    `json:"phone"`
	PlaceOfBirth     string    `json:"placeOfBirth"`
	DateOfBirth      time.Time `json:"dateOfBirth"`
	City             string    `json:"city"`
	PostalCode       int64     `json:"postalCode"`
	DrivingLicenseNo string    `json:"drivingLicenseNo"`
	Nationality      string    `json:"nationality"`
	Country          string    `json:"country"`
	Address          string    `json:"address"`
}
