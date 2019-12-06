package app

import (
	"fix-my-driveway/repo"
	"net/url"
)

// App this is the thing we run
type App struct {
	// Currently nothing
}

// NewApp This will create our app with all dependencies and blah
func NewApp()(*App, error){
	app := &App{}
	dbString := "postgres://postgres:docker@localhost:5432/postgres"
	dbURL, err := url.Parse(dbString)
	if err != nil {
		return nil, err
	}

	db, err := repo.InitDatabase(dbURL)
	if err != nil {
		return nil, err
	}
	db.InitUserTable()

	return app, nil
}


