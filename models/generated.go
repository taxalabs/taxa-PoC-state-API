// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"github.com/99designs/gqlgen/graphql"
)

type AddressInput struct {
	PostalCode   string  `json:"postalCode"`
	PostalTown   string  `json:"postalTown"`
	BuildingName string  `json:"buildingName"`
	StreetName   string  `json:"streetName"`
	Lon          float64 `json:"lon"`
	Lat          float64 `json:"lat"`
	RestaurantID string  `json:"restaurantID"`
}

type RestaurantInput struct {
	RestaurantName string `json:"restaurantName"`
	About          string `json:"about"`
	Telephone      string `json:"telephone"`
}

type UploadLicense struct {
	RestaurantID string         `json:"restaurantID"`
	File         graphql.Upload `json:"file"`
}
