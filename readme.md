# 💬 Simple Chatroom using RPC (GoLang)

This project is a simple **distributed chatroom** implemented using **Go RPC**.  
It allows multiple clients to send and receive messages through a single server.

---

## 🧠 Project Description

- The **server** receives messages from clients, stores them in a list, and returns the complete chat history.
- The **client** connects to the server using RPC, sends messages, and displays the chat history after each message.
- The chat runs continuously until the user types `exit` or presses `Ctrl + C`.

---

## ⚙️ How to Run the Project

### 🟢 Step 1 — Run the Server
```bash
cd server
go run server.go
````

You should see:

```
🚀 Chat Server is running on port 1234...
```

---

### 🔵 Step 2 — Run the Client

Open another terminal and run:

```bash
cd client
go run client.go
```

You will see:

```
💬 Welcome to the Simple Chatroom!
Type 'exit' to leave the chat.
```

Type your message and press Enter.
After sending a message, you will see the full chat history displayed.

---

## 🎥 Running Application Video

Here is the video showing the running chatroom:

🎬 [Watch Demo Video on Google Drive](https://drive.google.com/file/d/1W9bkgUy0l7ee452fp5s2IvG21fUNJgQQ/view?usp=sharing)
---


