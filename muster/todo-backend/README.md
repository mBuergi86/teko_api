# todo-backend
This repository contains a REST API backend for managing todos. The backend is based on the [HttpRouter](https://github.com/julienschmidt/httprouter).

# Data model
## Todo
| Field name  | Data type |
|-------------|-----------|
| Id          | string    |
| Title       | string    |
| Description | string    |
| Terminated  | bool      |

## JsonExtendedResponse
| Field name | Data type   |
|------------|-------------|
| Meta       | interface{} |
| Data       | interface{} |

## JsonDataResponse
| Field name | Data type |
|------------|-----------|
| Data       | []Todo    |

## ApiError
| Field name | Data type |
|------------|-----------|
| Status     | int16     |
| Title      | string    |

## JsonErrorResponse
| Field name  | Data type |
|-------------|-----------|
| Error       | ApiError  |

# Features
The following endpoints are implemented:

| No. | HTTP Verb | Path       | Expects (JSON) | Returns (JSON)                 | HTTP Status                         | Description         |
|-----|-----------|------------|----------------|--------------------------------|-------------------------------------|---------------------|
| 1   | GET       | /          | Nothing        | Welcome string                 | 200 (success)                       | Welcome string      |
| 2   | GET       | /todos     | Nothing        | An array with todo entries     | 200 (success)                       | Get a list of todos |
| 3   | GET       | /todos/:id | Nothing        | The todo with the specified ID | 200 (success) or 404 (not found)    | Get todo by ID      |
| 4   | POST      | /todos     | A todo entry   | The new todo entry             | 201 (created) or 400 (Bad Request)  | Create new todo     |

# Usage
The project can be built to desired platform and then the backend runs on port 8080, listening on any interface available at the place of execution.