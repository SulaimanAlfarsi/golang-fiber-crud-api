package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies = []Movie{
	{ID: "1", Isbn: "123456789", Title: "Inception", Director: &Director{Firstname: "Christopher", Lastname: "Nolan"}},
	{ID: "2", Isbn: "987654321", Title: "The Matrix", Director: &Director{Firstname: "Wachowski", Lastname: "Brothers"}},
	{ID: "3", Isbn: "12344321", Title: "Deadpool", Director: &Director{Firstname: "John", Lastname: "Wicker"}},
}

// Handler for the root route
func Welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to the Movie API!")
}

// This route retrieves all movies.
func getMovies(c *fiber.Ctx) error {
	return c.JSON(movies)
}

// Handler to get a specific movie by ID
func getMovie(c *fiber.Ctx) error {
	id := c.Params("id")

	for _, movie := range movies {
		if movie.ID == id {
			return c.JSON(movie)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Movie not found")

}

// Handler to create a new movie
func createMovie(c *fiber.Ctx) error {
	var movie Movie

	if err := c.BodyParser(&movie); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	movie.ID = fmt.Sprintf("%d", len(movies)+1)

	movies = append(movies, movie)

	return c.Status(fiber.StatusCreated).JSON(movie)
}

// Handler to update a movie by ID
func updateMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	var updatedMovie Movie

	if err := c.BodyParser(&updatedMovie); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	for i, movie := range movies {
		if movie.ID == id {
			movies[i].Isbn = updatedMovie.Isbn
			movies[i].Title = updatedMovie.Title
			movies[i].Director = updatedMovie.Director
			return c.JSON(movies[i])
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Movie not found")
}

// Handler to delete a movie by ID
func deleteMovie(c *fiber.Ctx) error {
	id := c.Params("id")

	for i, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:i], movies[i+1:]...)
			return c.SendString("Movie deleted successfully")
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Movie not found")
}

func main() {
	app := fiber.New()
	app.Get("/", Welcome)
	app.Get("/movies", getMovies)
	app.Get("/movies/:id", getMovie)
	app.Post("/movies", createMovie)
	app.Put("/movies/:id", updateMovie)
	app.Delete("/movies/:id", deleteMovie)

	fmt.Println("Starting server on :8080...")

	log.Fatal(app.Listen(":8080"))
}
