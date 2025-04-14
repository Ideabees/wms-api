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
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImszQGdtYWlsLmNvbSIsImV4cCI6MTc0NDkzMTI0OSwic3ViIjo0fQ.FuGV91C27_DaTb1GJgWYvzmQWDD9rEJ-CFgTo8NRPPw"
}

Login:
curl --location 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImszQGdtYWlsLmNvbSIsImV4cCI6MTc0NDkzMTI0OSwic3ViIjo0fQ.FuGV91C27_DaTb1GJgWYvzmQWDD9rEJ-CFgTo8NRPPw' \
--data-raw '{
    "email": "k3@gmail.com",
    "password": "k@123456"

}'

Response:
{
    "message": "Login successful",
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