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
		log.Fatal("Error connecting to server:", err)
	}
	defer client.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("💬 Welcome to the Simple Chatroom!")
	fmt.Println("Type 'exit' to leave the chat.")

	for {
		fmt.Print("You: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("👋 Exiting chat...")
			break
		}

		var reply struct{ History []string }
		err = client.Call("ChatServer.SendMessage", struct{ Text string }{Text: text}, &reply)
		if err != nil {
			log.Println("Error:", err)
			break
		}

		fmt.Println("\n--- Chat History ---")
		for _, msg := range reply.History {
			fmt.Println(msg)
		}
		fmt.Println("--------------------\n")
	}
}
