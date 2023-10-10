package category

import (
	"fmt"
	"todo/entity"
)

type (
	CreateRequest struct {
		Title  string
		Color  string
		UserId int
	}

	CreateResponse struct {
		entity.Category
		metadata string
	}

	repository interface {
		CreateCategory(c entity.Category) (entity.Category, error)
		ListCategory(userId int) ([]entity.Category, error)
	}

	Service struct {
		Repository repository
	}

	ListRequest struct {
		UserId int
	}

	ListResponse struct {
		categories []entity.Category
	}
)

func NewService(rep repository) Service {
	return Service{rep}
}

func (s Service) CreateCategory(req CreateRequest) (CreateResponse, error) {
	createdCategory, createCatErr := s.Repository.CreateCategory(entity.Category{
		Title:  req.Title,
		UserId: req.UserId,
		Color:  req.Color,
	})

	if createCatErr != nil {
		return CreateResponse{}, createCatErr
	}

	return CreateResponse{Category: createdCategory, metadata: ""}, nil
}

func (s Service) ListCategory(userId int) (ListResponse, error) {
	var response ListResponse
	var listError error

	response.categories, listError = s.Repository.ListCategory(userId)

	if listError != nil {
		fmt.Println(listError)
		return ListResponse{}, listError
	}

	return response, nil

}
