import React, { useState, useEffect } from 'react';

export default function ToDoList() {
    const [todoList, setTodoList] = useState([]);

    // 获取任务列表
    const fetchTodos = async () => {
        const res = await fetch("http://localhost:8080/todos");
        const data = await res.json();
        setTodoList(data);
    };

    useEffect(() => {
        fetchTodos();
    }, []);

    // 更新任务完成状态
    const toggleComplete = async (todo) => {
        const updated = {
            ...todo,
            isComplete: !todo.isComplete,
        };

        await fetch(`http://localhost:8080/todos/${todo.id}`, {
            method: "PUT",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(updated),
        });

        fetchTodos(); // 重新获取最新数据
    };

    // 删除任务
    const deleteTodo = async (id) => {
        await fetch(`http://localhost:8080/todos/${id}`, {
            method: "DELETE",
        });

        fetchTodos();
    };

    if (!todoList.length) {
        return <p>"There are no to-do items!"</p>;
    }

    return (
        <ul>
            {todoList.map((item) => (
                <div
                    key={item.id}
                    className={item.isComplete ? 'todolist_done' : 'todolist_undo'}
                >
                    <li>
                        <input
                            type="checkbox"
                            checked={item.isComplete}
                            onChange={() => toggleComplete(item)}
                        />
                        {item.description}
                        {item.isComplete && <span> (Done!)</span>}
                        <button onClick={() => deleteTodo(item.id)}>Delete</button>
                    </li>
                </div>
            ))}
        </ul>
    );
}
