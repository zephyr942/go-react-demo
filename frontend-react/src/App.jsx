// App.jsx
import React, { useState, useEffect } from 'react';
import './App.css';

export default function App() {
  const [todos, setTodos] = useState([]);
  const [newTask, setNewTask] = useState("");

  const fetchTodos = async () => {
    const res = await fetch("http://localhost:8080/todos");
    const data = await res.json();
    setTodos(data);
  };

  useEffect(() => {
    fetchTodos();
  }, []);

  const addTodo = async () => {
    if (!newTask.trim()) return;
    await fetch("http://localhost:8080/todos", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ description: newTask, isComplete: false }),
    });
    setNewTask("");
    fetchTodos();
  };

  const deleteTodo = async (id) => {
    await fetch(`http://localhost:8080/todos/${id}`, {
      method: "DELETE",
    });
    fetchTodos();
  };

  const toggleTodo = async (todo) => {
    await fetch(`http://localhost:8080/todos/${todo.id}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ ...todo, isComplete: !todo.isComplete }),
    });
    fetchTodos();
  };

  return (
    <div className="container">
      <h1 className="header">📝 To Do List</h1>

      <div className="input-section">
        <input
          value={newTask}
          onChange={(e) => setNewTask(e.target.value)}
          placeholder="Enter task description"
          className="task-input"
        />
        <button onClick={addTodo} className="add-button">Add</button>
      </div>

      <ul className="todo-list">
        {todos.map((todo) => (
          <li
            key={todo.id}
            className={`todo-item ${todo.isComplete ? 'completed' : ''}`}
          >
            <div className="todo-left">
              <input
                type="checkbox"
                checked={todo.isComplete}
                onChange={() => toggleTodo(todo)}
              />
              <span className="todo-text">
                {todo.description}
              </span>
            </div>
            <button
              onClick={() => deleteTodo(todo.id)}
              className="delete-button"
            >
              Delete
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
}