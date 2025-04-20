# wms-api
This api provides all the features related to whatsapp managment system

docker compose build

docker compose up

Register Endpoint 
Example:
curl --location 'http://localhost:8080/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name": "kripa1",
	"middle_name" : "shankar",
	"last_name": "sharma",
	"email": "k3@gmail.com",
	"password": "k@123456",
	"confirm_password": "k@123456",
	"mobile_number": "8956432193"
}'

Response:
{
    "message": "User registered successfully",
    "data": {
        "UserId": "0cc4fb2c-6e45-4928-b405-1a25c8da9f76",
        "FirstName": "kripat1",
        "MiddleName": "shankar",
        "LastName": "sharma",
        "Email": "kt1@gmail.com",
        "MobileNumber": "8956432194",
        "CreatedAt": "2025-04-20 05:54:41.540165 +0000 UTC",
        "UpdatedAt": "2025-04-20 05:54:41.540165 +0000 UTC"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Imt0MUBnbWFpbC5jb20iLCJleHAiOjE3NDUzODc5NTEsInN1YiI6NX0.2KII0cWBM3n8-x_XqR9_I-D_ataz8akoyscQRmRMUDs"
}
}

Login:
curl --location 'http://localhost:8080/api/login' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImszQGdtYWlsLmNvbSIsImV4cCI6MTc0NDkzMTI0OSwic3ViIjo0fQ.FuGV91C27_DaTb1GJgWYvzmQWDD9rEJ-CFgTo8NRPPw' \
--data-raw '{
    "email": "kt1@gmail.com",
    "password": "kt1@123456"

}'

Response:
{
    "message": "Login successful",
     "data": {
        "UserId": "0cc4fb2c-6e45-4928-b405-1a25c8da9f76",
        "FirstName": "kripat1",
        "MiddleName": "shankar",
        "LastName": "sharma",
        "Email": "kt1@gmail.com",
        "MobileNumber": "8956432194",
        "CreatedAt": "2025-04-20 05:54:41.540165 +0000 UTC",
        "UpdatedAt": "2025-04-20 05:54:41.540165 +0000 UTC"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImszQGdtYWlsLmNvbSIsImV4cCI6MTc0NDkzMjExOSwic3ViIjo0fQ.iTROg462q8b0xIdhoaNijAIre_4Gj0hYFR_9u9SKJKw"
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