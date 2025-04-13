package migrations

import (
	"database/sql"
	"fmt"
)

func InitiateDB(db *sql.DB) error {
	otpTableQuery := `
	CREATE TABLE IF NOT EXISTS OTPDETAILS (
	    id SERIAL PRIMARY KEY,
	    otp VARCHAR(100) NOT NULL,
	    mobile_number VARCHAR(100) UNIQUE NOT NULL,
	    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	// Execute the schema
	_, err := db.Exec(otpTableQuery)
	if err != nil {
		fmt.Println("Error while creating otp table", err)
		return err
	}

	tokenTableQuery := `
	CREATE TABLE IF NOT EXISTS TOKENDETAILS (
	    id SERIAL PRIMARY KEY,
	    token VARCHAR(100) NOT NULL,
	    mobile_number VARCHAR(100) UNIQUE NOT NULL,
	    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`	
	// Execute the schema
	_, tokenErr := db.Exec(tokenTableQuery)	
	if tokenErr != nil {
		fmt.Println("Error while creating token table", tokenErr)
		return tokenErr
	}

	//defer db.Close()
	return err
}