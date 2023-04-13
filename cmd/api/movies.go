package main

import (
	"net/http"
	"swagger-intro/internal/data"
	"swagger-intro/internal/validator"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateMovieHandler godoc
//	@Summary	Creates movie
//	@Produce	json
//	@Success	200
//	@Router		/v1/movies [post]
func (app *application) createMovieHandler(c *gin.Context) {
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}
	err := app.readJSON(c.Writer, c.Request, &input)
	if err != nil {
		app.badRequestResponse(c.Writer, c.Request, err)
		return
	}

	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}

	v := validator.New()

	if data.ValidateMovie(v, movie); !v.Valid() {
		app.failedValidationResponse(c.Writer, c.Request, v.Errors)
		return
	}

	var output struct {
		Title  string `json:"title"`
		Action string `json:"action"`
	}
	output.Title = input.Title
	output.Action = "create success"

	err = app.writeJSON(c.Writer, http.StatusOK, envelope{"movie": output}, nil)
	if err != nil {
		app.serverErrorResponse(c.Writer, c.Request, err)
	}
}

// ShowMovieHandler godoc
//	@Summary	Retrieves movie
//	@Produce	json
//	@Param		id	path	integer	true	"Movie ID"
//	@Success	200
//	@Router		/v1/movies/{id} [get]
func (app *application) showMovieHandler(c *gin.Context) {
	id := c.Param("id")

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}
	err := app.writeJSON(c.Writer, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(c.Writer, c.Request, err)
	}
}
