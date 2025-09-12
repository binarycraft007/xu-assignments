**Create a to-do list application**

1.  **Core Components:** The application must include a title, a form for inputting new tasks (containing an input field and an "Add" button), and an ordered list to display the to-do items.
2.  **Styling:** Design a set of modern and clean CSS styles. The main content area should be centered with a maximum width of 800px. The input field and buttons should be aesthetically pleasing. There must be spacing between list items. Implement visual feedback effects for mouse hover states on buttons and list items.
3.  **Add To-Do Item:** Implement the functionality to add a to-do item. When a user clicks the "Add" button, the content from the input field should be used to create a new list item and append it to the list. After adding, the input field should be cleared.
4.  **Mark as Complete and Delete:** Implement the ability to mark items as complete and to delete them. Each list item (li) should contain a "Complete" button and a "Delete" button. Clicking the "Complete" button should add a "completed" CSS class to the list item (which will be styled with a line-through). Clicking the "Delete" button should remove the list item from the list.
5.  **Filtering and Clearing:** Implement filtering functionality with the options: "All," "Unfinished," and "Completed." Additionally, include "Clear Completed" and "Clear All" buttons at the bottom of the list to perform these respective actions.

**Technical Requirements:**
*   **Frontend:** React
*   **Backend:** FastAPI
*   **Database:** SQLite

**Project Structure:**
The project should be organized into two main directories:
*   `backend`
*   `frontend`

**Database Table Design:**
Please design a database table for the to-do list application and provide the corresponding SQL statement.

