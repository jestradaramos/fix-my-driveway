package app

import (
	"fix-my-driveway/domain"
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

	db, err := repo.InitDatabase(dbURL)
	if err != nil {
		return nil, err
	}

	// Init all the tables
	db.InitUserTable()

	// Testing Delete later
	u := &domain.User{
		Name: "Jeff",
		Email: "email",
		Password: "pwd",
	}
	err = db.AddUser(u)
	if err != nil {
		return nil, err
	}
	fmt.Print("WE HERE")
	return app, nil
}


