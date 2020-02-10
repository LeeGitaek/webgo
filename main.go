package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.POST("/save", save)

	e.Logger.Fatal(e.Start(":1323")) // localhost:1323

}

func getUser(c echo.Context) error {
	// localhost:1323/users/gitaeklee
	// output ( method : GET )
	// gitaeklee
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
func show(c echo.Context) error {
	// localhost:1323/show?team=divtag&member=gitaeklee
	// output ( method : GET )
	// team : divtag, member : gitaeklee
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team :"+team+",member :"+member)
}

func save(c echo.Context) error {
	// localhost:1323/save
	// output ( method : POST )
	// name :Joe Smith, email :joe@labstack.com
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name :"+name+", email :"+email+"\n")
}
