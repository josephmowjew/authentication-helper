# Authentication Helper

A Go package providing authentication utilities for interacting with the Lyvepulse authentication service.

## Package Structure

```
.
├── examples/     # Example implementations
├── pkg/
│   ├── models/   # Data models
│   └── token/    # Authentication client
```

## Installation

```bash
go get github.com/josephmojoo/authentication-helper
```

## Features

- Secure authentication against Lyvepulse authentication service
- Customizable HTTP client configuration
- Token expiration validation
- Comprehensive user information retrieval
- Role-based access control support

## Usage

```go
package main

import (
    "fmt"
    "log"
    "time"
    
    auth_client "github.com/josephmojoo/authentication-helper/pkg/token"
)

func main() {
    // Create a custom configuration for the authentication service
    config := auth_client.NewConfig(
        "https://your-auth-service-url",
        func(c *auth_client.Config) {
            c.Timeout = 15 * time.Second // Custom timeout
        },
    )

    // Attempt authentication
    response, err := auth_client.Authenticate("username", "password", config)
    if err != nil {
        log.Fatalf("Authentication failed: %v", err)
    }

    // Access token and user information
    fmt.Printf("Token: %s\n", response.Token)
    user := response.AccessTicket

    // Check token expiration
    if auth_client.IsExpired(user.Exp) {
        fmt.Println("Token has expired!")
    }

    // Access user roles
    for _, role := range user.Roles {
        fmt.Printf("Role ID: %s, Branch ID: %s\n", role.RoleId, role.BranchId)
    }
}
```

## Configuration Options

The authentication client can be configured with the following options:

- `BaseURL`: The URL of the authentication service
- `Timeout`: Custom timeout for HTTP requests (default: 10 seconds)
- `HTTPClient`: Custom HTTP client for special requirements

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.