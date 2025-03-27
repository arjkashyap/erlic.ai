package activedir

import (
	"context"
	"fmt"

	"github.com/arjkashyap/erlic.ai/internal/directory"

	"github.com/go-ldap/ldap/v3"
)

// retrieves a user from Active Directory by username
func (m *ADManager) GetUser(ctx context.Context, username string) (*directory.User, error) {
	// TODO
	user := &directory.User{}

	return user, nil
}

// CreateUser creates a new user in Active Directory
func (m *ADManager) CreateUser(ctx context.Context, user *directory.User) error {
	return nil
}

// UpdateUser updates an existing user in Active Directory
func (m *ADManager) UpdateUser(ctx context.Context, user *directory.User) error {

	return nil
}

// deletes a user from Active Directory
func (m *ADManager) DeleteUser(ctx context.Context, username string) error {
	conn, err := m.getConnection()
	if err != nil {
		return err
	}
	defer m.releaseConnection(conn)

	if err := m.bind(conn); err != nil {
		return err
	}

	// Find the user's DN
	userDN, err := m.findUserDN(conn, username)
	if err != nil {
		return err
	}

	// Create a delete request
	delReq := ldap.NewDelRequest(userDN, nil)

	// Execute the delete operation
	if err := conn.Del(delReq); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

// ResetUserPassword resets a user's password in Active Directory
func (m *ADManager) ResetUserPassword(ctx context.Context, username, newPassword string) error {
	conn, err := m.getConnection()
	if err != nil {
		return err
	}
	defer m.releaseConnection(conn)

	if err := m.bind(conn); err != nil {
		return err
	}

	// Find the user's DN
	userDN, err := m.findUserDN(conn, username)
	if err != nil {
		return err
	}

	// For AD, the password needs to be in a specific format
	// It must be surrounded by quotes and encoded as UTF-16LE
	utf16Password := encodePassword(newPassword)

	modifyReq := ldap.NewModifyRequest(userDN, nil)
	modifyReq.Replace("unicodePwd", []string{utf16Password})

	// Execute the modify operation
	if err := conn.Modify(modifyReq); err != nil {
		return fmt.Errorf("failed to reset password: %w", err)
	}

	return nil
}
