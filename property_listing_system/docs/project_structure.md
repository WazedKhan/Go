## Project Structure for Property Listing System

Following **Clean Architecture** principles, the project is structured to separate concerns properly.

```
📂 property-listing-backend/   # Root directory
├── 📂 cmd/                     # Entry points for the application (main.go)
│   ├── main.go                 # Initializes the server and dependencies
│   └── config.go               # Configuration settings (env, ports, database, etc.)
│
├── 📂 internal/                 # Core business logic (not exposed to other modules)
│   ├── 📂 domain/               # Entities & core business logic
│   │   ├── user.go              # User entity
│   │   ├── property.go          # Property entity
│   │   ├── review.go            # Review entity
│   │   └── inquiry.go           # Inquiry entity
│   │
│   ├── 📂 usecase/              # Application use cases (services layer)
│   │   ├── user_service.go      # Business logic for user management
│   │   ├── property_service.go  # Business logic for properties
│   │   ├── review_service.go    # Business logic for reviews
│   │   └── inquiry_service.go   # Business logic for inquiries
│   │
│   ├── 📂 repository/           # Data access layer
│   │   ├── user_repo.go         # User repository (Postgres queries)
│   │   ├── property_repo.go     # Property repository
│   │   ├── review_repo.go       # Review repository
│   │   └── inquiry_repo.go      # Inquiry repository
│
├── 📂 graph/                    # GraphQL Schema & Resolvers
│   ├── schema.graphqls          # GraphQL schema definitions
│   ├── resolver.go              # Resolver implementations
│   ├── models_gen.go            # Auto-generated models from schema
│   ├── generated.go             # Auto-generated GraphQL types
│
├── 📂 infrastructure/           # Frameworks, databases, logging, and external services
│   ├── db/                      # Database setup
│   │   ├── postgres.go          # PostgreSQL connection setup
│   │   ├── redis.go             # Redis connection setup
│   │
│   ├── middleware/              # Custom middlewares (Auth, Logging, etc.)
│   │   ├── auth.go              # Authentication middleware
│   │   ├── logging.go           # Logging middleware
│
├── 📂 server/                   # HTTP & GraphQL Server
│   ├── routes.go                # GraphQL routes setup
│   ├── server.go                # Server initialization
│
├── 📂 config/                   # Configuration files (YAML, JSON, ENV)
│   ├── config.yaml              # App configuration file
│
├── 📂 scripts/                  # Helper scripts (DB migrations, seeding, etc.)
│   ├── migrate.sh               # Database migration script
│   ├── seed.sh                  # Seed database with initial data
│
├── go.mod                       # Go module file
├── go.sum                       # Dependencies lock file
├── Dockerfile                   # Dockerfile for containerization
├── README.md                    # Documentation
```

## 📌 Explanation of Key Directories

### 1️⃣ `cmd/` (Application Entry Point)
- Contains `main.go`, which initializes the **server**, **config**, and **dependencies**.
- `config.go` handles loading environment variables.

### 2️⃣ `internal/` (Business Logic)
- **`domain/`** → Contains core entities like `User`, `Property`, `Review`.
- **`usecase/`** → Implements business logic.
- **`repository/`** → Handles database queries.

### 3️⃣ `graph/` (GraphQL Schema & Resolvers)
- Defines GraphQL schema and **resolvers**.
- Uses `gqlgen` to generate Go models from the schema.

### 4️⃣ `infrastructure/` (Frameworks & External Services)
- Database connection setups (PostgreSQL, Redis, etc.).
- Custom middlewares for **authentication, logging, error handling**.

### 5️⃣ `server/` (GraphQL API Server)
- Initializes and runs the **GraphQL server** with Echo.

### 6️⃣ `config/` (Configuration Management)
- Stores environment and application configuration files.

### 7️⃣ `scripts/` (Database Migrations & Helpers)
- Scripts to **migrate and seed** the database.

