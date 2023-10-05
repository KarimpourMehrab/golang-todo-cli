package handler

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strings"
	"todo/refactor/entity"
	"todo/refactor/helper"
	"todo/refactor/service"
)

var categoryStorage service.File
var categories []entity.Category

func init() {
	fmt.Println(categories)
	categoryStorage.SetPath("/category.txt")
	loadCategory()

	fmt.Println(categories)
}

func CategoryStoreHandler(scanner *bufio.Scanner) {

	title := helper.ScanText("enter category title :", scanner)
	color := helper.ScanText("enter category color :", scanner)
	id := 1
	userId := 1
	categoryStore(&entity.Category{
		ID:     id,
		Title:  title,
		Color:  color,
		UserId: userId,
	})
}

func CategoryGetHandler() []entity.Category {
	fmt.Println(categories)
	return categories
}

func categoryStore(c *entity.Category) {
	data, _ := json.Marshal(c)
	categoryStorage.Store(data)
	fmt.Println("new category stored successful...")
	loadCategory()
}

func loadCategory() {
	var categoryStruct entity.Category
	var categoriesSlice []entity.Category
	data := categoryStorage.Get()
	categorySlice := strings.Split(string(data), "\n")
	categorySlice = categorySlice[:len(categorySlice)-1]
	for _, category := range categorySlice {
		err := json.Unmarshal([]byte(category), &categoryStruct)
		if err != nil {
			fmt.Println(err)
		}
		categoriesSlice = append(categoriesSlice, categoryStruct)
	}
	categories = categoriesSlice
}
