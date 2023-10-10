package main

import "fmt"

type UserEntity struct {
	ID   int
	name string
	role string
}

type CategoryUser struct {
	UserRepositoryInterface
	CategoryRepositoryInterface
}

type CategoryEntity struct {
	ID    int
	title string
}

func (c CategoryEntity) Get() []CategoryEntity {
	return make([]CategoryEntity, 0)
}

type UserRepositoryInterface interface {
	Get() []UserEntity
	IsAdmin() bool
}

type CategoryRepositoryInterface interface {
	Get() []CategoryEntity
}

func (u *UserEntity) Get() []UserEntity {
	return make([]UserEntity, 0)
}

func (u *UserEntity) IsAdmin() bool {
	return u.role == "admin"
}

func (c *CategoryUser) Get() []CategoryEntity {
	if c.IsAdmin() {
		return make([]CategoryEntity, 10)
	}
	return make([]CategoryEntity, 0)
}

func main() {
	var user UserRepositoryInterface
	var category CategoryRepositoryInterface

	user = &UserEntity{ID: 2, name: "mehrab", role: "admin"}
	category = &CategoryEntity{ID: 1, title: "test"}

	catUser := CategoryUser{user, category}
	res := catUser.Get()
	fmt.Println(res)

}
