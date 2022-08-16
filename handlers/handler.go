package handlers

import (
	"context"
	"fmt"
	m "github.com/marincor/lab_firebase_datarealtime/models"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

func ListAll(c echo.Context) error {
	ctx := context.Background()
	opt := option.WithCredentialsFile("credentials_firebase.json")
	config := &firebase.Config{
		DatabaseURL: "https://lab-firebase-dc9ec-default-rtdb.firebaseio.com",
	}

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatal(err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatal(err)
	}

	var players []m.Player

	client.NewRef("teams").Get(ctx, &players)

	fmt.Println(players)

	return c.JSON(http.StatusOK, players)
}

func SearchOne(c echo.Context) error {
	ctx := context.Background()
	opt := option.WithCredentialsFile("credentials_firebase.json")
	config := &firebase.Config{
		DatabaseURL: "https://lab-firebase-dc9ec-default-rtdb.firebaseio.com",
	}

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatal(err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatal(err)
	}

	var player m.Player
	id := c.Param("id")

	client.NewRef("teams"+"/"+id).Get(ctx, &player)

	fmt.Println(player)

	return c.JSON(http.StatusOK, player)
}

func CreateOne(c echo.Context) error {
	player := new(m.Player)

	if err := c.Bind(player); err != nil {
		return err
	}

	ctx := context.Background()
	opt := option.WithCredentialsFile("credentials_firebase.json")
	config := &firebase.Config{
		DatabaseURL: "https://lab-firebase-dc9ec-default-rtdb.firebaseio.com",
	}

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatal(err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatal(err)
	}

	client.NewRef("teams/3").Set(ctx, player)

	return c.JSON(http.StatusOK, player)
}
