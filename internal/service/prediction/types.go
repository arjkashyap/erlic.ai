package prediction

type FunctionCall struct {
	Function   string     `json:"function"`
	Parameters Parameters `json:"parameters"`
}

type Parameters struct {
	Ctx      string      `json:"ctx"`
	Username string      `json:"username,omitempty"`
	User     *UserParams `json:"user,omitempty"`
}

type UserParams struct {
	User UserDetails `json:"user"`
}

type UserDetails struct {
	Username    string   `json:"Username"`
	FirstName   string   `json:"FirstName,omitempty"`
	LastName    string   `json:"LastName,omitempty"`
	DisplayName string   `json:"DisplayName,omitempty"`
	Email       string   `json:"Email,omitempty"`
	Department  string   `json:"Department,omitempty"`
	Title       string   `json:"Title,omitempty"`
	Description string   `json:"Description,omitempty"`
	Enabled     bool     `json:"Enabled,omitempty"`
	Groups      []string `json:"Groups,omitempty"`
}
