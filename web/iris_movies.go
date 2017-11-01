package web

import (
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/context"
)

type Movie struct {
	Name   string `json:"name"`
	Year   int    `json:'year'`
	Poster string `json:'poster'`
}

var movies = []Movie{
	Movie{
		Name:   "Gone with the wind",
		Year:   2010,
		Poster: "Benjamin",
	},
	Movie{
		Name:   "Gold",
		Year:   2011,
		Poster: "Benjamin",
	},
}

type MovieController struct {
	mvc.C
}

func (c *MovieController) Get() []Movie {
	return movies
}

// demo:
// curl -i -X PUT -F "poster=wendy.c.yan" http://localhost:8080/movies/1
func (c *MovieController) PutBy(id int) Movie {
	movie := movies[id]
	poster := c.Ctx.FormValue("poster")
	movie.Poster = poster
	movies[id] = movie
	return movie
}

// demo
// curl -i -X DELETE http://localhost:8080/movies/1
func (c *MovieController) DeleteBy(id int) context.Map {
	deleted := movies[id].Name
	movies = append(movies[:id], movies[id+1:]...)
	return context.Map{"deleted": deleted}
}
