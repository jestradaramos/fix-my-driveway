package repo

import (
	"github.com/go-pg/pg/v9"
)

func NewRepo() {
	pg.Connect()
}
