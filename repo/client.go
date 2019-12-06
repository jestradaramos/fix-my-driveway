package repo


import (
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"fix-my-driveway/domain"
)

// Repository yup
type Repository struct {
	db *pg.DB
}


// NewRepo will be used to create a database client
func NewRepo() Repository {

	// Figure out the options
	db := pg.Connect(&pg.Options{
        User: "postgres",
	})
	
	defer db.Close()

	err := createSchema(db)
	if err != nil {
		panic(err)
	}

	return Repository {db}
}

func createSchema(db *pg.DB) error {
    for _, model := range []interface{}{(*domain.User)(nil),(*domain.TimeEntry)(nil)} {
        err := db.CreateTable(model, &orm.CreateTableOptions{
            Temp: true,
        })
        if err != nil {
            return err
        }
    }
    return nil
}