package main

import (
	"encoding/json"
	"fmt"
	"net"
	"todo/repository/memoryStore"
	"todo/service/task"
)

func main() {
	taskRepo := memoryStore.NewTaskRepo()
	taskService := task.NewService(taskRepo)
	listener, listenErr := net.Listen("tcp", "localhost:5000")
	if listenErr != nil {
		fmt.Println(listenErr)
	}

	defer func() {
		lisErr := listener.Close()
		if lisErr != nil {
			fmt.Println(lisErr)
		}
	}()

	type Request struct {
		Command string
	}

	for {
		data := make([]byte, 3024)
		conn, connErr := listener.Accept()
		if connErr != nil {
			fmt.Println(connErr)
			continue
		}

		_, readErr := conn.Read(data)
		if readErr != nil {
			fmt.Println(readErr)
		}

		var req = &Request{Command: string(data)}
		if reqErr := json.Unmarshal(data, req); reqErr != nil {
			fmt.Println(reqErr)
			continue
		}

		switch req.Command {
		case "create-task":
			response, cTaskErr := taskService.CreateTask(task.CreateRequest{
				Title:               "",
				DueDate:             "",
				CategoryId:          1,
				AuthenticatedUserId: 1,
			})

			if cTaskErr != nil {
				_, wErr := conn.Write([]byte(cTaskErr.Error()))
				if wErr != nil {
					fmt.Println(wErr)
				}
				continue
			}
			cleanedRes, _ := json.Marshal(response)
			_, wErr := conn.Write(cleanedRes)
			if wErr != nil {
				fmt.Println(wErr)
				continue
			}

		}

	}
}
