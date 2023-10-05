package handler

import (
	"bufio"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"todo/refactor/entity"
	"todo/refactor/helper"
	"todo/refactor/service"
)

var userStorage service.File
var users []entity.User
var AuthenticatedUser *entity.User

func init() {
	userStorage.SetPath("/user.txt")
	loadUsers()
}

func UserIsAuthenticated() bool {
	return AuthenticatedUser != nil
}

func UserIsNotGuest() bool {
	return AuthenticatedUser != nil
}

func userAlreadyRegistered(username string) bool {
	if UserIsNotGuest() {
		return true
	}

	for _, user := range users {
		if user.Username == username {
			return true
		}
	}

	return false
}

func Logout() {
	AuthenticatedUser = nil
	fmt.Println("logout is successful...")
}

func Login(scanner *bufio.Scanner) {
	if UserIsAuthenticated() {
		fmt.Printf("dear %s you are already authenticated!\n", AuthenticatedUser.Username)
		return
	}
	username := helper.ScanText("username :", scanner)
	password := helper.ScanText("password :", scanner)
	for _, user := range users {
		if user.Username == username && user.Password == password {
			AuthenticatedUser = &user
			fmt.Println("authentication successful...")
		}
	}
}

func Register(scanner *bufio.Scanner) {
	if UserIsNotGuest() {
		fmt.Printf("dear %s you are already registered!\n", AuthenticatedUser.Username)
		return
	}

	username := helper.ScanText("enter username :", scanner)
	if userAlreadyRegistered(username) {
		fmt.Printf("username %s already reserved !\n", username)
		return
	}
	password := helper.ScanText("enter password :", scanner)
	_, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		fmt.Println(err)
	}
	id := 1
	userStore(&entity.User{
		ID:       id,
		Username: username,
		Password: password,
	})
}

func userStore(u *entity.User) {
	user, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	userStorage.Store(user)
}

func loadUsers() {
	var userStruct entity.User
	data := userStorage.Get()
	userSlices := strings.Split(string(data), "\n")
	userSlices = userSlices[:len(userSlices)-1]
	for _, userSlice := range userSlices {
		err := json.Unmarshal([]byte(userSlice), &userStruct)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, userStruct)
	}
}
