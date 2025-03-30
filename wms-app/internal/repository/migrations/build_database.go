package migrations

import "database/sql"

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
	//defer db.Close()
	return err
}