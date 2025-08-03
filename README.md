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

## API Documentation (Swagger)

Swagger UI is available for interactive API documentation and testing.

**Access Swagger UI:**
```
http://localhost:8080/swagger/index.html
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
```json
{
    "status_code": 409,
    "error": "DuplicateError",
    "message": "User with this email or mobile number already exists"
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
    "last_name" : "sharma",
    "mobile_number": "78934566753",
    "email_id": "k3@gmail.com",
    "city": "bsr",
    "pincode": "203001"
}
```

**Curl Example:**
```powershell
curl --location 'http://localhost:8080/v1/api/customer' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QxQGdtYWlsLmNvbSIsImV4cCI6MTc1MjY2ODg4NSwiZmlyc3ROYW1lIjoia3JpcGF0MiIsImxhc3ROYW1lIjoic2hhcm1hIiwidXNlcklkIjoiYzdkYzFjYWYtZDhhMi00YTQwLThjMjUtZjViNjdlNWE4M2JlIn0.vHTOO_sp-Ho1-vrtL0QSYyB7uME48BZOCp-DPvWMlgs' \
--header 'Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QxQGdtYWlsLmNvbSIsImV4cCI6MTc1MjY2ODg4NSwiZmlyc3ROYW1lIjoia3JpcGF0MiIsImxhc3ROYW1lIjoic2hhcm1hIiwidXNlcklkIjoiYzdkYzFjYWYtZDhhMi00YTQwLThjMjUtZjViNjdlNWE4M2JlIn0.vHTOO_sp-Ho1-vrtL0QSYyB7uME48BZOCp-DPvWMlgs' \
--data-raw '{
    "first_name": "c3",
    "last_name" : "sharma",
    "mobile_number": "78934566753",
    "email_id": "k3@gmail.com",
    "city": "bsr",
    "pincode": "203001"
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
            "CustomerId": "e520cb0b-76c3-4028-9a80-a8cbec90fb53",
            "FirstName": "test104",
            "LastName": "Sharma",
            "MobileNumber": "784352655",
            "EmailID": "t104@gmail.com",
            "City": "BLR",
            "Pincode": "201005",
            "CreatedBy": " ",
            "UpdatedOn": "2025-07-13 12:35:30.024531 +0000 UTC"
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
curl --location 'http://localhost:8080/v1/api/create_bulk_customers' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QxQGdtYWlsLmNvbSIsImV4cCI6MTc1MjY2ODg4NSwiZmlyc3ROYW1lIjoia3JpcGF0MiIsImxhc3ROYW1lIjoic2hhcm1hIiwidXNlcklkIjoiYzdkYzFjYWYtZDhhMi00YTQwLThjMjUtZjViNjdlNWE4M2JlIn0.vHTOO_sp-Ho1-vrtL0QSYyB7uME48BZOCp-DPvWMlgs' \
--header 'Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QxQGdtYWlsLmNvbSIsImV4cCI6MTc1MjY2ODg4NSwiZmlyc3ROYW1lIjoia3JpcGF0MiIsImxhc3ROYW1lIjoic2hhcm1hIiwidXNlcklkIjoiYzdkYzFjYWYtZDhhMi00YTQwLThjMjUtZjViNjdlNWE4M2JlIn0.vHTOO_sp-Ho1-vrtL0QSYyB7uME48BZOCp-DPvWMlgs' \
--form 'file=@"/C:/KripaDev/Coding/wms-api/Bulk-customer-data.xlsx"'
```
**Response:**
```json
{
    "failed_count": 0,
    "message": "Bulk upload completed",
    "success_count": 4
}
```
---

### 9. Send Chat Message

**Endpoint:** `POST /v1/api/chat/chats/send/message`

**Headers:**
```
Authorization: Bearer <your_token>
Content-Type: application/json
```

**Request Example:**
```json
{
    "receiver_mobile_number": "917892360471",
    "message": "Kripa, we have an exciting offer for you."
}
```

**Curl Example:**
```powershell
curl --location 'http://localhost:8080/v1/api/chat/chats/send/message' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImtyaXBhdGVzdDEwMkBnbWFpbC5jb20iLCJleHAiOjE3NTQ1MDA3MDQsImZpcnN0TmFtZSI6ImtyaXBhdGVzdDEwMiIsImxhc3ROYW1lIjoic2hhcm1hIiwidXNlcklkIjoiMGU5MDUyOWUtM2IwMi00YjI0LThkY2EtNDQzNmNkOThjMjgzIn0.npsqLIN6VS7QmkTGj_wI20NgWLa3mETpSOXbCtMwWqQ' \
--header 'Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImtyaXBhdGVzdDEwMkBnbWFpbC5jb20iLCJleHAiOjE3NTQ1MDA3MDQsImZpcnN0TmFtZSI6ImtyaXBhdGVzdDEwMiIsImxhc3ROYW1lIjoic2hhcm1hIiwidXNlcklkIjoiMGU5MDUyOWUtM2IwMi00YjI0LThkY2EtNDQzNmNkOThjMjgzIn0.npsqLIN6VS7QmkTGj_wI20NgWLa3mETpSOXbCtMwWqQ' \
--data '{
    "receiver_mobile_number":"917892360471",
    "message": "Kripa, we have an exciting offer for you."
}'
```

**Response:**
```json
{
    "message": "Message sent successfully"
}
```

---

# Notes

- Replace `<your_token>` with your actual JWT token.
- All endpoints require authentication except registration and login.
- For bulk operations, ensure your Excel/CSV file matches the required format.