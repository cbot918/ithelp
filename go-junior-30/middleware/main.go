package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

const port = ":8886"

var userCount int
var users []User

var postCount int
var posts []Post

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	app.Post("/signup", Signup)
	app.Post("/signin", Signin)

	api := app.Group("/api", RequireAuth())
	api.Post("/post", Postt)

	if err := app.Listen(port); err != nil {
		log.Fatal(err)
	}
}

// handle request
func Signup(c *fiber.Ctx) error {
	user := User{}
	err := c.BodyParser(&user)
	if err != nil {
		return err
	}
	userCount += 1
	user.Id = userCount
	users = append(users, user)
	jsonFormat, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(jsonFormat))
	return c.JSON(user)
}

func Signin(c *fiber.Ctx) error {
	user := User{}
	err := c.BodyParser(&user)
	if err != nil {
		return err
	}

	loginSuccess := false

	for _, item := range users {
		if user.Email == item.Email {
			if user.Password == item.Password {
				loginSuccess = true
				user.Id = item.Id
			}
		}
	}
	var token string
	if loginSuccess {
		token, err = genJWT(user.Id, user.Email)
		if err != nil {
			return err
		}
		fmt.Println(token)
	} else {
		return c.JSON(fiber.Map{"message": "invalid email or password"})
	}

	return c.JSON(fiber.Map{
		"id":    user.Id,
		"token": token,
	})
}

func Postt(c *fiber.Ctx) error {
	post := Post{}
	err := c.BodyParser(&post)
	if err != nil {
		return err
	}
	fmt.Println(post)
	token := c.Get("Authorization")

	claim, err := DecodeJWT(token)
	if err != nil {
		return c.JSON(fiber.Map{"message": "token invalid"})
	}

	authSuccess := false
	for _, item := range users {
		if item.Id == claim.Id {
			authSuccess = true
		}
	}

	var resultPost Post
	if authSuccess {
		postCount += 1
		resultPost = Post{
			Id:       postCount,
			Title:    post.Title,
			Body:     post.Body,
			PostedBy: claim.Id,
		}
		posts = append(posts, resultPost)
	}

	return c.JSON(fiber.Map{
		"post_id":   resultPost.Id,
		"title":     resultPost.Title,
		"body":      resultPost.Body,
		"posted_by": resultPost.PostedBy,
	})
}

// types
type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Post struct {
	Id       int    `json:"post_id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	PostedBy int
}
