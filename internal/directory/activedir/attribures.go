package activedir

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/arjkashyap/erlic.ai/internal/directory"
	"github.com/go-ldap/ldap/v3"
)

type ADAttribute string

// constants for all AD attributes
const (
	// Common attributes
	ADAttributeUsername       ADAttribute = "sAMAccountName"
	ADAttributeFirstName      ADAttribute = "givenName"
	ADAttributeLastName       ADAttribute = "sn"
	ADAttributeDisplayName    ADAttribute = "displayName"
	ADAttributeEmail          ADAttribute = "mail"
	ADAttributeDepartment     ADAttribute = "department"
	ADAttributeTitle          ADAttribute = "title"
	ADAttributeDescription    ADAttribute = "description"
	ADAttributeEnabled        ADAttribute = "userAccountControl"
	ADAttributeCommonName     ADAttribute = "cn"
	ADAttributeDistinguished  ADAttribute = "distinguishedName"
	ADAttributeInitials       ADAttribute = "initials"
	ADAttributeObjectCategory ADAttribute = "objectCategory"
	ADAttributeObjectClass    ADAttribute = "objectClass"
	ADAttributeOrgUnit        ADAttribute = "ou"
	ADAttributeUserPrincipal  ADAttribute = "userPrincipalName"
	ADAttributeMemberOf       ADAttribute = "memberOf"
)

// maps User struct fields to AD attributes
var UserAttributeMapping = map[string]ADAttribute{
	"Username":    ADAttributeUsername,
	"FirstName":   ADAttributeFirstName,
	"LastName":    ADAttributeLastName,
	"DisplayName": ADAttributeDisplayName,
	"Email":       ADAttributeEmail,
	"Department":  ADAttributeDepartment,
	"Title":       ADAttributeTitle,
	"Description": ADAttributeDescription,
	// Map additional fields if you add them to your User struct
	// "Initials":      ADAttributeInitials,
	// "UserPrincipal": ADAttributeUserPrincipal,
}

// Inverse mapping - can be useful for processing search results
var ADAttributeToUserField = map[ADAttribute]string{
	ADAttributeUsername:    "Username",
	ADAttributeFirstName:   "FirstName",
	ADAttributeLastName:    "LastName",
	ADAttributeDisplayName: "DisplayName",
	ADAttributeEmail:       "Email",
	ADAttributeDepartment:  "Department",
	ADAttributeTitle:       "Title",
	ADAttributeDescription: "Description",
}

// Special attributes that need custom handling
var SpecialAttributes = map[ADAttribute]bool{
	ADAttributeEnabled:       true,
	ADAttributeDistinguished: true,
	ADAttributeMemberOf:      true,
}

// UserToADMap converts a User object to an AD attribute map
func UserToADMap(user *directory.User) map[ADAttribute][]string {
	attributes := make(map[ADAttribute][]string)

	// Use reflection to process non-empty fields
	val := reflect.ValueOf(*user)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name

		// Skip special fields or empty values
		if fieldName == "Groups" || fieldName == "Enabled" {
			continue
		}

		// Only process string fields that aren't empty
		if field.Kind() == reflect.String && field.String() != "" {
			if adAttr, exists := UserAttributeMapping[fieldName]; exists {
				attributes[adAttr] = []string{field.String()}
			}
		}
	}

	// Handle any special mappings or defaults

	// Ensure objectClass is set properly for user objects
	if _, ok := attributes[ADAttributeObjectClass]; !ok {
		attributes[ADAttributeObjectClass] = []string{
			"top", "person", "organizationalPerson", "user",
		}
	}

	return attributes
}

// ADMapToUser converts AD attributes to a User object
func ADMapToUser(attributes map[string][]string) *directory.User {
	user := &directory.User{}

	// Process each AD attribute
	for attr, values := range attributes {
		if len(values) == 0 {
			continue
		}

		// Convert string attribute to ADAttribute for lookups
		adAttr := ADAttribute(attr)

		// Process based on attribute type
		if fieldName, exists := ADAttributeToUserField[adAttr]; exists {
			// Use reflection to set the field
			field := reflect.ValueOf(user).Elem().FieldByName(fieldName)
			if field.IsValid() && field.CanSet() && field.Kind() == reflect.String {
				field.SetString(values[0])
			}
		} else if attr == string(ADAttributeEnabled) {
			// Handle userAccountControl specially
			if uac, err := strconv.Atoi(values[0]); err == nil {
				user.Enabled = (uac & 2) == 0 // Account is enabled if disabled bit is not set
			}
		} else if attr == string(ADAttributeMemberOf) {
			// Process group memberships
			groups := make([]string, 0, len(values))
			for _, dn := range values {
				// Extract CN from DN
				if cn := extractCNFromDN(dn); cn != "" {
					groups = append(groups, cn)
				}
			}
			user.Groups = groups
		}
	}

	return user
}

// Helper to extract CN from a DN string
func extractCNFromDN(dn string) string {
	parsedDN, err := ldap.ParseDN(dn)
	if err != nil {
		return "" // Return empty string if DN parsing fails
	}

	// Iterate through the RDNs (Relative Distinguished Names)
	for _, rdn := range parsedDN.RDNs {
		for _, attr := range rdn.Attributes {
			if strings.EqualFold(attr.Type, "CN") {
				return attr.Value
			}
		}
	}

	return ""
}
