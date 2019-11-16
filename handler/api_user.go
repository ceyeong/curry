package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/ceyeong/curry/database"
	"github.com/ceyeong/curry/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gookit/validate"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// RegisterUser : POST /register
func RegisterUser(c echo.Context) error {

	u := echo.Map{}
	if err := c.Bind(&u); err != nil {
		return err
	}
	v := validate.Map(u)

	v.AddRule("name", "required")
	v.StringRule("password", "required|string|minLen:6|maxLen:25")
	v.StringRule("email", "required|email")
	if !v.Validate() {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": v.Errors})
	}

	// check if user already exists
	collection := database.Database.Collection("user")
	user := new(model.User)
	safeUser := v.SafeData()
	err := collection.FindOne(context.TODO(), bson.M{"email": safeUser["email"]}).Decode(&user)

	// if decode successfull than user already exist
	if err == nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "user already exists"})
	}
	// if error is other than not found than return that error
	if err != mongo.ErrNoDocuments {
		return err
	}
	user.Name = safeUser["name"].(string)
	user.Email = safeUser["email"].(string)
	user.Password = safeUser["password"].(string)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.PasswordUpdatedAt = time.Now()

	// insert user to collection
	res, err := collection.InsertOne(context.TODO(), u)

	if err != nil {
		return err
	}

	user.ID = res.InsertedID.(primitive.ObjectID)
	return c.JSON(http.StatusOK, &user)
}

// LoginUser : POST /login
func LoginUser(c echo.Context) error {
	return nil
}

// Me : GET /me
func Me(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	exp := claims["exp"].(string)

	//parse expiry time and compare it.
	expTime, _ := time.Parse(time.RFC3339, exp)
	if expTime.Before(time.Now()) {
		return echo.ErrUnauthorized
	}

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	collection := database.Database.Collection("user")

	user := new(model.User)
	if err := collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&user); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
