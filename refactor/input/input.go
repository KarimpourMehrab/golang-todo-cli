package input

import (
	"bufio"
	"fmt"
	"todo/refactor/handler"
	"todo/refactor/helper"
)

func Handler(command string, scanner *bufio.Scanner) {
	//flag.Parse()
	fmt.Printf("command :")
	scanner.Scan()
	command = scanner.Text()
	command = helper.RemoveSpace(command)

	if handler.UserIsAuthenticated() {
		switch command {
		case "category-store":
			handler.CategoryStoreHandler(scanner)
		case "logout":
			handler.Logout()
		case "category-list":
			handler.CategoryGetHandler()
		default:
			fmt.Printf("valid commands : category-store\nlogout\ncategory-list\n")
		}
	} else {
		fmt.Println("you are not authenticated ! please login or register ")
		switch command {
		case "register":
			handler.Register(scanner)
		case "login":
			handler.Login(scanner)
		default:
			fmt.Println("valid commands : login | register")
		}
	}

}
