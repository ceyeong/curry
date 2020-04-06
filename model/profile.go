package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Profile : User profile
type Profile struct {
	UserID  primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Gender  string             `json:"gender,omitempty" bson:"gender,omitempty"`
	Phone   string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Website string             `json:"website,omitempty" bson:"website,omitempty"`
	Bio     string             `json:"bio,omitempty" bson:"bio,omitempty"`
}

func NewProfile() *Profile {
	return new(Profile)
}
