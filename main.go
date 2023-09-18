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

type Task struct {
	ID       int
	Title    string
	DueDate  string
	Category string
	isDone   bool
	UserId   int
}

var userStorage []User
var taskStorage []Task
var AuthenticatedUser *User

func (u *User) print() {
	fmt.Println(u.ID, u.Name, u.Email)
}

func main() {
	command := flag.String("command", "no command", "create a new to do !")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		runCommand(*command)
		fmt.Println("please enter another command :")
		scanner.Scan()
		*command = scanner.Text()
	}
}

func runCommand(command string) {
	scanner := bufio.NewScanner(os.Stdin)

	if command != "register" && command != "exit" && AuthenticatedUser == nil {
		loginUser(scanner)
		if AuthenticatedUser == nil {
			return
		}
	}

	fmt.Println("cmd : ", command)
	switch command {
	case "create-task":
		createTask(scanner)
	case "create-category":
		createCategory(scanner)
	case "register":
		registerUser(scanner)
	case "list-task":
		listTasks()
	case "exit":
		os.Exit(0)
	case "q":
		os.Exit(0)
	default:
		fmt.Printf("command ' %s ' is not valid !\n", command)
	}
}

func createTask(scanner *bufio.Scanner) {
	AuthenticatedUser.print()
	var title, dueDate, category string

	title = scanText("please enter the task title : ", scanner)
	dueDate = scanText("please enter the task due date : ", scanner)
	category = scanText("please enter the task category : ", scanner)

	newTask := Task{
		ID:       len(taskStorage) + 1,
		Title:    title,
		Category: category,
		DueDate:  dueDate,
		isDone:   false,
		UserId:   AuthenticatedUser.ID,
	}

	taskStorage := append(taskStorage, newTask)
	fmt.Println(taskStorage)
}
func createCategory(scanner *bufio.Scanner) {
	AuthenticatedUser.print()
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
			AuthenticatedUser = &user

			break
		}
	}
	if AuthenticatedUser != nil {
		fmt.Println("you logged in successfully...")
	}
}

func scanText(str string, scanner *bufio.Scanner) string {
	fmt.Print(str)
	scanner.Scan()
	return scanner.Text()
}

func listTasks() {
	for _, task := range taskStorage {
		if task.UserId == AuthenticatedUser.ID {
			fmt.Println(task)

			break
		}
	}
}
