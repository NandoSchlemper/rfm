# RFM Transportes API Documentation

## Overview

This documentation describes the RFM Transportes API, which provides endpoints for managing driver information in the transportation system. The API follows RESTful principles and uses JSON for data exchange.

## Base URL

All API endpoints are relative to the base URL:

```
http://localhost:8080/api/v1
```

## Authentication

Authentication is required for accessing the API endpoints. The API uses JWT (JSON Web Tokens) for authentication.

Include the JWT token in the Authorization header of your requests:

```
Authorization: Bearer {your_jwt_token}
```

## Endpoints

### Drivers

#### Get All Drivers

Retrieves a list of all drivers in the system.

- **URL**: `/drivers`
- **Method**: `GET`
- **Response**: 200 OK
  
**Response Format**:
```json
[
  {
    "id": "string",
    "name": "string",
    "cpf": "string",
    "rg": "string",
    "license_category": "string",
    "license_number": "string",
    "phone": "string",
    "address": "string",
    "emergency_contact": "string",
    "birth_date": "string",
    "admission_date": "string",
    "created_at": "string",
    "updated_at": "string"
  }
]
```

#### Get Driver by ID

Retrieves a specific driver by ID.

- **URL**: `/drivers/{id}`
- **Method**: `GET`
- **URL Parameters**: `id=[string]` (required)
- **Response**: 200 OK

**Response Format**:
```json
{
  "id": "string",
  "name": "string",
  "cpf": "string",
  "rg": "string",
  "license_category": "string",
  "license_number": "string",
  "phone": "string",
  "address": "string",
  "emergency_contact": "string",
  "birth_date": "string",
  "admission_date": "string",
  "created_at": "string",
  "updated_at": "string"
}
```

**Error Responses**:
- **404 Not Found**: If the driver with the specified ID does not exist

#### Create Driver

Creates a new driver in the system.

- **URL**: `/drivers`
- **Method**: `POST`
- **Content-Type**: `application/json`
- **Response**: 201 Created

**Request Body**:
```json
{
  "name": "string",
  "cpf": "string",
  "rg": "string",
  "license_category": "string",
  "license_number": "string",
  "phone": "string",
  "address": "string",
  "emergency_contact": "string",
  "birth_date": "string",
  "admission_date": "string"
}
```

**Response Format**:
```json
{
  "id": "string",
  "name": "string",
  "cpf": "string",
  "rg": "string",
  "license_category": "string",
  "license_number": "string",
  "phone": "string",
  "address": "string",
  "emergency_contact": "string",
  "birth_date": "string",
  "admission_date": "string",
  "created_at": "string",
  "updated_at": "string"
}
```

**Error Responses**:
- **400 Bad Request**: If the request body is invalid or missing required fields

#### Update Driver

Updates an existing driver's information.

- **URL**: `/drivers/{id}`
- **Method**: `PUT`
- **URL Parameters**: `id=[string]` (required)
- **Content-Type**: `application/json`
- **Response**: 200 OK

**Request Body**:
```json
{
  "name": "string",
  "cpf": "string",
  "rg": "string",
  "license_category": "string",
  "license_number": "string",
  "phone": "string",
  "address": "string",
  "emergency_contact": "string",
  "birth_date": "string",
  "admission_date": "string"
}
```

**Response Format**:
```json
{
  "id": "string",
  "name": "string",
  "cpf": "string",
  "rg": "string",
  "license_category": "string",
  "license_number": "string",
  "phone": "string",
  "address": "string",
  "emergency_contact": "string",
  "birth_date": "string",
  "admission_date": "string",
  "created_at": "string",
  "updated_at": "string"
}
```

**Error Responses**:
- **400 Bad Request**: If the request body is invalid
- **404 Not Found**: If the driver with the specified ID does not exist

#### Delete Driver

Deletes a driver from the system.

- **URL**: `/drivers/{id}`
- **Method**: `DELETE`
- **URL Parameters**: `id=[string]` (required)
- **Response**: 204 No Content

**Error Responses**:
- **404 Not Found**: If the driver with the specified ID does not exist

## Usage Examples

### Getting All Drivers

```bash
curl -X GET \
  http://localhost:8080/api/v1/drivers \
  -H 'Authorization: Bearer your_jwt_token'
```

### Getting a Specific Driver

```bash
curl -X GET \
  http://localhost:8080/api/v1/drivers/1 \
  -H 'Authorization: Bearer your_jwt_token'
```

### Creating a New Driver

```bash
curl -X POST \
  http://localhost:8080/api/v1/drivers \
  -H 'Authorization: Bearer your_jwt_token' \
  -H 'Content-Type: application/json' \
  -d '{
    "name": "João Silva",
    "cpf": "123.456.789-10",
    "rg": "12.345.678-9",
    "license_category": "E",
    "license_number": "12345678901",
    "phone": "(11) 98765-4321",
    "address": "Rua das Flores, 123, São Paulo - SP",
    "emergency_contact": "(11) 12345-6789",
    "birth_date": "1980-01-01",
    "admission_date": "2022-01-01"
  }'
```

### Updating a Driver

```bash
curl -X PUT \
  http://localhost:8080/api/v1/drivers/1 \
  -H 'Authorization: Bearer your_jwt_token' \
  -H 'Content-Type: application/json' \
  -d '{
    "name": "João Silva",
    "cpf": "123.456.789-10",
    "rg": "12.345.678-9",
    "license_category": "E",
    "license_number": "12345678901",
    "phone": "(11) 98765-4321",
    "address": "Rua das Rosas, 456, São Paulo - SP",
    "emergency_contact": "(11) 12345-6789",
    "birth_date": "1980-01-01",
    "admission_date": "2022-01-01"
  }'
```

### Deleting a Driver

```bash
curl -X DELETE \
  http://localhost:8080/api/v1/drivers/1 \
  -H 'Authorization: Bearer your_jwt_token'
```

## Error Handling

The API returns standard HTTP status codes to indicate the success or failure of a request.

Common error codes:
- **400 Bad Request**: The request was malformed or invalid
- **401 Unauthorized**: Authentication failed or was not provided
- **404 Not Found**: The requested resource was not found
- **500 Internal Server Error**: An unexpected error occurred on the server

Error responses will include a JSON body with additional details:

```json
{
  "error": "string",
  "message": "string"
}
```

