// App.jsx
import React, { useState, useEffect } from 'react';
import './App.css';
import { API_URL } from './config';

export default function App() {

  const [todos, setTodos] = useState([]);
  const [newTask, setNewTask] = useState("");

  // Load todos when component mounts
  const fetchTodos = async () => {
    try {
      const res = await fetch(API_URL);
      const data = await res.json();
      setTodos(data);
    } catch (err) {
      console.error("Failed to fetch todos:", err);
    }
  };

  // Run fetchTodos once on page load
  useEffect(() => {
    fetchTodos();
  }, []);

  // Add a new todo
  const addTodo = async () => {
    if (!newTask.trim()) return;
    try {
      await fetch(API_URL, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ description: newTask }),
      });
      // Clear input
      setNewTask(""); 
      // Refresh list
      fetchTodos();    
    } catch (err) {
      console.error("Failed to add todo:", err);
    }
  };

  // Delete a todo by ID
  const deleteTodo = async (id) => {
    try {
      await fetch(`${API_URL}/${id}`, {
        method: "DELETE",
      });
      fetchTodos();
    } catch (err) {
      console.error("Failed to delete todo:", err);
    }
  };

  // Toggle todo status (complete/incomplete)
  const toggleTodo = async (todo) => {
    try {
      await fetch(`${API_URL}/${todo._id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          ...todo,
          isComplete: !todo.isComplete,
        }),
      });
      fetchTodos();
    } catch (err) {
      console.error("Failed to update todo:", err);
    }
  };

  return (
    <div className="container">
      <h1 className="header">📝 To Do List</h1>

      {/* Input field and button to add new todo */}
      <div className="input-section">
        <input
          value={newTask}
          onChange={(e) => setNewTask(e.target.value)}
          placeholder="Enter task description"
          className="task-input"
        />
        <button onClick={addTodo} className="add-button">Add</button>
      </div>

      {/* List of todo items */}
      <ul className="todo-list">
        {todos && todos.map((todo) => (
          <li
            key={todo._id}
            className={`todo-item ${todo.isComplete ? 'completed' : ''}`}
          >
            <div className="todo-left">
              {/* Toggle checkbox */}
              <input
                type="checkbox"
                checked={todo.isComplete}
                onChange={() => toggleTodo(todo)}
              />
              <span className="todo-text">
                {todo.description}
              </span>
            </div>
            {/* Delete button */}
            <button
              onClick={() => deleteTodo(todo._id)}
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
