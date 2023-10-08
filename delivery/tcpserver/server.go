package main

import (
	"encoding/json"
	"fmt"
	"net"
	"todo/delivery/deliveryparam"
	"todo/repository/memoryStore"
	"todo/service/task"
)

func main() {
	taskRepo := memoryStore.NewTaskRepo()
	taskService := task.NewService(taskRepo)
	listener, listenErr := net.Listen("tcp", ":5000")
	if listenErr != nil {
		fmt.Println(listenErr)
	} else {
		fmt.Printf("tcp server listened on %s\n", listener.Addr())
	}

	defer func() {
		lisErr := listener.Close()
		if lisErr != nil {
			fmt.Println(lisErr)
		}
	}()

	for {
		data := make([]byte, 1024)
		conn, connErr := listener.Accept()
		if connErr != nil {
			fmt.Println("connErr", connErr)
			continue
		}

		numberOfReadData, readErr := conn.Read(data)
		if readErr != nil {
			fmt.Println("readErr", readErr)
		}

		var req = &deliveryparam.Request{Command: string(data)}
		if reqErr := json.Unmarshal(data[:numberOfReadData], req); reqErr != nil {
			fmt.Println("reqErr : ", reqErr)
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
					fmt.Println("wErr", wErr)
				}
				continue
			}
			cleanedRes, _ := json.Marshal(response)
			_, wErr := conn.Write(cleanedRes)
			if wErr != nil {
				fmt.Println("wErr", wErr)
				continue
			}
		}

	}
}
