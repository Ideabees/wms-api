package repository

import (
	"fmt"
	"wms-app/config"
)



func SetOTP(otp string, mobile_number string) (string, error){
	// it will store otp in db
	insertOTPQuery := fmt.Sprintf("INSERT INTO OTPDETAILS (otp, mobile_number) VALUES (%s, %s) ON CONFLICT (mobile_number) DO UPDATE SET otp = %s", otp, mobile_number, otp)
	
	conn, err := config.GetDBConnection()
	if err != nil {
		fmt.Println("Error while making connection for insert otp", err)
		return "", err
	}
	_ , err1 := conn.Exec(insertOTPQuery)
	if err1 != nil {
		fmt.Println("Not able to insert otp", err1)
		return "", err1
	}
	fmt.Println("OTP successfully inserted!")
	return "ok", nil
}

func GetOTP(mobile_number string) (string, error){
	//fetchOTPQuery := fmt.Sprintf("SELECT otp FROM OTPDETAILS WHERE mobile_number=%s",mobile_number)
	
	conn, err := config.GetDBConnection()
	if err != nil {
		fmt.Println("Error while making connection for fetch otp", err)
		return "", err
	}

	var otp string
	err = conn.QueryRow("SELECT otp FROM OTPDETAILS WHERE mobile_number = $1", mobile_number).Scan(&otp)
	if err != nil {
        fmt.Println("Not able fetch otp", err)
		return "", err
    }
	
	fmt.Println("OTP successfully fetched!", otp)
	return otp, nil
}

