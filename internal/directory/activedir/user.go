package activedir

import (
	"context"
	"fmt"

	"github.com/arjkashyap/erlic.ai/internal/directory"
	"golang.org/x/text/encoding/unicode"

	"github.com/go-ldap/ldap/v3"
)

// retrieves a user from Active Directory by username
func (m *ADManager) GetUser(ctx context.Context, username string) (*directory.User, error) {
	// TODO
	user := &directory.User{}

	return user, nil
}

func (m *ADManager) CreateUser(ctx context.Context, user *directory.User) error {
	default_pass := "BigH3ro@101_ComplexP@ss!"
	conn, err := m.getConnection()

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer m.releaseConnection(conn)

	if err := m.bind(conn); err != nil {
		return err
	}

	fmt.Println("Creating new user request")

	dn := fmt.Sprintf("CN=%s,CN=Users,%s", user.Username, m.baseDN)
	fmt.Println("DN = " + dn)

	addReq := ldap.NewAddRequest(dn, []ldap.Control{})
	addReq.Attribute("objectClass", []string{"top", "organizationalPerson", "user", "person"})
	addReq.Attribute("name", []string{user.Username})
	addReq.Attribute("sAMAccountName", []string{user.Username})
	addReq.Attribute("userAccountControl", []string{fmt.Sprintf("%d", 0x0202)}) // Disabled account at first
	addReq.Attribute("userPrincipalName", []string{user.Email})
	addReq.Attribute("accountExpires", []string{fmt.Sprintf("%d", 0x00000000)})

	// Add name attributes
	if user.FirstName != "" {
		addReq.Attribute("givenName", []string{user.FirstName})
	}
	if user.LastName != "" {
		addReq.Attribute("sn", []string{user.LastName})
	}
	if user.DisplayName != "" {
		addReq.Attribute("displayName", []string{user.DisplayName})
	}

	if user.Department != "" {
		addReq.Attribute("department", []string{user.Department})
	}
	if user.Title != "" {
		addReq.Attribute("title", []string{user.Title})
	}
	if user.Description != "" {
		addReq.Attribute("description", []string{user.Description})
	}

	if err := conn.Add(addReq); err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to create user: %w", err)
	}

	fmt.Println("User is created")

	// Setting up user password
	quotedPassword := fmt.Sprintf("\"%s\"", default_pass)
	utf16le := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	pwdEncoded, err := utf16le.NewEncoder().String(quotedPassword)
	if err != nil {
		return err
	}

	modReq := ldap.NewModifyRequest(dn, nil)
	modReq.Replace("unicodePwd", []string{pwdEncoded})

	if err := conn.Modify(modReq); err != nil {
		fmt.Println("Error with PasswordModify operation:", err)
		return err

	}

	fmt.Println("User Password Set")

	// Enable the account by setting userAccountControl to 512 (normal account)
	enableReq := ldap.NewModifyRequest(dn, nil)
	enableReq.Replace("userAccountControl", []string{"512"})
	if err := conn.Modify(enableReq); err != nil {
		fmt.Println("Error enabling user account:", err)
		return err
	}

	fmt.Println("User account enabled")
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
