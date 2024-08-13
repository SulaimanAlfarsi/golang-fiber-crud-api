package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
)


type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	 Firstname string `json:"firstname"`
	 Lastname string `json:"lastname"`
}

var movies = []Movie{
	{ID: "1", Isbn: "123456789", Title: "Inception", Director: &Director{Firstname: "Christopher", Lastname: "Nolan"}},
	{ID: "2", Isbn: "987654321", Title: "The Matrix", Director: &Director{Firstname: "Wachowski", Lastname: "Brothers"}},
	{ID: "3", Isbn: "12344321", Title: "Deadpool", Director: &Director{Firstname: "John", Lastname: "Wicker"}},

}

// Handler for the root route
func Welcome(c *fiber.Ctx) error{
	return c.SendString("Welcome to the Movie API!")
}

//This route retrieves all movies.
func getMovies(c *fiber.Ctx) error {
	return c.JSON(movies)
}

// Handler to get a specific movie by ID
func getMovie(c *fiber.Ctx) error {
	id:= c.Params("id") // Extract the ID from the URL parameters

	for _, movie := range movies {
        if movie.ID == id {
            return c.JSON(movie)
        }
    }
	return c.Status(fiber.StatusNotFound).SendString("Movie not found")

}

// Handler to update a movie by ID



func main() {
	app := fiber.New()
	app.Get("/",Welcome)
	app.Get("/movies", getMovies)
	app.Get("/movies/:id", getMovie) 
	// app.Post("/movies", createMovie)
	// app.Put("/movies/{id}", updateMovie)
	// app.Delete("/movies/{id}", deleteMovie)

	fmt.Println("Starting server on :8080...")

	log.Fatal(app.Listen(":8080"))
}
