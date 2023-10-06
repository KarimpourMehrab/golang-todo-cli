package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
	"strings"
	"todo/repository/memoryStore"
	"todo/service/task"
)

type (
	User struct {
		ID       int
		Name     string
		Email    string
		Password string
	}

	Task struct {
		ID         int
		Title      string
		DueDate    string
		CategoryId int
		isDone     bool
		UserId     int
	}

	Category struct {
		ID     int
		Title  string
		Color  string
		UserId int
	}
)

var (
	userStorage       []User
	taskStorage       []Task
	categoryStorage   []Category
	AuthenticatedUser *User
	serializationMode string
)

const (
	userStoragePath        = "./files/user.txt"
	otherSerializationMode = "other"
	jsonSerializationMode  = "json"
)

func (u *User) print() {
	fmt.Println(u.ID, u.Name, u.Email)
}

func main() {
	TaskMemoryRepo := memoryStore.NewTaskRepo()
	taskService := task.NewService(TaskMemoryRepo)
	flag.String("serializationMode", jsonSerializationMode, "serialization mode to write data in storage")
	serializationScanner := bufio.NewScanner(os.Stdin)
	serializationM := scanText(fmt.Sprintf("please enter the serialization mode, %s | %s ", jsonSerializationMode, otherSerializationMode), serializationScanner)

	if serializationM == otherSerializationMode {
		serializationMode = otherSerializationMode
	} else {
		serializationMode = jsonSerializationMode
	}

	loadUserStorage(serializationMode)

	command := flag.String("command", "no command", "create a new to do !")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		runCommand(*command, &taskService)
		fmt.Println("please enter another command :")
		scanner.Scan()
		*command = scanner.Text()
	}
}

func loadUserStorage(serializationMode string) {
	file, _ := os.Open(userStoragePath)
	var data = make([]byte, 300)
	_, err := file.Read(data)
	if err != nil {
		fmt.Println("read file error ", err)
	}
	usersSlice := strings.Split(string(data), "\n")
	usersSlice = usersSlice[:len(usersSlice)-1]

	for _, usrStr := range usersSlice {
		var userStruct User
		switch serializationMode {
		case otherSerializationMode:
			userStruct, err = deserializeOtherSerializationUser(usrStr)
			if err != nil {
				fmt.Println(err)
				return
			}

			break
		default:
			uErr := json.Unmarshal([]byte(usrStr), &userStruct)
			if uErr != nil {
				fmt.Println("error in user unmarshal process !", uErr)
				return
			}
		}
		userStruct.print()
		userStorage = append(userStorage, userStruct)
	}

}
func deserializeOtherSerializationUser(usrStr string) (User, error) {

	if usrStr == "" {
		return User{}, fmt.Errorf("invalid user string for deserilize it ! please pass the valid string")
	}
	userFields := strings.Split(usrStr, ",")

	var user = User{}
	for _, field := range userFields {
		fieldSlice := strings.Split(field, ":")
		if len(fieldSlice) != 2 {
			fmt.Println("field is not valid !")
			continue
		}
		fieldName := strings.ReplaceAll(fieldSlice[0], " ", "")
		fieldValue := strings.ReplaceAll(fieldSlice[1], " ", "")
		switch fieldName {
		case "id":
			id, err := strconv.Atoi(fieldValue)
			if err != nil {
				return User{}, errors.New("cannot convert string to integer for user ID filed ")
			}
			user.ID = id
		case "name":
			user.Name = fieldValue
		case "email":
			user.Email = fieldValue
		case "password":
			user.Password = fieldValue
		}
	}
	return user, nil
}

func runCommand(command string, taskService *task.Service) {
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
		createTask(scanner, taskService)
	case "create-category":
		createCategory(scanner)
	case "register":
		registerUser(scanner)
	case "list-task":
		listTasks(taskService)
	case "exit":
		os.Exit(0)
	case "q":
		os.Exit(0)
	default:
		fmt.Printf("command ' %s ' is not valid !\n", command)
	}
}

func createTask(scanner *bufio.Scanner, service *task.Service) {
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

	response, errCreateTask := service.CreateTask(task.CreateRequest{
		Title:               title,
		CategoryId:          categoryId,
		DueDate:             dueDate,
		AuthenticatedUserId: AuthenticatedUser.ID,
	})
	if errCreateTask != nil {
		fmt.Println(errCreateTask)
	}

	fmt.Println(response)
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

	passHashedByte, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		fmt.Println("hashing password has error : ", err)
	}

	user := User{
		ID:       len(userStorage) + 1,
		Name:     name,
		Password: string(passHashedByte),
		Email:    email,
	}
	userStorage = append(userStorage, user)
	writeToFile(user)
}

func writeToFile(user User) {
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
	var userData []byte

	if serializationMode == otherSerializationMode {
		userData = []byte(fmt.Sprintf("id:%v, name:%s, email:%s, password:%s \n", len(userStorage)+1, user.Name, user.Email, user.Password))
	} else {
		var jsonMarshalingError error
		userData, jsonMarshalingError = json.Marshal(user)
		if jsonMarshalingError != nil {
			fmt.Println("error in marshaling json ! ", jsonMarshalingError)
			return
		}
	}
	userData = append(userData, []byte("\n")...)
	_, err = file.Write(userData)
	defer file.Close()
	if err != nil {
		fmt.Println("the error of write file ! ", err)
	}

	fmt.Println(err)
	fmt.Println(userStorage)
}

func loginUser(scanner *bufio.Scanner) {
	email := scanText("please enter the user email :", scanner)
	password := scanText("please enter the user password :", scanner)

	for _, user := range userStorage {

		if user.Email == email && nil == bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) {
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

func listTasks(service *task.Service) {
	usersTasks, err := service.List(task.ListRequest{
		UserId: AuthenticatedUser.ID,
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(usersTasks)
}
