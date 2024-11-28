# Monolithic Application

This project implements a simple CRUD (Create, Read, Update, Delete) application for managing **Person** and **Address** entities in Go using a **Monolithic** architecture. In this structure, all components (data access, business logic, HTTP handlers, etc.) are tightly coupled in one place.

## Architecture Overview

### Monolithic Architecture

A **Monolithic** architecture involves a single codebase that handles all aspects of the application, including HTTP routing, business logic, and data persistence. All logic and components are contained within a single package or project structure.

- **Advantages**:
  - Simple to implement.
  - Easy to understand and manage for small projects.
  - Requires less boilerplate code.
  - Fewer components, so faster initial development.

- **Disadvantages**:
  - Difficult to scale as the codebase grows.
  - Harder to test individual parts.
  - Changes in one part of the application can affect others (tight coupling).
  - Not flexible for adding new features or services.

## Folder Structure

```text
monolithic/
├── README.md
├── main.go: Contains all code for routing, handling HTTP requests, implementing business logic, and accessing the in-memory storage.
└── go.mod: Defines the Go module and its dependencies.
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