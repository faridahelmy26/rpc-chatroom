---

# 💬 Real-Time Distributed Chatroom using Go RPC & Concurrency

This project is a **real-time distributed chat system** built using **Go RPC**, **goroutines**, **channels**, and **mutexes**.
It allows multiple clients to join, send messages, and receive broadcasts instantly without reloading the full history.

---

## 🧠 Project Overview

Unlike a simple request–response chat, this version implements:

### ✅ **Real-Time Broadcasting**

* When a client joins, all other clients instantly receive:

  ```
  🔵 User [ID] joined
  ```
* Any message sent by a client is broadcast to **all other connected clients**.

### ✅ **No Self-Echo**

* The sender does NOT receive their own message.

### ✅ **Concurrency with Channels & Goroutines**

* Each client has a dedicated message channel.
* A broadcasting goroutine pushes new messages to all clients.
* Shared state (client list) is protected using a `sync.Mutex`.

### ✅ **Multiple Clients Supported**

* Any number of clients can join the chat.
* All receive updates in real time.

---

## 🗂 Project Structure

```
📁 rpc-realtime-chat
 ├── server.go     # Real-time broadcasting server
 ├── client.go     # Concurrent client with listener goroutine
 └── README.md
```

---

## ⚙️ How to Run the Project

### 🟢 1. Run the Server

```bash
go run server.go
```

You should see:

```
🚀 Real-Time RPC Chat Server Running...
```

---

### 🔵 2. Run the Client

Open a new terminal for each client:

```bash
go run client.go
```

You will see:

```
💬 Joined chat as User X
Type messages and press Enter.
```

The client has two goroutines:

* **Sender** → when YOU type a message
* **Listener** → receives real-time broadcasts from the server

---

## 🧪 Example Interaction

### Client 1:

```
💬 Joined chat as User 0
You: Hello everyone!
```

### Client 2:

```
🔵 User 0 joined
User 0: Hello everyone!
You:
```

### Client 3 joins:

All other clients receive:

```
🔵 User 2 joined
```

---

## 🧵 Concurrency Features Used

### ✔ Goroutines

* Server broadcasting
* Client background message listener

### ✔ Channels

* Each client gets its own message channel

### ✔ Mutex

* Protects shared map of connected clients

---

## 📺 Demo Video

🎬 **Real-Time Chat Application Demo**
[https://drive.google.com/file/d/1lYxbbbgiKfi3V1DqOCwnbtZ41A6k6FQw/view?usp=sharing](https://drive.google.com/file/d/1lYxbbbgiKfi3V1DqOCwnbtZ41A6k6FQw/view?usp=sharing)

---



