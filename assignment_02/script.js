document.addEventListener('DOMContentLoaded', () => {
    const todoForm = document.getElementById('todo-form');
    const taskInput = document.getElementById('task-input');
    const taskList = document.getElementById('task-list');

    // Function to create a new task item
    function createTaskElement(taskText) {
        const listItem = document.createElement('li');

        const taskSpan = document.createElement('span');
        taskSpan.textContent = taskText;
        listItem.appendChild(taskSpan);

        const actionsDiv = document.createElement('div');
        actionsDiv.classList.add('task-actions');

        const completeButton = document.createElement('button');
        completeButton.textContent = 'Complete';
        completeButton.classList.add('complete-btn');
        completeButton.addEventListener('click', () => {
            listItem.classList.toggle('completed');
        });
        actionsDiv.appendChild(completeButton);

        const deleteButton = document.createElement('button');
        deleteButton.textContent = 'Delete';
        deleteButton.classList.add('delete-btn');
        deleteButton.addEventListener('click', () => {
            taskList.removeChild(listItem);
        });
        actionsDiv.appendChild(deleteButton);

        listItem.appendChild(actionsDiv);

        return listItem;
    }

    // Add new task functionality
    todoForm.addEventListener('submit', (event) => {
        event.preventDefault(); // Prevent form submission

        const taskText = taskInput.value.trim();

        if (taskText !== '') {
            const newTask = createTaskElement(taskText);
            taskList.appendChild(newTask);
            taskInput.value = ''; // Clear the input field
        }
    });
});