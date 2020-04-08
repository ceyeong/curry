package utils

import (
	cctx "github.com/ceyeong/curry/context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetObjectID : returns primitive.ObjectID form of passed string
func GetObjectID(str string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(str)
}

// GetUserObjectIDFromSession : returns user id object from session
func GetUserObjectIDFromSession(c echo.Context) (primitive.ObjectID, error) {
	cc := c.(*cctx.CurryContext)

	uID, err := cc.GetFromSession("userID")
	if err != nil {
		return primitive.NewObjectID(), err
	}
	return GetObjectID(uID.(string))
}
