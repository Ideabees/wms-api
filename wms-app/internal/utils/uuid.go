package utils

import ( 
    "github.com/google/uuid"
)


func CreateUUID() string{
	// Create UUID
	id := uuid.New() 
	return id.String()
}