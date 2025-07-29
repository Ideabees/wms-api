package config

import (
	"time"
)

const (
	// API constants
	APIBaseURL       = "https://api.gupshup.io"
	APITimeout       = 60 * time.Second
	APIUserAgent     = "WMS-App/1.0"
	APIDefaultHeader = "application/json"
	APIMaxRetries    = 3

	// Endpoints
	EndpointOptInNumber = "/sm/api/v1/app/opt/in/appname"
	EndpointSendMessage = "/wa/api/v1/msg"
	AppName             = "convoEngage"
	APIKey              = "<Your API Key>" // Replace with your actual API key
)
