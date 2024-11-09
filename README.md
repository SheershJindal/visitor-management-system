# Visitor Management System

Welcome to the **Visitor Management System** project! This system allows businesses to manage and track visitors coming into their premises, including registration, check-in, and check-out. This README outlines the basic project structure, setup instructions, and our commit message conventions.

## Table of Contents
- [Project Setup](#project-setup)
- [Technology Stack](#technology-stack)
- [Commit Message Conventions](#commit-message-conventions)
- [License](#license)

## Project Setup

To get started with the project, clone this repository and follow the instructions below to set up the development environment.

1. **Clone the repository**:
    ```bash
    git clone https://github.com/yourusername/visitor-management-system.git
    ```

2. **Install dependencies**:
    - **Frontend**: Navigate to the `frontend` folder and install the necessary packages.
    ```bash
    cd frontend
    npm install
    ```

    - **Backend**: Navigate to the `backend` folder and install the necessary packages.
    ```bash
    cd backend
    npm install
    ```

3. **Run the application** using `docker-compose`:
    If you're using Docker, you can start both the frontend and backend services with:
    ```bash
    docker-compose up --build
    ```

<!-- 4. **Run tests** (optional):
    - To run tests for the frontend and backend, you can use the following commands:
    ```bash
    cd frontend
    npm test
    cd ../backend
    npm test
    ``` -->

## Technology Stack

- **Frontend**: TBD
- **Backend**: GoLang
- **Database**: Postgres
- **Containerization**: Docker
<!-- - **CI/CD**: GitHub Actions -->

## Commit Message Conventions

We follow a simple and structured commit message convention to ensure our version history is clear and organized. Here's how to format commit messages:

### Format

```
<type> | <scope> | <short description of the change>
```

- **`<type>`**: The type of change (e.g., `CHORE`, `FEAT`, `FIX`, `CI`, etc.)
- **`<scope>`**: The component or part of the project that the commit pertains to (e.g., `FRONTEND`, `BACKEND`, `DOCKER`, `CI`, `COMMON`).
- **`<short description>`**: A brief, descriptive message in the present tense.

### Types of Commits

- **`CHORE`**: Routine tasks, configuration, and setup.
- **`FEAT`**: New features or functionality.
- **`FIX`**: Bug fixes.
- **`CI`**: Changes to the continuous integration pipeline.
- **`DOCS`**: Documentation changes.
- **`REFACTOR`**: Code restructuring and optimization without changing functionality.
- **`PERF`**: Performance improvements.
- **`WIP`**: Work in progress (used for unfinished or incomplete work).

### Example Commit Messages

#### 1. **Initial Setup Commits**

These commits set up the foundational structure of the project, including configurations, Docker setup, and basic structure.

```bash
CHORE | COMMON | Init the project structure
CHORE | COMMON | Initialize React app with create-react-app
CHORE | COMMON | Initialize Node.js backend with Express
CHORE | COMMON | Add Dockerfile for frontend and backend services
CHORE | COMMON | Add .gitignore for Node.js, React, and Docker
CI | COMMON | Add initial GitHub Actions CI workflow
DOCS | COMMON | Add initial README with project description
```

#### 2. **Adding Features**

Once the basic setup is done, we start adding features to the system. For example, adding a visitor registration form or a check-in/check-out API.

```bash
FEAT | FRONTEND | Add visitor registration form component
FEAT | BACKEND | Implement visitor check-in API
FEAT | BACKEND | Add database schema for visitors
FEAT | FRONTEND | Display visitor list in dashboard
```

#### 3. **Bug Fixes and Improvements**

Whenever a bug is fixed or improvements are made, we follow the `FIX` type.

```bash
FIX | FRONTEND | Fix layout issue with visitor registration form
FIX | BACKEND | Correct database connection error in check-out API
```

#### 4. **Continuous Integration (CI) Setup**

If you modify the CI/CD pipeline, such as adding a step to deploy the system, you would use the `CI` type.

```bash
CI | COMMON | Add build step for frontend in CI pipeline
CI | COMMON | Configure Docker container for production build
```

#### 5. **Documentation Changes**

Changes to the `README.md` or other documentation files would be committed using the `DOCS` type.

```bash
DOCS | COMMON | Update README with setup instructions for backend
DOCS | FRONTEND | Document visitor registration form component usage
```

#### 6. **Code Refactoring**

If you make changes to improve or restructure existing code without changing its external behavior, use the `REFACTOR` type.

```bash
REFACTOR | FRONTEND | Refactor visitor form for better code readability
REFACTOR | BACKEND | Simplify visitor check-in API logic
```

#### 7. **Performance Improvements**

For commits related to optimizing performance, such as reducing load times or optimizing queries, use the `PERF` type.

```bash
PERF | BACKEND | Optimize visitor search query to improve response time
PERF | FRONTEND | Improve form rendering speed
```

#### 8. **Work In Progress (WIP)**

Use `WIP` for commits that represent unfinished work or ongoing development. These should be temporary commits and should be replaced with a more appropriate commit type when the work is complete.

```bash
WIP | FRONTEND | Start implementing visitor registration form (unfinished)
WIP | BACKEND | Begin implementing visitor check-out API (unfinished)
```

---

### Summary of Commit Types and Scopes

| Commit Type | Scope (Example) | Description |
|-------------|-----------------|-------------|
| **`CHORE`** | `COMMON`, `FRONTEND`, `BACKEND` | Non-functional changes like configuration, project structure, Docker setup, etc. |
| **`FEAT`**  | `FRONTEND`, `BACKEND` | Adding new features, e.g., new components or APIs |
| **`FIX`**   | `FRONTEND`, `BACKEND` | Bug fixes or small corrections |
| **`CI`**    | `COMMON` | Changes to the CI/CD pipeline |
| **`DOCS`**  | `COMMON`, `FRONTEND`, `BACKEND` | Documentation updates, such as adding or updating the README |
| **`REFACTOR`** | `FRONTEND`, `BACKEND` | Code restructuring or optimization without changing external functionality |
| **`PERF`**  | `FRONTEND`, `BACKEND` | Performance optimizations, such as faster queries or improved rendering |
| **`WIP`**   | `COMMON`, `FRONTEND`, `BACKEND` | Work in progress; used for unfinished or incomplete tasks |

---

By following this convention, our commit history remains organized, making it easier for everyone to understand what changes were made, why they were made, and what part of the project they affect. This helps maintain consistency and clarity as the project grows.

---

### Notes:
- **Keep commit messages short and to the point**. If further detail is required, use the body of the commit message (after the short description).
- **Be consistent**: Ensure your commit messages follow the same format and conventions throughout the project.

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.