package deliveryparam

type Request struct {
	Command               string
	CreateTaskRequest     CreateTaskRequest
	CreateCategoryRequest CreateCategoryRequest
}

type CreateTaskRequest struct {
	Title      string
	DueDate    string
	CategoryId int
}

type CreateCategoryRequest struct {
	Title string
	Color string
}
