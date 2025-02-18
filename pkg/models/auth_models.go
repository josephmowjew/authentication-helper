package models

// AuthRequest represents the authentication request payload
type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Role represents a user role in the system
type Role struct {
	RoleId           string `json:"roleId"`
	BranchId         string `json:"branchId"`
	OrganisationalId string `json:"organisationalId"`
}

// AccessTicket represents the user access information
type AccessTicket struct {
	Sub          string `json:"sub"`
	Username     string `json:"username"`
	EmployeeId   string `json:"employeeId"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	PhoneNumber  string `json:"phoneNumber"`
	Enabled      bool   `json:"enabled"`
	PendingReset bool   `json:"pendingReset"`
	Roles        []Role `json:"roles"`
	Iat          float64  `json:"iat"`
	Exp          float64  `json:"exp"`
}

// AuthResponse represents the authentication response from the service
type AuthResponse struct {
	Token        string       `json:"token"`
	AccessTicket AccessTicket `json:"accessTicket"`
}