package directory

import "context"

// interface for directory management operations
type DirectoryManager interface {
	// User operations
	GetUser(ctx context.Context, username string) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, username string) error
	ResetUserPassword(ctx context.Context, username, newPassword string) error

	// Group operations
	// GetGroup(ctx context.Context, groupName string) (*Group, error)
	AddUserToGroup(ctx context.Context, username, groupName string) error
	RemoveUserFromGroup(ctx context.Context, username, groupName string) error
	// ListGroups(ctx context.Context, filter string) ([]*Group, error)
	ListGroupMembers(ctx context.Context, groupName string) ([]*User, error)
}

// Directory User
type User struct {
	Username    string
	FirstName   string
	LastName    string
	DisplayName string
	Email       string
	Department  string
	Title       string
	Description string
	Enabled     bool
	Groups      []string
}

// Directory Group
type Group struct {
	Name        string
	Description string
	Email       string
	Members     []string
}

// ErrorNotFound is returned when a requested resource is not found
type ErrorNotFound struct {
	Message string
}

func (e *ErrorNotFound) Error() string {
	return e.Message
}

// ErrorPermissionDenied is returned when the operation is not permitted
type ErrorPermissionDenied struct {
	Message string
}

func (e *ErrorPermissionDenied) Error() string {
	return e.Message
}
