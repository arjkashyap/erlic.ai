# Active Directory USER

## Creating a new User
### [(m *ADManager) CreateUser(ctx context.Context, user *directory.User) error]

*Following are the REQUIRED attributes for creating a new user*

- objectClass - Must include "user" (typically set to multiple values like ["top", "person", "organizationalPerson", "user"])
- sAMAccountName - The login name used for older Windows clients (must be unique in the domain)
- cn (Common Name) - Used to create the Distinguished Name (DN) of the user
- userAccountControl - Controls account properties (whether it's disabled, etc.)
- userPrincipalName - Modern login format (username@domain.com)


#### Steps for creation
- Create a new disabled user
- Set password (requires TLS)
- Enable the user

#### Secure Communication
All Active Directory password management operations (setting, resetting, or changing passwords) must be performed over LDAP with TLS/SSL (LDAPS) for security reasons. This is an Active Directory requirement.

* Certificate Validation *

When connecting to LDAPS, our application must validate the server's certificate:

- If the Active Directory server uses a certificate from a publicly trusted Certificate Authority (CA), no additional configuration is needed.
- If the server uses a self-signed certificate or one signed by an internal CA, the client must:
  1. Obtain the CA certificate from the server administrator
  2. Configure the application to trust this certificate by providing the path to the CA certificate file

This additional configuration ensures secure and validated connections when performing sensitive password operations.