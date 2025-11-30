package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Error connecting:", err)
	}

	// Join server
	var joinRes struct{ ID int }
	client.Call("ChatServer.Join", struct{}{}, &joinRes)
	id := joinRes.ID

	fmt.Println("💬 Joined chat as User", id)
	fmt.Println("Type messages and press Enter.")

	// Start listening goroutine
	go func() {
		for {
			var msgs []string
			err := client.Call("ChatServer.ReceiveMessages", id, &msgs)
			if err == nil && len(msgs) > 0 {
				fmt.Println("\n" + msgs[0])
				fmt.Print("You: ")
			}
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	// Message loop
	for {
		fmt.Print("You: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("👋 Leaving...")
			return
		}

		client.Call("ChatServer.SendMessage", struct {
			ID   int
			Text string
		}{ID: id, Text: text}, &struct{}{})
	}
}
