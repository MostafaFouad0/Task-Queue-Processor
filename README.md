# Task Queue Processor

## Overview

Task Queue Processor is a Go-based backend service that allows users to submit tasks via an API, which are then processed asynchronously in the background using worker goroutines. The system supports concurrent task execution, making it ideal for I/O-bound and CPU-bound operations like email sending Tasks are managed and stored in a relational database, and their statuses can be queried via API endpoints.

## Features

- Task Submission API: Submit tasks such as sending emails, processing images, or making API calls, in this project we are just sending Emails
- Worker Pool: Concurrency with worker goroutines that process tasks from a queue.
- Task Status Management: Monitor task status (e.g., pending, in-progress, Ok, failed) with real-time updates.
- Retry Mechanism: Automatically retry failed tasks up to a configurable limit.
- Database Persistence: Task metadata and statuses are stored in a database (PostgreSQL).

## Tech Stack

- Language: Go
- API Framework: Gin Web Framework
- Database: PostgreSQL

## Endpoints
**1 - GET /task/:id**

Retrieve the status of a task by its ID.

### Request

- Method `GET`
- URL `/task/{id}`
### Path Parameters
| Parameter | Type   | Description               |
|-----------|--------|---------------------------|
| `id`      | `String` | `Unique ID of the task `     |

### Response
- Status Code :`200 OK`
- Body : `JSON object containing the task status.`

**2 - POST /task/**

Submit a new task to the queue for processing.

### Request

- Method `POST`
- URL `/task/`
- Body `JSON object containing the task data to be processed.`

### Response
- Status Code :`201 Created`
- Body : `JSON object containing the task ID if submission is successful`

## Contributing

Feel free to contribute by submitting pull requests, reporting issues, or suggesting features!
