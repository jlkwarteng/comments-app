package main

import (
	"context"
	"fmt"

	db "github.com/jlkwarteng/comments-app/internal/database"
)

// Run - Is going to be responsible for starting the application
func Run() error {
	fmt.Println("Starting the Application")
	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to Connect to Database")
		return err
	}

	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	fmt.Println("Successfully Connected and Pinged Database ")
	return nil

}

func main() {
	fmt.Println("Rest Api")
	err := Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
