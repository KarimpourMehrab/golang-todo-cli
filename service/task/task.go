package task

import "todo/entity"

type (
	CreateRequest struct {
		Title               string
		DueDate             string
		CategoryId          int
		AuthenticatedUserId int
	}

	CreateResponse struct {
		entity.Task
		MetaData string
	}

	repository interface {
		//UserHasCategory(UserID, CategoryID int) (bool, error)
		CreateTask(t entity.Task) (entity.Task, error)
		ListTask(userId int) ([]entity.Task, error)
	}

	Service struct {
		Repository repository
	}

	ListRequest struct {
		UserId int
	}
	ListResponse struct {
		Tasks []entity.Task
	}
)

func NewService(rep repository) Service {
	return Service{rep}
}

func (s Service) CreateTask(req CreateRequest) (CreateResponse, error) {
	//ok, err := s.Repository.UserHasCategory(req.AuthenticatedUserId, req.CategoryId)
	//if err != nil || !ok {
	//	return CreateResponse{}, err
	//}

	createdTask, createTaskErr := s.Repository.CreateTask(entity.Task{
		ID:         0,
		Title:      req.Title,
		UserId:     req.AuthenticatedUserId,
		CategoryId: req.CategoryId,
		IsDone:     false,
		DueDate:    req.DueDate,
	})

	if createTaskErr != nil {
		return CreateResponse{}, createTaskErr
	}

	return CreateResponse{
		Task:     createdTask,
		MetaData: "",
	}, nil

}

func (s Service) List(req ListRequest) (ListResponse, error) {
	var response ListResponse
	var listError error

	response.Tasks, listError = s.Repository.ListTask(req.UserId)
	if listError != nil {
		return ListResponse{}, listError
	}

	return response, nil
}
