package activedir

import (
	"context"
	"fmt"
	"reflect"
	"strconv"

	"golang.org/x/text/encoding/unicode"

	"github.com/arjkashyap/erlic.ai/internal/logger"
	"github.com/arjkashyap/erlic.ai/internal/service/directory"
	"github.com/go-ldap/ldap/v3"
)

func (m *ADManager) CreateUser(ctx context.Context, user *directory.User) error {
	default_pass := "BigH3ro@101_ComplexP@ss!"
	conn, err := m.getConnection()

	if err != nil {
		logger.Logger.Error(err)
		return err
	}

	defer m.releaseConnection(conn)

	if err := m.bind(conn); err != nil {
		return err
	}

	logger.Logger.Info("Creating new user request")

	dn := fmt.Sprintf("CN=%s,CN=Users,%s", user.Username, m.baseDN)
	logger.Logger.Info("DN = " + dn)

	addReq := ldap.NewAddRequest(dn, []ldap.Control{})
	addReq.Attribute("objectClass", []string{"top", "organizationalPerson", "user", "person"})
	addReq.Attribute("name", []string{user.Username})
	addReq.Attribute("sAMAccountName", []string{user.Username})
	addReq.Attribute("userAccountControl", []string{fmt.Sprintf("%d", 0x0202)}) // Disabled account at first
	addReq.Attribute("userPrincipalName", []string{user.Email})
	addReq.Attribute("accountExpires", []string{fmt.Sprintf("%d", 0x00000000)})

	if user.FirstName != "" {
		addReq.Attribute("givenName", []string{user.FirstName})
	}
	if user.LastName != "" {
		addReq.Attribute("sn", []string{user.LastName})
	}
	if user.DisplayName != "" {
		addReq.Attribute("displayName", []string{user.DisplayName})
	}
	if user.Email != "" {
		addReq.Attribute("mail", []string{user.Email})
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
		logger.Logger.Error(err)
		return fmt.Errorf("failed to create user: %w", err)
	}

	logger.Logger.Info("User is created")

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
		logger.Logger.Error("Error with PasswordModify operation:", err)
		return err

	}

	logger.Logger.Info("User Password Set")

	// Enable the account by setting userAccountControl to 512 (normal account)
	enableReq := ldap.NewModifyRequest(dn, nil)
	enableReq.Replace("userAccountControl", []string{"512"})
	if err := conn.Modify(enableReq); err != nil {
		logger.Logger.Error("Error enabling user account:", err)
		return err
	}

	logger.Logger.Info("User account enabled")
	return nil
}

// retrieves a user from Active Directory by username
func (m *ADManager) GetUser(ctx context.Context, username string) (*directory.User, error) {
	conn, err := m.getConnection()
	if err != nil {
		return nil, err
	}
	defer m.releaseConnection(conn)

	if err := m.bind(conn); err != nil {
		return nil, err
	}

	// Create search filter for the user
	filter := fmt.Sprintf("(&(objectClass=user)(sAMAccountName=%s))", ldap.EscapeFilter(username))

	// attributes to retrieve from search query
	attributes := []string{
		"sAMAccountName", "givenName", "sn", "displayName",
		"mail", "department", "title", "description",
		"userAccountControl", "memberOf",
	}

	// Create search request
	// TODO: adjust params - [Scope,DerefAliases,SizeLimit] according to search
	searchRequest := ldap.NewSearchRequest(
		m.baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter,
		attributes,
		nil,
	)

	// Execute search
	sr, err := conn.Search(searchRequest)
	if err != nil {
		return nil, fmt.Errorf("LDAP search error: %w", err)
	}

	// Check if user was found
	if len(sr.Entries) == 0 {
		return nil, &directory.ErrorNotFound{Message: fmt.Sprintf("User '%s' not found", username)}
	}

	// Get first entry
	entry := sr.Entries[0]

	// Extract user account control value to determine if account is enabled
	uacStr := entry.GetAttributeValue("userAccountControl")
	uac, _ := strconv.Atoi(uacStr)
	enabled := (uac & 2) == 0 // Account is disabled if bit 1 is set

	// TODO: Use function Get groups

	// Create user object
	user := &directory.User{
		Username:    entry.GetAttributeValue("sAMAccountName"),
		FirstName:   entry.GetAttributeValue("givenName"),
		LastName:    entry.GetAttributeValue("sn"),
		DisplayName: entry.GetAttributeValue("displayName"),
		Email:       entry.GetAttributeValue("mail"),
		Department:  entry.GetAttributeValue("department"),
		Title:       entry.GetAttributeValue("title"),
		Description: entry.GetAttributeValue("description"),
		Enabled:     enabled,
		Groups:      make([]string, 0),
	}

	return user, nil
}

// UpdateUser updates an existing user in Active Directory
func (m *ADManager) UpdateUser(ctx context.Context, user *directory.User) error {
	conn, err := m.getConnection()
	if err != nil {
		return err
	}
	defer m.releaseConnection(conn)

	if err := m.bind(conn); err != nil {
		return err
	}

	userDN, err := m.findUserDN(conn, user.Username)
	if err != nil {
		return fmt.Errorf("cannot update user, user not found: %w", err)
	}

	modReq := ldap.NewModifyRequest(userDN, nil)

	// Build modifications based on non-empty fields
	modifications := BuildModifications(user)

	for attr, values := range modifications {
		modReq.Replace(string(attr), values)
	}

	// Handle enabled state specially (similar to your current code)
	if user != nil {
		// ... handle enabled state with userAccountControl ...
	}

	if len(modReq.Changes) > 0 {
		if err := conn.Modify(modReq); err != nil {
			return fmt.Errorf("failed to update user: %w", err)
		}
	}

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

	logger.Logger.Info("User %s deleted from AD %s", username, m.serverHostName)
	return nil
}

// resets a user's password in Active Directory
func (m *ADManager) ResetUserPassword(ctx context.Context, username, newPassword string) error {
	conn, err := m.getConnection()
	if err != nil {
		return err
	}
	defer m.releaseConnection(conn)

	if err := m.bind(conn); err != nil {
		return err
	}

	userDN, err := m.findUserDN(conn, username)
	if err != nil {
		return fmt.Errorf("failed to find user for password reset: %w", err)
	}

	quotedPassword := fmt.Sprintf("\"%s\"", newPassword)
	utf16le := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	pwdEncoded, err := utf16le.NewEncoder().String(quotedPassword)
	if err != nil {
		return fmt.Errorf("failed to encode password: %w", err)
	}

	modReq := ldap.NewModifyRequest(userDN, nil)
	modReq.Replace("unicodePwd", []string{pwdEncoded})

	// Execute the password modification
	if err := conn.Modify(modReq); err != nil {
		return fmt.Errorf("failed to reset password: %w", err)
	}

	logger.Logger.Info("Password reset successful for user: %s", username)
	return nil
}

func BuildModifications(user *directory.User) map[ADAttribute][]string {
	modifications := make(map[ADAttribute][]string)

	val := reflect.ValueOf(*user)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name

		if fieldName == "Groups" || fieldName == "Enabled" {
			continue
		}

		if field.Kind() == reflect.String && field.String() != "" {
			if adAttr, exists := UserAttributeMapping[fieldName]; exists {
				modifications[adAttr] = []string{field.String()}
			}
		}
	}

	// Handle the Enabled field specially since it maps to userAccountControl
	// and requires bit manipulation
	if _, ok := typ.FieldByName("Enabled"); ok {
		// This field requires special handling in the caller since we need
		// current value to toggle bits
	}

	return modifications
}
