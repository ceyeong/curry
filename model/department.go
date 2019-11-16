package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Department :
type Department struct {
	ID      primitive.ObjectID `json:"id,omitempty"`
	Name    string             `json:"name,omitempty"`
	Code    string             `json:"code,omitempty"`
	Head    User               `json:"head,omitempty"`
	Members []User             `json:"-"`
}
