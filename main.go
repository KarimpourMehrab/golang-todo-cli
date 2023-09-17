package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

var loggedIn = false
var userStorage []User

func main() {
	command := flag.String("command", "no command", "create a new to do !")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	if loggedIn {
		for {
			runCommand(*command)
			fmt.Println("please enter another command :")
			scanner.Scan()
			*command = scanner.Text()
		}
	} else {
		loginUser(scanner)
	}
}

func runCommand(command string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("cmd : ", command)
	switch command {
	case "create-task":
		createTask(scanner)
	case "create-category":
		createCategory(scanner)
	case "register-user":
		registerUser(scanner)
	case "login-user":
		loginUser(scanner)
	case "exit":
		os.Exit(0)
	case "q":
		os.Exit(0)
	default:
		fmt.Printf("command ' %s ' is not valid !\n", command)
	}
}

func createTask(scanner *bufio.Scanner) {
	var name, dueDate, category string

	name = scanText("please enter the task title : ", scanner)
	dueDate = scanText("please enter the task due date : ", scanner)
	category = scanText("please enter the task category : ", scanner)

	fmt.Println(name, dueDate, category)
}
func createCategory(scanner *bufio.Scanner) {

	category := scanText("please enter the task category : ", scanner)
	color := scanText("please enter the color of category : ", scanner)

	fmt.Println(category, color)
}
func registerUser(scanner *bufio.Scanner) {

	name := scanText("please enter the user  full name : ", scanner)
	email := scanText("please enter the user email : ", scanner)
	password := scanText("please enter the user password : ", scanner)

	user := User{
		ID:       len(userStorage) + 1,
		Name:     name,
		Password: password,
		Email:    email,
	}
	userStorage = append(userStorage, user)
	fmt.Println(userStorage)

}

func loginUser(scanner *bufio.Scanner) {
	email := scanText("please enter the user email :", scanner)
	password := scanText("please enter the user password :", scanner)

	for _, user := range userStorage {
		if user.Email == email && user.Password == password {
			loggedIn = true
			break
		}
	}
	if loggedIn {
		fmt.Println("you logged in successfully...")
	} else {
		fmt.Println("your email or password is incorrect ! ")
		registerUser(scanner)
	}

}

func scanText(str string, scanner *bufio.Scanner) string {
	fmt.Print(str)
	scanner.Scan()
	return scanner.Text()
}
