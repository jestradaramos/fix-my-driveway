package main
import (
	"fmt"
	"fix-my-driveway/app"
)

func main() {
	a, err := app.NewApp()
	if err !=nil {
		panic(err)
	}
	fmt.Print(a)
}