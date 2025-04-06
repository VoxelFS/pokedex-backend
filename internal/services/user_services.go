package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type User struct {
	ID             string `json:"id,omitempty" bson:"_id,omitempty"`
	User           string `json:"user,omitempty" bson:"user,omitempty"`
	HashedPassword string `json:"hashed_password,omitempty" bson:"hashed_password,omitempty"`
	SessionToken   string `json:"session_token,omitempty" bson:"session_token,omitempty"`
	CSRFToken      string `json:"csrf_token,omitempty" bson:"csrf_token,omitempty"`
}

// GetUser fetches a user from the DB by their username.
// Returns a User object and an error (if any).
func (u *User) GetUser(username string) (User, error) {
	collection := returnCollectionPointer("user")

	var user User

	err := collection.FindOne(context.Background(), bson.M{"user": username}).Decode(&user)
	if err != nil {
		log.Println(err)
		return User{}, err
	}
	return user, nil
}

// SetSessionToken updates a user's session and CSRF tokens in the DB.
func (u *User) SetSessionToken(sessionToken string, csrfToken string, username string) error {
	collection := returnCollectionPointer("user")
	update := bson.M{
		"$set": bson.M{
			"session_token": sessionToken,
			"csrf_token":    csrfToken,
		},
	}
	_, err := collection.UpdateOne(context.Background(), bson.M{"user": username}, update)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// ClearTokens clears session and CSRF tokens for a given user.
func (u *User) ClearTokens(username string) error {
	collection := returnCollectionPointer("user")
	update := bson.M{
		"$set": bson.M{
			"session_token": "",
			"csrf_token":    "",
		},
	}
	_, err := collection.UpdateOne(context.Background(), bson.M{"user": username}, update)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
