package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type Task struct {
	ID         int
	Title      string
	DueDate    string
	CategoryId int
	isDone     bool
	UserId     int
}

type Category struct {
	ID     int
	Title  string
	Color  string
	UserId int
}

var userStorage []User
var taskStorage []Task
var categoryStorage []Category
var AuthenticatedUser *User

const userStoragePath = "./files/user.txt"

func (u *User) print() {
	fmt.Println(u.ID, u.Name, u.Email)
}

func main() {
	//TODO : load the user storage from txt file
	loadUserStorage()

	return
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

func loadUserStorage() {
	file, _ := os.Open(userStoragePath)
	//fileScanner := bufio.NewScanner(file)
	//fileScanner.Split(bufio.ScanLines)
	//for fileScanner.Scan() {
	//	fmt.Println(fileScanner.Text())
	//}
	var data = make([]byte, 500)
	_, err := file.Read(data)
	if err != nil {
		fmt.Println("read file error ", err)
	}
	usersSlice := strings.Split(string(data), "\n")
	usersSlice = usersSlice[:len(usersSlice)-1]

	for _, user := range usersSlice {
		userSlice := strings.Split(user, ",")

		for _, user := range userSlice {
			user := strings.Split(user, ":")
			if len(user) > 1 {
				fmt.Println(user[0])
				//fmt.Printf("%v : %v\n", user[0], user[1])
			}
		}
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
	categoryId, err := strconv.Atoi(category)
	if err != nil {
		fmt.Println("please enter the valid category ! ")
		return
	}
	categoryIdIsValid := false
	for _, c := range categoryStorage {
		if c.ID == categoryId && c.UserId == AuthenticatedUser.ID {
			categoryIdIsValid = true
			break
		}
	}
	if !categoryIdIsValid {
		fmt.Printf("authenticated user dosnt have category with id %v", categoryId)
		return
	}

	newTask := Task{
		ID:         len(taskStorage) + 1,
		Title:      title,
		CategoryId: categoryId,
		DueDate:    dueDate,
		isDone:     false,
		UserId:     AuthenticatedUser.ID,
	}

	taskStorage := append(taskStorage, newTask)
	fmt.Println(taskStorage)
}
func createCategory(scanner *bufio.Scanner) {
	AuthenticatedUser.print()
	title := scanText("please enter the task category : ", scanner)
	color := scanText("please enter the color of category : ", scanner)

	c := Category{
		ID:     len(categoryStorage) + 1,
		Title:  title,
		Color:  color,
		UserId: AuthenticatedUser.ID,
	}
	categoryStorage = append(categoryStorage, c)
	fmt.Println(categoryStorage)

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

	_, err := os.Stat(userStoragePath)
	var file *os.File
	if err != nil {
		fmt.Printf("the file %v dosn't exists ! \n", userStoragePath)
		var createErr error
		file, createErr = os.Create(userStoragePath)
		if createErr != nil {
			fmt.Println(createErr)
			return
		}
	} else {
		var openError error
		file, openError = os.OpenFile(userStoragePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if openError != nil {
			fmt.Println(openError)
			return
		}

	}
	//newData := []byte("name: " + name + ", email: " + email + ", password: " + password + "\n")
	newData := []byte(fmt.Sprintf("id: %v, name: %s, email: %s, password: %s \n", len(userStorage)+1, name, email, password))
	_, err = file.Write(newData)
	if err != nil {
		fmt.Println("the error of write file ! ", err)
	}
	file.Close()

	fmt.Println(err)
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
