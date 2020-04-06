package handler

import (
	"context"
	"net/http"

	"github.com/ceyeong/curry/database"
	"github.com/ceyeong/curry/model"
	"github.com/ceyeong/curry/utils"
	"github.com/gookit/validate"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// GetProfile : GET /me/profile
func GetProfile(c echo.Context) error {
	userID, err := utils.GetUserObjectIDFromSession(c)
	if err != nil {
		return err
	}
	collection := database.GetDatabase().Collection("user")
	profile := new(model.Profile)

	if err := collection.FindOne(context.TODO(), bson.M{"user_id": userID}).Decode(profile); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, profile)
}

// PostProfile : POST /me/profile
func PostProfile(c echo.Context) error {
	userID, err := utils.GetUserObjectIDFromSession(c)
	if err != nil {
		return err
	}
	p := model.NewProfile()
	if err := c.Bind(p); err != nil {
		return err
	}
	v := validate.Struct(p)
	v.AddRule("gender", "in:m,f")
	v.AddRule("phone", "minLen:10|maxLen13")
	v.AddRule("bio", "maxLen:100")
	v.AddRule("website", "url")
	if !v.Validate() {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": v.Errors})
	}
	collection := database.GetDatabase().Collection("profile")
	p.UserID = userID
	_, err = collection.UpdateOne(context.TODO(), bson.M{"user_id": userID}, p)

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "success", "data": p})
}
