# Task Queue Processor

## Overview

Task Queue Processor is a Go-based backend service that allows users to submit tasks via an API, which are then processed asynchronously in the background using worker goroutines. The system supports concurrent task execution, making it ideal for I/O-bound and CPU-bound operations like email sending Tasks are managed and stored in a relational database, and their statuses can be queried via API endpoints.
