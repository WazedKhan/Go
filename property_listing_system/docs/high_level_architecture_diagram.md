## High-Level Architecture Diagram

```
                           +----------------------+
                           |  Web / Mobile Client |
                           +----------------------+
                                     |
                                     ▼
                          +--------------------+
                          |    API Gateway     |  <-- Handles authentication & request routing
                          +--------------------+
                                     |
    -----------------------------------------------------------------------------
    |                   |                 |                  |                  |
    ▼                   ▼                 ▼                  ▼                  ▼
+----------------+  +----------------+  +----------------+  +----------------+  +-------------------+
| User Service  |  | Property Service | | Search Service | | Inquiry Service | | Review Service     |
| (Auth, Roles) |  | (Listings CRUD)  | | (Filters, DB)  | | (Messaging)     | | (Ratings, Reviews) |
+----------------+  +----------------+  +----------------+  +----------------+  +-------------------+
       |                     |                |                  |                 |
       ▼                     ▼                ▼                  ▼                 ▼
+------------------------------------------------------------------------------------+
|                       PostgreSQL Database                                          |
|  (Users, Properties, Inquiries, Reviews, Search Index)                             |
+------------------------------------------------------------------------------------+
        |
        ▼
+-----------------+
|   Redis Cache   |  <-- Stores frequently accessed queries
+-----------------+
```

## Explanation of Each Component

### 1️⃣ Client Apps (Frontend)
- Users interact with the system via **Web & Mobile Apps**.
- Sends requests to the backend via the **API Gateway**.

### 2️⃣ API Gateway
- Acts as a single entry point for all requests.
- Handles **authentication & authorization** (JWT / OAuth).
- Routes requests to appropriate **microservices**.

### 3️⃣ Microservices (Core Business Logic)
- **User Service** → Manages authentication, roles (Admin, Owner, Renter).
- **Property Service** → Manages property **CRUD operations**.
- **Search Service** → Provides **fast search & filtering** (can use PostgreSQL indexes or Elasticsearch if needed).
- **Inquiry Service** → Manages **messages between renters & owners**.
- **Review Service** → Handles property **ratings & reviews**.

### 4️⃣ Database Layer
- **PostgreSQL**: Stores structured data (Users, Listings, Reviews, Inquiries).
- **Redis**: Caches frequently accessed data (e.g., popular properties).

