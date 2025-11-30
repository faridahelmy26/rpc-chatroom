package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
)

type Client struct {
	ID int
	Ch chan string
}

type ChatServer struct {
	mu      sync.Mutex
	clients map[int]*Client
	nextID  int
}

type JoinRequest struct{}
type JoinResponse struct {
	ID int
}

type MessageRequest struct {
	ID   int
	Text string
}

type Empty struct{}

func NewChatServer() *ChatServer {
	return &ChatServer{
		clients: make(map[int]*Client),
	}
}

func (s *ChatServer) Join(req JoinRequest, res *JoinResponse) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := s.nextID
	s.nextID++

	client := &Client{
		ID: id,
		Ch: make(chan string, 10),
	}
	s.clients[id] = client

	// Notify others
	go s.broadcast(fmt.Sprintf("🔵 User %d joined", id), id)

	res.ID = id
	return nil
}

func (s *ChatServer) SendMessage(req MessageRequest, _ *Empty) error {
	go s.broadcast(fmt.Sprintf("User %d: %s", req.ID, req.Text), req.ID)
	return nil
}

func (s *ChatServer) ReceiveMessages(id int, reply *[]string) error {
	s.mu.Lock()
	ch := s.clients[id].Ch
	s.mu.Unlock()

	msg := <-ch
	*reply = []string{msg}
	return nil
}

func (s *ChatServer) broadcast(msg string, senderID int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for id, client := range s.clients {
		if id != senderID {
			client.Ch <- msg
		}
	}
}

func main() {
	server := NewChatServer()
	rpc.Register(server)

	ln, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("🚀 Real-Time RPC Chat Server Running...")

	for {
		conn, _ := ln.Accept()
		go rpc.ServeConn(conn)
	}
}
