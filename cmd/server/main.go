package main

import (
	"fmt"

	"github.com/jlkwarteng/comments-app/internal/comment"
	db "github.com/jlkwarteng/comments-app/internal/database"
	transportHttp "github.com/jlkwarteng/comments-app/internal/transport/http"
)

// Run - Is going to be responsible for starting the application
func Run() error {
	fmt.Println("Starting the Application")
	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to Connect to Database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("Failed to Migrate Database")
		return err
	}
	fmt.Println("Successfully Connected and Pinged Database ")
	cmtService := comment.NewService(db)
	// cmtService.Store.CreateComment(context.Background(), comment.Comment{Id: "e4de69eb-366b-4d68-b61a-95dad6fdf3a7", Body: "My Test Body", Slug: "Tesst Slug", Author: "Jlkwarteng"})

	// fmt.Println(cmtService.Store.GetComment(context.Background(), "3bfb3afc-9e63-4002-a085-9a40c8a4aeec"))
	// err = cmtService.DeleteComment(context.Background(), "e4de69eb-366b-4d68-b61a-95dad6fdf3a7")
	// if err != nil {
	// 	fmt.Errorf("Failed to Delete Comment", err)
	// }
	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}
	return nil

}

func main() {
	fmt.Println("Rest Api")
	err := Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
