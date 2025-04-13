package repository

import (
	"fmt"
	"wms-app/config"
)

func SetToken(token string, mobile_number string) (string, error) {
	conn, err := config.GetDBConnection()
	if err != nil {
		fmt.Println("Error while making connection for insert token", err)
		return "", err
	}

	insertTokenQuery := fmt.Sprintf("INSERT INTO TOKENDETAILS (token, mobile_number) VALUES (%s, %s)", token, mobile_number)
	_, err = conn.Exec(insertTokenQuery)
	if err != nil {
		fmt.Println("Error while inserting token", err)
		return "", err
	}

	fmt.Println("Token successfully inserted!")
	return token, nil
}

func GetToken(mobile_number string) (string, error) {
	conn, err := config.GetDBConnection()
	if err != nil {
		fmt.Println("Error while making connection for get token", err)
		return "", err
	}

	var token string
	err = conn.QueryRow("SELECT token FROM TOKENDETAILS WHERE mobile_number = $1", mobile_number).Scan(&token)
	if err != nil {
        fmt.Println("Not able fetch otp", err)
		return "", err
    }
	
	fmt.Println("OTP successfully fetched!", token)
	return token, nil
}

