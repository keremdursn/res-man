package utils

import (
	"context"
	"golang-restaurant-management/database"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	jwt.RegisteredClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func init() {
	if SECRET_KEY == "" {
		log.Fatal("SECRET_KEY environment değişkeni tanımlı değil!")
	}
}

func GenerateAllTokens(email string, firstName string, lastName string, uid string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email:      email,
		First_name: firstName,
		Last_name:  lastName,
		Uid:        uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	refreshClaims := &SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(168 * time.Hour)),
		},
	}

	token, tokenErr := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, refreshErr := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if tokenErr != nil || refreshErr != nil {
		log.Panicf("Token üretme hatası: tokenErr=%v, refreshErr=%v", tokenErr, refreshErr)
	}

	return token, refreshToken, nil
}

func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	updateObj := bson.D{
		{"token", signedToken},
		{"refresh_token", signedRefreshToken},
		{"updated_at", time.Now()},
	}

	filter := bson.M{"user_id": userId}
	upsert := true
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := userCollection.UpdateOne(
		ctx,
		filter,
		bson.D{{"$set", updateObj}},
		&opt,
	)

	if err != nil {
		log.Panic(err)
	}
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
		jwt.WithLeeway(5*time.Second), // tolerans süresi
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "token claim dönüşümü başarısız"
		return
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		msg = "token is expired"
		return
	}

	return claims, ""
}
