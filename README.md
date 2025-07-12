# wms-api
This api provides all the features related to whatsapp managment system

docker compose build

docker compose up

Register Endpoint 
Example:
curl --location 'http://localhost:8080/api/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name": "kripat2",
    "middle_name" : "shankar",
    "last_name": "sharma",
    "email": "test1@gmail.com",
    "password": "test1@123456",
    "confirm_password": "test1@123456",
    "mobile_number": "8956432134"
}'

Response:
{
    "data": {
        "UserId": "c7dc1caf-d8a2-4a40-8c25-f5b67e5a83be",
        "FirstName": "kripat2",
        "MiddleName": "shankar",
        "LastName": "sharma",
        "Email": "test1@gmail.com",
        "MobileNumber": "8956432134",
        "CreatedAt": "2025-04-27 06:55:37.453284953 +0000 UTC",
        "UpdatedAt": "2025-04-27 06:55:37.453284953 +0000 UTC"
    },
    "message": "User registered successfully",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QxQGdtYWlsLmNvbSIsImV4cCI6MTc0NTk5NjEzNywiZmlyc3ROYW1lIjoia3JpcGF0MiIsImxhc3ROYW1lIjoic2hhcm1hIiwidXNlcklkIjoiYzdkYzFjYWYtZDhhMi00YTQwLThjMjUtZjViNjdlNWE4M2JlIn0.dCaGOoC37CnGTTe3WRbjS42fcOOPmpkanun3pbinxac"
}


Login:
curl --location 'http://localhost:8080/api/login' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImszQGdtYWlsLmNvbSIsImV4cCI6MTc0NDkzMTI0OSwic3ViIjo0fQ.FuGV91C27_DaTb1GJgWYvzmQWDD9rEJ-CFgTo8NRPPw' \
--data-raw '{
    "email": "test1@gmail.com",
    "password": "test1@123456"
}'

Response:
{
    "data": {
        "UserId": "c7dc1caf-d8a2-4a40-8c25-f5b67e5a83be",
        "FirstName": "kripat2",
        "MiddleName": "shankar",
        "LastName": "sharma",
        "Email": "test1@gmail.com",
        "MobileNumber": "8956432134",
        "CreatedAt": "2025-04-27 06:55:37.453284 +0000 UTC",
        "UpdatedAt": "2025-04-27 06:55:37.453284 +0000 UTC"
    },
    "message": "Login successful",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QxQGdtYWlsLmNvbSIsImV4cCI6MTc0NTk5NjE3NCwiZmlyc3ROYW1lIjoia3JpcGF0MiIsImxhc3ROYW1lIjoic2hhcm1hIiwidXNlcklkIjoiYzdkYzFjYWYtZDhhMi00YTQwLThjMjUtZjViNjdlNWE4M2JlIn0.aWenIv5n4TJJOb0s2BEAEx6g_AezQXPzbK85XBTw0cE"
}

Permissions:
curl --location 'http://localhost:8080/v1/api/permissions' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImszQGdtYWlsLmNvbSIsImV4cCI6MTc0NDkzMjExOSwic3ViIjo0fQ.iTROg462q8b0xIdhoaNijAIre_4Gj0hYFR_9u9SKJKw' \
--data ''

Response:
{
    "email": "k3@gmail.com",
    "message": "Protected profile access",
    "user_id": ""
}

Logout:
curl --location --request POST 'http://localhost:8080/v1/api/logout' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImszQGdtYWlsLmNvbSIsImV4cCI6MTc0NDkzMjExOSwic3ViIjo0fQ.iTROg462q8b0xIdhoaNijAIre_4Gj0hYFR_9u9SKJKw'

Response:
{
    "message": "Logged out successfully"
}

Add customer one by one:
curl --location 'http://localhost:8080/v1/api/customer' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QxQGdtYWlsLmNvbSIsImV4cCI6MTc0NTk5NjE3NCwiZmlyc3ROYW1lIjoia3JpcGF0MiIsImxhc3ROYW1lIjoic2hhcm1hIiwidXNlcklkIjoiYzdkYzFjYWYtZDhhMi00YTQwLThjMjUtZjViNjdlNWE4M2JlIn0.aWenIv5n4TJJOb0s2BEAEx6g_AezQXPzbK85XBTw0cE' \
--data '{
    "first_name": "c3",
    "last_name" : "sharma",
    "mobile_number": "78934566764"
}'

Response:
{
    "message": "Customer succefully created",
    "status": "Success"
}

{
    "message": "DB insertion Failed",
    "status": "Failed"
}

Get Customers:
curl --location 'http://localhost:8080/v1/api/customers' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QxQGdtYWlsLmNvbSIsImV4cCI6MTc0NTk5NjE3NCwiZmlyc3ROYW1lIjoia3JpcGF0MiIsImxhc3ROYW1lIjoic2hhcm1hIiwidXNlcklkIjoiYzdkYzFjYWYtZDhhMi00YTQwLThjMjUtZjViNjdlNWE4M2JlIn0.aWenIv5n4TJJOb0s2BEAEx6g_AezQXPzbK85XBTw0cE'

Response:
{
    "data": [
        {
            "CustomerId": "f290a2a8-6482-4aef-b4e3-f7e506594be1",
            "FirstName": "c1",
            "LastName": "sharma",
            "MobileNumber": "78934566760",
            "CreatedBy": " ",
            "UpdatedOn": "2025-04-27 06:57:41.400123 +0000 UTC"
        },
        {
            "CustomerId": "3ff5ab7e-81c0-45ea-956b-fb944fd7bbc2",
            "FirstName": "c1",
            "LastName": "sharma",
            "MobileNumber": "78934566761",
            "CreatedBy": " ",
            "UpdatedOn": "2025-04-27 06:57:46.156479 +0000 UTC"
        },
        {
            "CustomerId": "a72df95b-0738-467a-84c7-54aa4b04bffb",
            "FirstName": "c2",
            "LastName": "sharma",
            "MobileNumber": "78934566762",
            "CreatedBy": " ",
            "UpdatedOn": "2025-04-27 06:58:04.547567 +0000 UTC"
        },
        {
            "CustomerId": "0bfa33f5-3420-4464-a668-84377c7bd4e7",
            "FirstName": "c3",
            "LastName": "sharma",
            "MobileNumber": "78934566764",
            "CreatedBy": " ",
            "UpdatedOn": "2025-04-27 06:58:20.775889 +0000 UTC"
        }
    ],
    "message": "",
    "status": "Success"
}

Delete Customers:

curl --location --request DELETE 'http://localhost:8080/v1/api/customers' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QxQGdtYWlsLmNvbSIsImV4cCI6MTc0NTk5NjE3NCwiZmlyc3ROYW1lIjoia3JpcGF0MiIsImxhc3ROYW1lIjoic2hhcm1hIiwidXNlcklkIjoiYzdkYzFjYWYtZDhhMi00YTQwLThjMjUtZjViNjdlNWE4M2JlIn0.aWenIv5n4TJJOb0s2BEAEx6g_AezQXPzbK85XBTw0cE' \
--data '{
    "customer_ids":["0bfa33f5-3420-4464-a668-84377c7bd4e7", "a72df95b-0738-467a-84c7-54aa4b04bffb"]
}'

Response:
{
    "message": "Customer succefully deleted",
    "status": "Success"
}

## Bulk Customer Creation via Excel

Endpoint:
```
POST http://localhost:8080/v1/api/create_bulk_customers
```

Headers:
```
Authorization: Bearer <your_token>
```

Form Data:
```
file=@"C:/KripaDev/Coding/wms-api/Bulk-customer-data.xlsx"
```

Example using curl (Windows path):
```powershell
curl --location 'http://localhost:8080/v1/api/create_bulk_customers' `
--header 'Authorization: Bearer <your_token>' `
--form 'file=@"C:/KripaDev/Coding/wms-api/Bulk-customer-data.xlsx"'
```