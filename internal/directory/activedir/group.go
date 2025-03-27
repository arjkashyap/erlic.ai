package activedir

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/arjkashyap/erlic.ai/internal/directory"
	"github.com/go-ldap/ldap/v3"
)

// AddUserToGroup adds a user to a group
func (m *ADManager) AddUserToGroup(ctx context.Context, username, groupName string) error {
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

	// Find the group's DN
	groupDN, err := m.findGroupDN(conn, groupName)
	if err != nil {
		return err
	}

	// Create a modify request to add the member
	modReq := ldap.NewModifyRequest(groupDN, nil)
	modReq.Add("member", []string{userDN})

	// Execute the modify operation
	if err := conn.Modify(modReq); err != nil {
		if strings.Contains(err.Error(), "LDAP: Type or value exists") {
			// User is already a member of the group
			return nil
		}
		return fmt.Errorf("failed to add user to group: %w", err)
	}

	return nil
}

// removes a user from a group
func (m *ADManager) RemoveUserFromGroup(ctx context.Context, username, groupName string) error {
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

	// Find the group's DN
	groupDN, err := m.findGroupDN(conn, groupName)
	if err != nil {
		return err
	}

	// Create a modify request to remove the member
	modReq := ldap.NewModifyRequest(groupDN, nil)
	modReq.Delete("member", []string{userDN})

	// Execute the modify operation
	if err := conn.Modify(modReq); err != nil {
		if strings.Contains(err.Error(), "LDAP: No such attribute") {
			// User is not a member of the group
			return nil
		}
		return fmt.Errorf("failed to remove user from group: %w", err)
	}

	return nil
}

// ListGroupMembers lists all members of a group
func (m *ADManager) ListGroupMembers(ctx context.Context, groupName string) ([]*directory.User, error) {
	conn, err := m.getConnection()
	if err != nil {
		return nil, err
	}
	defer m.releaseConnection(conn)

	if err := m.bind(conn); err != nil {
		return nil, err
	}

	// Find the group's DN
	groupDN, err := m.findGroupDN(conn, groupName)
	if err != nil {
		return nil, err
	}

	// Get the group's members
	searchReq := ldap.NewSearchRequest(
		groupDN,
		ldap.ScopeBaseObject, ldap.NeverDerefAliases, 0, 0, false,
		"(objectClass=*)",
		[]string{"member"},
		nil,
	)

	sr, err := conn.Search(searchReq)
	if err != nil {
		return nil, fmt.Errorf("failed to search group members: %w", err)
	}

	if len(sr.Entries) == 0 {
		return nil, fmt.Errorf("group not found during member listing")
	}

	members := sr.Entries[0].GetAttributeValues("member")
	users := make([]*directory.User, 0, len(members))

	// For each member DN, fetch the user details
	for _, memberDN := range members {
		// Skip non-user members (like nested groups)
		if !strings.Contains(memberDN, "CN=Users") && !strings.Contains(memberDN, "OU=Users") {
			continue
		}

		searchReq := ldap.NewSearchRequest(
			memberDN,
			ldap.ScopeBaseObject, ldap.NeverDerefAliases, 0, 0, false,
			"(objectClass=user)",
			[]string{
				"sAMAccountName", "givenName", "sn", "displayName",
				"mail", "userAccountControl",
			},
			nil,
		)

		sr, err := conn.Search(searchReq)
		if err != nil {
			// Log the error but continue with other members
			continue
		}

		if len(sr.Entries) == 0 {
			continue
		}

		entry := sr.Entries[0]
		uacStr := entry.GetAttributeValue("userAccountControl")
		uac, _ := strconv.Atoi(uacStr)
		enabled := (uac & 2) == 0

		user := &directory.User{
			Username:    entry.GetAttributeValue("sAMAccountName"),
			FirstName:   entry.GetAttributeValue("givenName"),
			LastName:    entry.GetAttributeValue("sn"),
			DisplayName: entry.GetAttributeValue("displayName"),
			Email:       entry.GetAttributeValue("mail"),
			Enabled:     enabled,
		}

		users = append(users, user)
	}

	return users, nil
}

// encodePassword converts a password to the format required by Active Directory
func encodePassword(password string) string {
	// Password must be enclosed in quotes
	quoted := fmt.Sprintf("\"%s\"", password)

	// Convert to UTF-16LE bytes
	utf16 := make([]byte, 0, len(quoted)*2)
	for _, r := range quoted {
		utf16 = append(utf16, byte(r))
		utf16 = append(utf16, 0)
	}

	return string(utf16)
}
