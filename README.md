# 📝 go-react-demo

A full-stack To-Do List demo using **React (frontend)** + **Go (backend)** + **MongoDB (local)**.  
This project demonstrates how Go can replace Node.js in a MERN-style stack with a clean architecture and high performance.

---

## 📌 Project Purpose

This demo is part of a university assignment to explore web technologies beyond the MERN stack. It demonstrates:

- Building a backend API using **Go + Gin**
- Connecting to a **local MongoDB** database using Go’s native driver
- Integrating with a **React frontend**
- Applying **best practices** in project structure and maintainability

---

## ⚙️ Tech Stack

| Layer     | Technology              |
|-----------|--------------------------|
| Frontend  | React (Vite + Fetch API) |
| Backend   | Go + Gin (REST API)      |
| Database  | MongoDB (local)          |

---

## 📁 Folder Structure

```bash
go-react-demo/
├── frontend/                 # React client
│   ├── src/
│   │   ├── App.jsx
│   │   ├── App.css
│   │   └── config.js         # API URL config
│   └── index.html
│
└── backend/                  # Go API server
    ├── main.go               # Entry point
    ├── database/
    │   └── connection.go     # MongoDB connection
    ├── models/
    │   └── todo_model.go     # Data structure
    ├── services/
    │   └── todo_service.go   # DB operations
    ├── controllers/
    │   └── todo_controller.go # API logic
    └── routes/
        └── todo_routes.go    # Route registration
```

---

## 🚀 How to Run

### 1. Start MongoDB (local)

Make sure MongoDB is installed and running at `localhost:27017`.

```bash
brew services start mongodb-community@6.0
```

---

### 2. Start the Go backend

```bash
cd backend
go mod tidy
go run main.go
```

Expected output:

```
✅ Connected to MongoDB
[GIN-debug] Listening and serving HTTP on :8080
```

---

### 3. Start the React frontend

```bash
cd frontend
npm install
npm run dev
```

Then open [http://localhost:5173](http://localhost:5173) in your browser.

---

## 📡 API Endpoints

Base URL: `http://localhost:8080/api/todos`

| Method | Endpoint           | Description         |
|--------|--------------------|---------------------|
| GET    | `/api/todos`       | Get all todos       |
| POST   | `/api/todos`       | Create new todo     |
| PUT    | `/api/todos/:id`   | Update a todo       |
| DELETE | `/api/todos/:id`   | Delete a todo       |

### Example JSON:

```json
{
  "_id": "ObjectId",
  "description": "Finish homework",
  "isComplete": false
}
```

---

## 🔍 MERN vs Go Comparison

| Feature          | MERN (Node.js)            | This Demo (Go)             |
|------------------|----------------------------|-----------------------------|
| Backend Language | JavaScript + Express       | Go + Gin                    |
| Performance      | Async, single-threaded     | Compiled, multi-threaded    |
| Tooling          | NPM ecosystem              | Strong typing, fast compile |
| Learning Curve   | Easy for JS developers     | Ideal for backend learners  |

---

## ✅ Features Implemented

- [x] View todo list (GET)
- [x] Add new todo (POST)
- [x] Update todo status (PUT)
- [x] Delete todo (DELETE)
- [x] React frontend with live update
- [x] MongoDB storage with persistence

---

## 🙋 Author

Zephyr — COMPSCI732 Tech Demo  
**Focus**: Exploring Go as a modern backend alternative to Node.js

---

## 📦 Notes

- `.gitignore` should include `node_modules` and `dist`
- Make sure MongoDB is running before starting backend
- Project emphasizes simplicity, modularity, and clean code for learning
