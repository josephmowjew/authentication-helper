package auth_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/josephmowjew/authentication-helper/pkg/models"
)

// Config holds configuration for authentication service
type Config struct {
	BaseURL    string
	HTTPClient *http.Client
	Timeout    time.Duration
}

// DefaultConfig creates a new configuration with default values
func DefaultConfig() Config {
	return Config{
		Timeout: 10 * time.Second,
	}
}

// NewConfig creates a new configuration for authentication service with custom settings
func NewConfig(baseURL string, opts ...func(*Config)) Config {
	config := DefaultConfig()
	config.BaseURL = baseURL

	// Apply any custom options
	for _, opt := range opts {
		opt(&config)
	}

	// Initialize HTTP client if not set
	if config.HTTPClient == nil {
		config.HTTPClient = &http.Client{
			Timeout: config.Timeout,
		}
	}

	return config
}

// Authenticate performs authentication against the external service
func Authenticate(username, password string, config Config) (*models.AuthResponse, error) {
	if config.BaseURL == "" {
		return nil, errors.New("base URL is required")
	}

	if config.HTTPClient == nil {
		config.HTTPClient = &http.Client{
			Timeout: config.Timeout,
		}
	}

	reqBody := models.AuthRequest{
		Username: username,
		Password: password,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", config.BaseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	// Set Content-Type header
	req.Header.Set("Content-Type", "application/json")

	resp, err := config.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Read error response body
		var errorResponse struct {
			Message string `json:"message"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			// If we can't decode the error response, return a generic error with status code
			return nil, fmt.Errorf("authentication failed with status code: %d", resp.StatusCode)
		}
		// Return the error message from the response if available
		if errorResponse.Message != "" {
			return nil, fmt.Errorf("authentication failed: %s", errorResponse.Message)
		}
		return nil, fmt.Errorf("authentication failed with status code: %d", resp.StatusCode)
	}

	var authResponse models.AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		return nil, err
	}

	return &authResponse, nil
}

// IsExpired checks if a token has expired based on its expiration timestamp
func IsExpired(exp float64) bool {
	return float64(time.Now().Unix()) > exp
}
