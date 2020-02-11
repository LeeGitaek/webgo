package main

import (
	"database/sql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"webgo/handlers"
	//"github.com/graphql-go/graphql"
	//"encoding/json"
)

func main() {
	//schema, _ := graphql.NewSchema(graphql.SchemaConfig{})

	db := initDB("storage.db")
	migrate(db)

	// new instance.
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.File("/", "public/index.html")
	e.GET("/tasks", handlers.GetTasks(db))
	e.PUT("/tasks", handlers.PutTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))

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

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil { // if err is not nil
		panic(err)
	}

	if db == nil { // if db is nil
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `CREATE TABLE IF NOT EXISTS tasks(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name VARCHAR NOT NULL );`

	_, err := db.Exec(sql)

	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}

}
