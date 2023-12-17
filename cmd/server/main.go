package main

import "fmt"

// Run - Is going to be responsible for starting the application
func Run() error {
	fmt.Println("Starting the Application")
	return nil

}

func main() {
	fmt.Println("Rest Api")
	err := Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
