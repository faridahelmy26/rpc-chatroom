package main

import (
    "fmt"
    "log"
    "net"
    "net/rpc"
    "sync"
)

type ChatServer struct {
    messages []string
    mu       sync.Mutex
}

// Request struct for sending messages
type MessageRequest struct {
    Text string
}

// Response struct for returning chat history
type MessageResponse struct {
    History []string
}

// Remote procedure to send a message and return chat history
func (s *ChatServer) SendMessage(req MessageRequest, res *MessageResponse) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    // Append message to chat history
    s.messages = append(s.messages, req.Text)

    // Return all messages
    res.History = s.messages
    return nil
}

func main() {
    server := new(ChatServer)
    rpc.Register(server)

    listener, err := net.Listen("tcp", ":1234")
    if err != nil {
        log.Fatal("Error starting server:", err)
    }

    fmt.Println("🚀 Chat Server is running on port 1234...")
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Println("Connection error:", err)
            continue
        }
        go rpc.ServeConn(conn)
    }
}
