package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"todo/delivery/deliveryparam"
)

func main() {
	conn, connErr := net.Dial("tcp", "localhost:5000")
	command := os.Args[1]

	if connErr != nil {
		fmt.Println(connErr)
		return
	}

	req := deliveryparam.Request{
		Command: command,
	}
	serializedRequest, mErr := json.Marshal(req)
	if mErr != nil {
		fmt.Println(mErr)
		return
	}
	_, wErr := conn.Write(serializedRequest)

	if wErr != nil {
		fmt.Println(wErr)
	}

	response := make([]byte, 1024)
	_, rErr := conn.Read(response)
	if rErr != nil {
		fmt.Println(rErr)
	}

	fmt.Println(string(response))
}
