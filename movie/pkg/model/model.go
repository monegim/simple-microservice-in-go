package model

import model "github.com/monegim/simple-microservice-in-go/metadata/pkg"

type MovieDetails struct {
	Rating   *float64       `json:"rating,omitempty"`
	Metadata model.Metadata `json:"metadata"`
}
