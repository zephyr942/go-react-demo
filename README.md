# 📝 Go + React Todo List Tech Demo

This project demonstrates how to build a simple full-stack Todo List application using **React** (frontend), **Go with Gin** (backend), and **MongoDB** (database). It is created as part of the Tech Demo Assignment for the COMPSCI732 Web Development course.

> ✅ This demo is designed for **new learners** and includes simple, clean code with helpful comments.

---

## 🔧 Tech Stack

| Layer      | Tech Used               |
|------------|-------------------------|
| Frontend   | React + Vite            |
| Backend    | Go (Gin framework)      |
| Database   | MongoDB (local instance)|

---

## 🗂️ Project Structure
```bash
go-react-demo/
├── backend-go/                 # Go backend
│   ├── controllers/            # Route handlers (GET, POST, etc.)
│   ├── models/                 # MongoDB schema (Todo model)
│   ├── routes/                 # API route registration
│   ├── database/               # MongoDB connection
│   └── main.go                 # Entry point of the backend
│
├── frontend-react/            # React frontend
│   └── src/
│       ├── App.jsx            # Main React component
│       ├── App.css            # Styling
│       └── main.jsx           # Vite entry point
│
├── index.html
├── README.md                  # ← You are here
└── go.mod                     # Go module config
```
---

## 🚀 How to Run This Project

### 1️⃣ Prerequisites

To run this project locally, ensure the following are installed and properly set up:

- Go 1.20+ installed

- Node.js + npm installed

- MongoDB running locally (`mongodb://localhost:27017`)

#### ✅ Go 1.20+ installed

- Download: https://go.dev/dl/
or brew install：
```bash
  brew install go
```
- Verify installation:
```bash
  go version
```
#### ✅ Node.js + npm installed (for React frontend)
- Download: https://nodejs.org/
or brew install：
```bash
  brew install node
```
- Verify installation:
```bash
node -v
npm -v
```
#### ✅ MongoDB running locally
Download and install: https://www.mongodb.com/try/download/community
or brew install：
```bash
  brew tap mongodb/brew
  brew install mongodb-community
```

Start MongoDB service (Mac example):
```bash
brew services start mongodb-community
```

To stop the service:
```bash
brew services stop mongodb-community
```

To Check MongoDB service status::
```bash
brew services list
```

Connect to the database using the MongoDB shell:
```bash
mongosh
```

Default connection string used in this project:
```bash
mmongodb://localhost:27017
```

Example output when mongosh connects::
```bash
Connecting to:          mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.4.2
Using MongoDB:          7.0.17
Using Mongosh:          2.4.2
```

---

### 2️⃣ Start the Backend

```bash
cd backend-go
go run main.go
```
You should see: ✅ Connected to MongoDB successfully
```bash
The backend runs on: http://localhost:8080/api/todos
```
### 3️⃣ Start the Frontend

```bash
cd frontend-react
npm install
npm run dev
```
```bash
The frontend runs on: http://localhost:5173
```

---

## 🧪 Features:
 - Add a new todo

 - Mark todo as complete/incomplete

 - Delete a todo

 - View live list of todos

 - Update todo status in MongoDB

---

## 🎥 Demo Video
👉 The full walkthrough video (15 mins) will be uploaded to Canvas.
It includes:

 - Architecture explanation

 - Code walkthrough (frontend + backend)

 - Live demo

 - Summary

 ---

## ❗ Notes
This project is not built for production. It is designed for educational/demo purposes.

MongoDB is assumed to be running locally. If you want to connect to Atlas or Docker, modify the connection string in database/connection.go.

---

## 🙋‍♂️ Author
Zephyr (COMPSCI732 - Tech Demo)\
UID:bche942
University of Auckland

---

## 📄 License
MIT License. Use freely for learning and academic purposes.