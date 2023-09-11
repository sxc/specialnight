package main

import (
	"context"
	"flag"
	"log"

	"github.com/sxc/specialnight/db"

	"github.com/gofiber/fiber/v2"
	"github.com/sxc/specialnight/api"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"
const dbname = "specialnight"
const userColl = "users"

var config = fiber.Config{
	// Override default error handler
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		// 	// Status code defaults to 500
		// 	code := fiber.StatusInternalServerError

		// 	// Retrieve the custom status code if it's a *fiber.Error
		// 	var e *fiber.Error
		// 	if errors.As(err, &e) {
		// 		code = e.Code
		// 	}

		// 	// Send custom error page
		// 	err = ctx.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
		// 	if err != nil {
		// 		// In case the SendFile fails
		// 		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		// 	}

		// 	// Return from handler
		// 	return nil
		return c.JSON(map[string]string{
			"error": err.Error(),
		})
	},
}

func main() {

	// ctx := context.Background()
	// coll := client.Database(dbname).Collection(userColl)

	// user := types.User{
	// FirstName: "John",
	// LastName:  "Appleseed",
	// }

	// fmt.Println(res)
	// var john types.User
	// if err := coll.FindOne(ctx, bson.M{}).Decode(&john); err != nil {
	// log.Fatal(err)
	// }

	// fmt.Println(john)

	listenAddr := flag.String("listenAddr", ":3000", "The listen address of the API server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))

	if err != nil {
		log.Fatal(err)
	}

	// _, err = coll.InsertOne(ctx, user)
	// if err != nil {
	// log.Fatal(err)
	// }

	// handlers initialization
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)

	app.Listen(*listenAddr)
}
