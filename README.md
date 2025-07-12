docker compose build
docker compose up
curl --location --request DELETE 'http://localhost:8080/v1/api/customers' \

# WhatsApp Management System API (wms-api)

This API provides endpoints for managing WhatsApp customers and users. Below are instructions for setup and usage.

---

## Getting Started

### Build and Run with Docker

```powershell
docker compose build
docker compose up
```

---

## API Endpoints

### 1. Register User

**Endpoint:** `POST /api/register`

**Request Example:**
```json
{
  "first_name": "kripat2",
  "middle_name": "shankar",
  "last_name": "sharma",
  "email": "test1@gmail.com",
  "password": "test1@123456",
  "confirm_password": "test1@123456",
  "mobile_number": "8956432134"
}
```

**Curl Example:**
```powershell
curl --location 'http://localhost:8080/api/register' `
--header 'Content-Type: application/json' `
--data-raw '{
    "first_name": "kripat2",
    "middle_name": "shankar",
    "last_name": "sharma",
    "email": "test1@gmail.com",
    "password": "test1@123456",
    "confirm_password": "test1@123456",
    "mobile_number": "8956432134"
}'
```

**Response:**
```json
{
  "data": {
    "UserId": "...",
    "FirstName": "kripat2",
    "MiddleName": "shankar",
    "LastName": "sharma",
    "Email": "test1@gmail.com",
    "MobileNumber": "8956432134",
    "CreatedAt": "...",
    "UpdatedAt": "..."
  },
  "message": "User registered successfully",
  "token": "..."
}
```

---

### 2. Login

**Endpoint:** `POST /api/login`

**Request Example:**
```json
{
  "email": "test1@gmail.com",
  "password": "test1@123456"
}
```

**Curl Example:**
```powershell
curl --location 'http://localhost:8080/api/login' `
--header 'Content-Type: application/json' `
--data-raw '{
    "email": "test1@gmail.com",
    "password": "test1@123456"
}'
```

**Response:**
```json
{
  "data": {
    "UserId": "...",
    "FirstName": "kripat2",
    "MiddleName": "shankar",
    "LastName": "sharma",
    "Email": "test1@gmail.com",
    "MobileNumber": "8956432134",
    "CreatedAt": "...",
    "UpdatedAt": "..."
  },
  "message": "Login successful",
  "token": "..."
}
```

---

### 3. Get Permissions

**Endpoint:** `GET /v1/api/permissions`

**Headers:**
```
Authorization: Bearer <your_token>
```

**Curl Example:**
```powershell
curl --location 'http://localhost:8080/v1/api/permissions' `
--header 'Authorization: Bearer <your_token>'
```

**Response:**
```json
{
  "email": "k3@gmail.com",
  "message": "Protected profile access",
  "user_id": "..."
}
```

---

### 4. Logout

**Endpoint:** `POST /v1/api/logout`

**Headers:**
```
Authorization: Bearer <your_token>
```

**Curl Example:**
```powershell
curl --location --request POST 'http://localhost:8080/v1/api/logout' `
--header 'Authorization: Bearer <your_token>'
```

**Response:**
```json
{
  "message": "Logged out successfully"
}
```

---

### 5. Add Customer (Single)

**Endpoint:** `POST /v1/api/customer`

**Headers:**
```
Authorization: Bearer <your_token>
Content-Type: application/json
```

**Request Example:**
```json
{
  "first_name": "c3",
  "last_name": "sharma",
  "mobile_number": "78934566764"
}
```

**Curl Example:**
```powershell
curl --location 'http://localhost:8080/v1/api/customer' `
--header 'Content-Type: application/json' `
--header 'Authorization: Bearer <your_token>' `
--data '{
    "first_name": "c3",
    "last_name": "sharma",
    "mobile_number": "78934566764"
}'
```

**Response:**
```json
{
  "message": "Customer successfully created",
  "status": "Success"
}
```

---

### 6. Get Customers

**Endpoint:** `GET /v1/api/customers`

**Headers:**
```
Authorization: Bearer <your_token>
```

**Curl Example:**
```powershell
curl --location 'http://localhost:8080/v1/api/customers' `
--header 'Authorization: Bearer <your_token>'
```

**Response:**
```json
{
  "data": [
    {
      "CustomerId": "...",
      "FirstName": "c1",
      "LastName": "sharma",
      "MobileNumber": "78934566760",
      "CreatedBy": " ",
      "UpdatedOn": "..."
    }
    // ...more customers
  ],
  "message": "",
  "status": "Success"
}
```

---

### 7. Delete Customers (Bulk)

**Endpoint:** `DELETE /v1/api/customers`

**Headers:**
```
Authorization: Bearer <your_token>
Content-Type: application/json
```

**Request Example:**
```json
{
  "customer_ids": ["id1", "id2"]
}
```

**Curl Example:**
```powershell
curl --location --request DELETE 'http://localhost:8080/v1/api/customers' `
--header 'Content-Type: application/json' `
--header 'Authorization: Bearer <your_token>' `
--data '{
    "customer_ids": ["id1", "id2"]
}'
```

**Response:**
```json
{
  "message": "Customer successfully deleted",
  "status": "Success"
}
```

---

### 8. Bulk Customer Creation via Excel

**Endpoint:** `POST /v1/api/create_bulk_customers`

**Headers:**
```
Authorization: Bearer <your_token>
```

**Form Data:**
```
file=@"C:/KripaDev/Coding/wms-api/Bulk-customer-data.xlsx"
```

**Curl Example (Windows):**
```powershell
curl --location 'http://localhost:8080/v1/api/create_bulk_customers' `
--header 'Authorization: Bearer <your_token>' `
--form 'file=@"C:/KripaDev/Coding/wms-api/Bulk-customer-data.xlsx"'
```

---

# Notes

- Replace `<your_token>` with your actual JWT token.
- All endpoints require authentication except registration and login.
- For bulk operations, ensure your Excel/CSV file matches the required format.