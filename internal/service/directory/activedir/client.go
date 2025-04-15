package activedir

import (
	"crypto/tls"
	"fmt"
	"sync"
	"time"

	"github.com/arjkashyap/erlic.ai/internal/logger"
	"github.com/arjkashyap/erlic.ai/internal/service/directory"
	"github.com/arjkashyap/erlic.ai/internal/utils"
	"github.com/go-ldap/ldap/v3"
)

// ADManager implements the DirectoryManager interface for Active Directory
type ADManager struct {
	ldapURL        string
	baseDN         string
	bindUsername   string
	bindPassword   string
	serverHostName string
	useTLS         bool
	connPool       sync.Pool
	caCertPath     string
}

// NewADManager creates a new Active Directory manager
func NewADManager(ldapURL string, baseDN string, bindUsername string, bindPassword string, serverHostName string, useTLS bool, caCertPath string) *ADManager {
	adm := &ADManager{
		ldapURL:        ldapURL,
		baseDN:         baseDN,
		bindUsername:   bindUsername,
		bindPassword:   bindPassword,
		serverHostName: serverHostName,
		useTLS:         useTLS,
		caCertPath:     caCertPath,
	}
	// Initialize connection pool
	adm.connPool = sync.Pool{
		New: func() interface{} {
			conn, err := adm.dial()
			if err != nil {
				return nil
			}
			return conn
		},
	}
	return adm
}

// establishes a connection to the LDAP server
func (m *ADManager) dial() (*ldap.Conn, error) {
	var conn *ldap.Conn
	var err error

	if m.useTLS {
		host := m.ldapURL
		logger.Logger.Info("Using TLS with host: " + host)

		caCertPool, err := utils.LoadCACert(m.caCertPath)
		if err != nil {
			return nil, fmt.Errorf("failed to load CA cert: %w", err)
		}

		tlsConfig := &tls.Config{
			RootCAs:            caCertPool,
			ServerName:         m.serverHostName,
			InsecureSkipVerify: false,
		}

		conn, err = ldap.DialURL(m.ldapURL, ldap.DialWithTLSConfig(tlsConfig))
		if err != nil {
			return nil, fmt.Errorf("LDAP TLS connection error: %w", err)
		}
	} else {
		conn, err = ldap.DialURL(m.ldapURL)
		if err != nil {
			return nil, fmt.Errorf("LDAP connection error: %w", err)
		}
	}

	logger.Logger.Info("Connection Established with LDAP server")
	conn.SetTimeout(5 * time.Second)
	return conn, nil
}

// gets a connection from the pool or creates a new one
func (m *ADManager) getConnection() (*ldap.Conn, error) {
	logger.Logger.Info("getConnection() - called")
	conn := m.connPool.Get()
	if conn == nil {
		return m.dial()
	}

	ldapConn, ok := conn.(*ldap.Conn)
	if !ok || ldapConn == nil {
		return m.dial()
	}

	return ldapConn, nil
}

// returns a connection to the pool
func (m *ADManager) releaseConnection(conn *ldap.Conn) {
	if conn != nil {
		m.connPool.Put(conn)
	}
}

// bind authenticates with the LDAP server
func (m *ADManager) bind(conn *ldap.Conn) error {
	err := conn.Bind(m.bindUsername, m.bindPassword)
	if err != nil {
		return fmt.Errorf("LDAP bind error: %w", err)
	}
	return nil
}

// findUserDN finds a user's DN by their username
func (m *ADManager) findUserDN(conn *ldap.Conn, username string) (string, error) {
	searchRequest := ldap.NewSearchRequest(
		m.baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=user)(sAMAccountName=%s))", ldap.EscapeFilter(username)),
		[]string{"dn"},
		nil,
	)

	sr, err := conn.Search(searchRequest)
	if err != nil {
		return "", fmt.Errorf("LDAP search error: %w", err)
	}

	if len(sr.Entries) == 0 {
		return "", &directory.ErrorNotFound{Message: fmt.Sprintf("User '%s' not found", username)}
	}

	return sr.Entries[0].DN, nil
}

// findGroupDN finds a group's DN by its name
func (m *ADManager) findGroupDN(conn *ldap.Conn, groupName string) (string, error) {
	searchRequest := ldap.NewSearchRequest(
		m.baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=group)(cn=%s))", ldap.EscapeFilter(groupName)),
		[]string{"dn"},
		nil,
	)

	sr, err := conn.Search(searchRequest)
	if err != nil {
		return "", fmt.Errorf("LDAP search error: %w", err)
	}

	if len(sr.Entries) == 0 {
		return "", &directory.ErrorNotFound{Message: fmt.Sprintf("Group '%s' not found", groupName)}
	}

	return sr.Entries[0].DN, nil
}
