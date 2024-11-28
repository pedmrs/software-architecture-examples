# Hexagonal Application

This project implements a CRUD (Create, Read, Update, Delete) application for managing **Person** and **Address** entities using **Hexagonal Architecture** (also known as **Ports and Adapters**) in Go. This structure is a design pattern that separates the core business logic from external concerns (such as web frameworks or databases) using **Ports** and **Adapters**. The application core is completely decoupled from infrastructure and can easily be adapted for different external interfaces (e.g., HTTP, CLI, etc.).

- **Advantages**:
  - High testability: Business logic can be tested independently of external concerns.
  - Decoupling: Core business logic is separated from the infrastructure.
  - Flexibility: Adapters can be swapped out without changing the core business logic.
  - Easier maintenance as it avoids direct dependencies on external systems (e.g., database, APIs).

- **Disadvantages**:
  - Higher complexity due to the decoupling of components.
  - More boilerplate code required.
  - Can be overkill for small applications.

## Folder Structure

```text
hexagonal/
├── cmd/
│   └── server/
│       └── main.go: The entry point of the application. Wires up the repositories, services, and handlers, and starts the HTTP server.
├── internal/
│   ├── application/: Contains the services (business logic) that manage Person and Address.
│   │   ├── person_service.go
│   │   └── address_service.go
│   ├── domain/: The core domain models/entities (Person and Address).
│   │   ├── person.go
│   │   └── address.go
│   ├── port/: Defines the interfaces (ports) for the repositories.
│   │   ├── person_repository.go
│   │   └── address_repository.go
│   └── adapters/
│       ├── http/: Contains the HTTP handlers that expose the endpoints for interacting with the service layer.
│       │   ├── person_handler.go
│       │   └── address_handler.go
│       └── repository/:  The in-memory repository implementations that store Person and Address.
│           ├── in_memory_person_repository.go
│           └── in_memory_address_repository.go
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