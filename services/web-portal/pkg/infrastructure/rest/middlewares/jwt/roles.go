package jwt

// Role represents a requesting user role
type Role int

const (
	// VISITOR role
	VISITOR Role = iota
	// USER role
	USER Role = 1
	// ADMIN role
	ADMIN Role = 99999
)

// ToRole converts string to Role type
func ToRole(role string) Role {
	switch role {
	case "USER":
		return USER
	case "ADMIN":
		return ADMIN
	}
	return VISITOR
}
