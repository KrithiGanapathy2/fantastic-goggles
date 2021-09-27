package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"os"
)

type Author struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

type Book struct {
	Name string `json:"name"`
	Genre string `json:"genre"`
	Author Author `json:"author"`
}

func jsonify(b []Book, f string) {
	var raw, _ = json.Marshal(b)
	var file, _ = os.Create("data.json")
	file.WriteString(string(raw))
	file.Close()

}

func getJson(f string) []Book {
	var data, _ = os.ReadFile("data.json")
	var reading []Book
	json.Unmarshal(data, &reading)
	return reading
}

func getRawJson(f string) string {
	var data, _ = os.ReadFile(f)
	return string(data)
}

func main() {
	var port = os.Getenv("PORT")
	var app = fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("It works")
	})
	app.Get("/api/v1", func(c *fiber.Ctx) error {
		return c.SendString(getRawJson("data.json"))
	})
	app.Listen(":"+port)
}