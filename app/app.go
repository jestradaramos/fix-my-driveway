package app

import (
	"fix-my-driveway/repo"
	"fmt"
	"net/url"
)

// App this is the thing we run
type App struct {
	// Currently nothing
}

// NewApp This will create our app with all dependencies and blah
func NewApp()(*App, error){
	app := &App{}
	dbString := "postgres://postgres:docker@database:5432/postgres?sslmode=disable"
	dbURL, err := url.Parse(dbString)
	if err != nil {
		return nil, err
	}
	fmt.Print("WE PARSED THE URL!")
	db, err := repo.InitDatabase(dbURL)
	if err != nil {
		return nil, err
	}
	fmt.Print("WE INIT THE DB")
	db.InitUserTable()
	fmt.Print("IT WORKED!")
	return app, nil
}


