package memoryStore

import (
	"todo/entity"
)

type Task struct {
	tasks []entity.Task
}

func NewTaskRepo() *Task {
	return &Task{
		tasks: make([]entity.Task, 0),
	}
}

func (t *Task) CreateTask(task entity.Task) (entity.Task, error) {
	task.ID = len(t.tasks) + 1
	t.tasks = append(t.tasks, task)

	return task, nil
}

func (t *Task) ListTask(userId int) ([]entity.Task, error) {
	var currentUserTasks []entity.Task
	for _, task := range t.tasks {
		if task.UserId == userId {
			currentUserTasks = append(currentUserTasks, task)
		}
	}

	return currentUserTasks, nil
}
