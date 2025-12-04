---

# 💬 Simple RPC Chatroom (Go)

This project is a simple chatroom implemented using **Go RPC**.
Multiple clients can connect to the server, send messages, and view the full chat history returned by the server.

It is based on a classic request/response RPC model where the server stores all messages in memory and returns the full chat history to each client after every new message.

---

## 📌 Features

* Connect multiple clients to the chat server
* Send text messages
* View complete chat history after each message
* Uses **Go net/rpc** for communication
* Thread-safe message storage using **sync.Mutex**

---

## 📁 Project Structure

```
.
├── server.go   # RPC server that stores and returns chat messages
└── client.go   # Client that sends messages and displays chat history
```

---

## 🚀 How to Run

### **1️⃣ Start the Server**

Open a terminal and run:

```bash
go run server.go
```

You should see:

```
🚀 Chat Server is running on port 1234...
```

---

### **2️⃣ Start a Client**

Open another terminal window and run:

```bash
go run client.go
```

You will see:

```
💬 Welcome to the Simple Chatroom!
Type 'exit' to leave the chat.
```

Start typing messages!

---

### **3️⃣ Open Multiple Clients**

Open multiple terminals and run:

```bash
go run client.go
```

Each client will send messages to the same server and receive the full chat history.

---

## 🧠 How It Works

### **Server**

* Stores all chat messages in a slice
* Protects the message list using a Mutex
* Each call to `SendMessage`:

  * Appends the message
  * Returns all messages back to the client

### **Client**

* Reads user input
* Sends the message via RPC to the server
* Displays the entire chat history returned by the server
* Exits when the user types `exit`

---

## 📌 Example Interaction

**Client 1:**

```
You: Hello
--- Chat History ---
Hello
```

**Client 2:**

```
You: Hi everyone!
--- Chat History ---
Hello
Hi everyone!
```

---

🎬 **Real-Time Chat Application Demo**
[https://drive.google.com/drive/my-drive?hl=ar](https://drive.google.com/drive/my-drive?hl=ar)


## 🛠 Requirements

* Go 1.18+
* Basic terminal to run multiple processes

---