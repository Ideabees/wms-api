package thirdparty

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"wms-app/config"
)

func SendMessageRequest(phoneNumber string, message string) (string, error) {
	apiURL := GetSendMessageURL()

	// Prepare the form data
	formData := url.Values{}
	formData.Set("channel", "whatsapp")
	formData.Set("source", "<Your WhatsApp Number in E.164 format>") // Source is usually the WhatsApp number in E.164 format, but can be left empty if not required
	formData.Set("destination", phoneNumber)
	formData.Set("message", `{"type":"text","text":"`+message+`"}`)
	formData.Set("src.name", "<Your WhatsApp Business Name>")

	// Create a new POST request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(formData.Encode()))
	if err != nil {
		return "", err
	}

	// Set headers
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("apikey", getAPIKey())
	req.Header.Set("cache-control", "no-cache")

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

func GetSendMessageURL() string {
	return config.APIBaseURL + config.EndpointSendMessage
}

func getAPIKey() string {
	return config.APIKey
}