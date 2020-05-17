package handler

import (
	"context"
	"net/http"
	"time"

	cctx "github.com/ceyeong/curry/context"
	"github.com/ceyeong/curry/database"
	"github.com/ceyeong/curry/model"
	"github.com/ceyeong/curry/utils"
	"github.com/gookit/validate"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// RegisterUser : POST /register
func RegisterUser(c echo.Context) error {

	u := model.NewUser()
	if err := c.Bind(u); err != nil {
		return err
	}
	v := validate.Struct(u)

	v.AddRule("name", "required")
	v.StringRule("password", "required|string|minLen:6|maxLen:25")
	v.StringRule("email", "required|email")
	if !v.Validate() {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": v.Errors})
	}

	// check if user already exists
	collection := database.GetDatabase().Collection("user")
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

	//hash user password
	user.HashPassword()

	// insert user to collection
	res, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		return err
	}

	user.ID = res.InsertedID.(primitive.ObjectID)
	return c.JSON(http.StatusOK, &user)
}

// LoginUser : POST /login
func LoginUser(c echo.Context) error {
	u := model.NewUser()
	if err := c.Bind(u); err != nil {
		return err
	}

	v := validate.Struct(u)
	v.StringRule("password", "required|string|minLen:6|maxLen:25")
	v.StringRule("email", "required|email")

	if !v.Validate() {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": v.Errors})
	}

	safeData := v.SafeData()
	email := safeData["email"].(string)
	password := safeData["password"].(string)

	collection := database.GetDatabase().Collection("user")

	user := model.NewUser()
	//search for user via email; if not found return error
	if err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user); err != nil {
		return echo.ErrUnauthorized
	}
	if err := user.ComparePassword(password); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "user or password doesn't match"})
	}
	cc := c.(*cctx.CurryContext)
	err := cc.SetToSession("userID", user.ID.Hex())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "success"})
}

// LogoutUser : POST /logout
func LogoutUser(c echo.Context) error {
	cc := c.(*cctx.CurryContext)
	err := cc.ClearSession()
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

// Token : POST /auth/token
func Token(c echo.Context) error {
	if err := LoginUser(c); err != nil {
		return err
	}
	token, err := utils.GenerateJWTTokens(c.Get("user").(string))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "success", "token": token})
}

// Me : GET /me
func Me(c echo.Context) error {
	cc := c.(*cctx.CurryContext)

	uID, err := cc.GetFromSession("userID")
	if err != nil {
		return err
	}
	objectID, err := primitive.ObjectIDFromHex(uID.(string))
	if err != nil {
		return err
	}
	collection := database.GetDatabase().Collection("user")
	user := new(model.User)
	if err := collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&user); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// Csrf : GET /csrf
func Csrf(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
