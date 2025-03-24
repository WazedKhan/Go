# 📌 Service Interactions
Each microservice will communicate via GraphQL API, following these interaction flows:

## 1️⃣ User Registration & Authentication
- Client → API Gateway → User Service
- User registers or logs in using email/password or OAuth.
- User service returns JWT token for authentication.

## 2️⃣ Property Listing CRUD
- Client → API Gateway → Property Service
- Owners can add, update, delete their properties.
- Property details are stored in PostgreSQL.

## 3️⃣ Searching for Properties
- Client → API Gateway → Search Service
- Renters can search properties by location, price, type.
- Search Service queries PostgreSQL (or Elasticsearch for full-text search).

## 4️⃣ Inquiry & Messaging
- Client → API Gateway → Inquiry Service
- Renters send messages to property owners.
- Inquiry service manages message history.

## 5️⃣ Reviews & Ratings
- Client → API Gateway → Review Service
- Renters leave reviews & ratings for properties.
- Review data is stored in PostgreSQL.
