## **App Idea:**

1. **Prompt the User for Input:**
   - The program uses a `for` loop to continuously prompt the user with: 
     ```
     Enter a command (add, list, exit):
     ```

2. **Accept Commands like `add`, `list`, and `exit`:**
   - **`add`**: Allows the user to add a new task to the `tasks` slice.
   - **`list`**: Displays all tasks in the `tasks` slice in a numbered format.
   - **`exit`**: Terminates the application gracefully.

3. **Handle Invalid Commands:**
   - If the user enters a command that isn’t `add`, `list`, or `exit`, the program displays:
     ```
     Invalid command. Please use 'add', 'list', or 'exit'.
     ```

4. **Handle Adding Tasks:**
   - The `createTask` function prompts the user for a task description and ensures:
     - Input is trimmed of leading/trailing whitespace.
     - Empty task descriptions are rejected with an appropriate error message:
       ```
       Task description cannot be empty.
       ```

5. **Handle Listing Tasks:**
   - The `showTasks` function checks if the `tasks` slice is empty:
     - If empty, it displays:
       ```
       No tasks available!
       ```
     - Otherwise, it prints each task in a numbered list:
       ```
       [1]: Buy groceries
       [2]: Complete Go tutorial
       ```

6. **Graceful Exit:**
   - When the user enters `exit`, the program terminates with a farewell message:
     ```
     Exiting the application.
     ```

---

### **Code Quality Highlights:**

- **Readability**: Functions are small, well-named (`createTask`, `showTasks`), and focused on single responsibilities.
- **Validation**: Input is sanitized (trimmed) and validated (no empty tasks).
- **Error Messages**: Informative feedback is provided for invalid commands or empty task descriptions.
- **Scalability**: The program structure makes it easy to add more commands in the future.
- **Idiomatic Go**: Uses idiomatic constructs like slices, `for` loops, `switch` statements, and `range`.

---

### **Example Interaction:**

#### Input:
```
Welcome to the CLI To-Do List!
Enter a command (add, list, exit): add
Enter task description: Buy groceries
Task created successfully!

Enter a command (add, list, exit): add
Enter task description: Complete Go tutorial
Task created successfully!

Enter a command (add, list, exit): list
Your tasks:
[1]: Buy groceries
[2]: Complete Go tutorial

Enter a command (add, list, exit): exit
Exiting the application.
```

#### Output:
```
Welcome to the CLI To-Do List!
Enter a command (add, list, exit): add
Enter task description: Buy groceries
Task created successfully!

Enter a command (add, list, exit): add
Enter task description: Complete Go tutorial
Task created successfully!

Enter a command (add, list, exit): list
Your tasks:
[1]: Buy groceries
[2]: Complete Go tutorial

Enter a command (add, list, exit): exit
Exiting the application.
```

---

### **What You Learned from This Project:**

1. **Working with Slices:**
   - Adding items dynamically using `append`.
   - Iterating over slices using `range`.

2. **Input Handling:**
   - Using `bufio.Scanner` to read user input from the command line.
   - Validating and sanitizing user input with `strings.TrimSpace`.

3. **Conditional Logic:**
   - Using `switch` statements for command handling.
   - Providing meaningful feedback for valid and invalid inputs.

4. **Looping Mechanisms:**
   - Creating a continuous `for` loop for a CLI application.

5. **Functions and Modularity:**
   - Designing reusable functions for specific tasks (`createTask`, `showTasks`).

---

### Next Steps:

If you'd like to expand this project, here are some ideas:
1. **Mark Tasks as Completed:**
   - Add a `complete` command to mark tasks as done.
   - Show completed tasks with a different format (e.g., `[X]`).

2. **Delete Tasks:**
   - Add a `delete` command to remove tasks by their number.

3. **Persistent Storage:**
   - Save tasks to a file so they persist across program runs.

4. **Enhance UX:**
   - Provide additional commands like `help` to display instructions.
