const API_URL = "http://localhost:8080";

// Fetch tasks and display them
async function fetchTasks() {
    const res = await fetch(`${API_URL}/tasks`);
    const tasks = await res.json();
    const taskList = document.getElementById("taskList");
    taskList.innerHTML = "";

    tasks.forEach(task => {
        const li = document.createElement("li");
        li.className = `list-group-item d-flex justify-content-between align-items-center`;
        li.innerHTML = `
            ${task.title} 
            <span>
                <button class="btn btn-info btn-sm" onclick="editTask(${task.id}, '${task.title}')">Edit</button>
                <button class="btn btn-danger btn-sm" onclick="deleteTask(${task.id})">Delete</button>
            </span>
        `;
        taskList.appendChild(li);
    });
}

// Add a new task
document.getElementById("taskForm").addEventListener("submit", async (e) => {
    e.preventDefault();
    const title = document.getElementById("title").value;

    await fetch(`${API_URL}/tasks`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ title, completed: false, user_id: 1 }),
    });

    document.getElementById("title").value = "";
    fetchTasks();
});

// Delete a task
async function deleteTask(id) {
    if (confirm("Are you sure you want to delete this task?")) {
        await fetch(`${API_URL}/tasks/${id}`, { method: "DELETE" });
        fetchTasks();
    }
}

// Edit a task
async function editTask(id, currentTitle) {
    const newTitle = prompt("Edit task title:", currentTitle);
    if (newTitle) {
        await fetch(`${API_URL}/tasks/${id}`, {
            method: "PUT",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ title: newTitle, completed: false }),
        });
        fetchTasks();
    }
}

// Load tasks on page load
fetchTasks();
