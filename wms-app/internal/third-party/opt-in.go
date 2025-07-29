package thirdparty

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"wms-app/config"
)

func MakeOptInRequest(phoneNumber string) (string, error) {
	url := GetOptInURL()

	// Prepare the request body
	data := []byte("user=" + phoneNumber)

	// Create a new POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("apikey", GetAPIKey())

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("non-200 response: " + string(body))
	}

	return string(body), nil
}

func GetOptInURL() string {
	return config.APIBaseURL + config.EndpointOptInNumber
}

func GetAPIKey() string {
	return config.APIKey
}