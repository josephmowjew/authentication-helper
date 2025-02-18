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
		"https://your-auth-service-url", // Replace with your authentication service URL
		// Custom options
		func(c *auth_client.Config) {
			c.Timeout = 15 * time.Second // Custom timeout
		},
	)

	// Attempt authentication
	username := "your-username@example.com" // Replace with your username
	password := "your-password"           // Replace with your password

	// Perform authentication
	response, err := auth_client.Authenticate(username, password, config)
	if err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}

	// Successfully authenticated
	fmt.Printf("Authentication successful!\n")
	fmt.Printf("Token: %s\n", response.Token)

	// Print user information
	user := response.AccessTicket
	fmt.Printf("User Details:\n")
	fmt.Printf("Name: %s %s\n", user.FirstName, user.LastName)
	fmt.Printf("Employee ID: %s\n", user.EmployeeId)
	fmt.Printf("Phone: %s\n", user.PhoneNumber)

	// Check if token is expired
	if auth_client.IsExpired(user.Exp) {
		fmt.Println("Warning: Token has expired!")
	} else {
		fmt.Println("Token is valid")
		// Calculate remaining time
		remaining := time.Unix(int64(user.Exp), 0).Sub(time.Now())
		fmt.Printf("Token expires in: %v\n", remaining.Round(time.Second))
	}

	// Print user roles
	fmt.Println("\nUser Roles:")
	for _, role := range user.Roles {
		fmt.Printf("- Role ID: %s\n  Branch ID: %s\n  Org ID: %s\n",
			role.RoleId,
			role.BranchId,
			role.OrganisationalId)
	}

	// Additional status checks
	if user.PendingReset {
		fmt.Println("\nWarning: Password reset is pending!")
	}

	if !user.Enabled {
		fmt.Println("Warning: User account is disabled!")
	}
}
