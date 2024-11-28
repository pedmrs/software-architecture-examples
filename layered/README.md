# Layered Application

This project implements a simple CRUD (Create, Read, Update, Delete) application for managing **Person** and **Address** entities in Go using a **Layered** architecture. In this structure, the codebase is divided into distinct layers where each layer has a specific responsibility. 

Common layers include:
- **Presentation Layer**: Handles HTTP requests and responses.
- **Service Layer**: Contains business logic.
- **Repository Layer**: Responsible for data persistence.

- **Advantages**:
  - Clear separation of concerns.
  - Each layer is focused on one responsibility (e.g., data access vs. business logic).
  - Easier to test individual layers.
  - Improves maintainability and scalability for medium-sized applications.

- **Disadvantages**:
  - Can introduce extra complexity and boilerplate.
  - May lead to performance overhead with multiple layers calling each other.
  - Inflexible if components in different layers need to communicate directly.

## Folder Structure

```text
layered/
├── cmd/
│   └── server/
│       └── main.go: The application entry point. Starts the HTTP server.
├── internal/
│   ├── handler/: Contains HTTP handlers for routing and input/output handling.
│   │   ├── person_handler.go
│   │   └── address_handler.go
│   ├── service/: Business logic for processing Person and Address entities.
│   │   ├── person_service.go
│   │   └── address_service.go
│   └── repository/: Data access layer (in-memory repositories for this example).
│       ├── in_memory_person_repository.go
│       └── in_memory_address_repository.go
├── go.mod
└── README.md
```

Routes
Person CRUD

    POST /persons: Create a new person.
    GET /persons/{id}: Get a person by ID.
    PUT /persons/{id}: Update a person by ID.
    DELETE /persons/{id}: Delete a person by ID.

Address CRUD

    POST /addresses: Create a new address.
    GET /addresses/{id}: Get an address by ID.
    PUT /addresses/{id}: Update an address by ID.
    DELETE /addresses/{id}: Delete an address by ID.